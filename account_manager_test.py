import unittest
from unittest.mock import patch, mock_open
import json

class TestAccountManager(unittest.TestCase):

    @patch("builtins.open", new_callable=mock_open, read_data='{"nodes": ["http://node1", "http://node2"]}')
    def test_setHttpProvidersFromFile_valid_file(self, mock_file):
        am = AccountManager()
        am.setHttpProvidersFromFile("httpProviders.json")
        self.assertEqual(am.httpProviders, ["http://node1", "http://node2"])

    @patch("builtins.open", new_callable=mock_open, read_data='{"invalid": "data"}')
    def test_setHttpProvidersFromFile_no_nodes_key(self, mock_file):
        am = AccountManager()
        am.setHttpProvidersFromFile("httpProviders.json")
        self.assertEqual(am.httpProviders, [])

    @patch("builtins.open", new_callable=mock_open, read_data='invalid json')
    def test_setHttpProvidersFromFile_invalid_json(self, mock_file):
        am = AccountManager()
        with self.assertRaises(ValueError):
            am.setHttpProvidersFromFile("httpProviders.json")

    def test_setHttpProvidersFromFile_empty_file(self):
        am = AccountManager()
        with self.assertRaises(ValueError):
            am.setHttpProvidersFromFile("")

    @patch("builtins.open", side_effect=FileNotFoundError)
    def test_setHttpProvidersFromFile_file_not_found(self, mock_file):
        am = AccountManager()
        with self.assertRaises(FileNotFoundError):
            am.setHttpProvidersFromFile("nonexistent.json")

    def test_nextHttpProvider_no_providers(self):
        am = AccountManager()
        with self.assertRaises(IndexError):
            am.nextHttpProvider()

    def test_nextHttpProvider_valid_providers(self):
        am = AccountManager()
        am.httpProviders = ["http://node1", "http://node2"]
        am.nextHttpProvider()
        self.assertEqual(am.httpProviderIndex, 1)

    def test_getTx_minimal_params(self):
        am = AccountManager()
        tx = am.getTx(to="0x456")
        expected = {
            "from": "0x123",
            "to": "0x456",
            "nonce": 1,
            "value": 0
        }
        self.assertEqual(tx, expected)

    def test_getTx_with_data(self):
        am = AccountManager()
        tx = am.getTx(to="0x456", data=b"test_data", nonce=5)
        expected = {
            "from": "0x123",
            "to": "0x456",
            "nonce": 5,
            "value": 0,
            "data": b"test_data"
        }
        self.assertEqual(tx, expected)


if __name__ == "__main__":
    unittest.main()
