import json
from typing import List, Dict, Optional, Union
from eth_typing import Address


class AccountManager:
    def __init__(self):
        self.httpProviders: List[str] = []
        self.httpProviderIndex: int = 0
        self.account_address: str = "0x123"  # Default address for testing

    def setHttpProvidersFromFile(self, httpProvidersFile: str) -> None:
        """Load HTTP providers from a JSON file."""
        if not httpProvidersFile:
            raise ValueError("httpProvidersFile cannot be empty")

        try:
            with open(httpProvidersFile, "r") as f:
                data = json.load(f)

            if "nodes" not in data:
                self.httpProviders = []
                return

            self.httpProviders = [
                node for node in data["nodes"] if isinstance(node, str)
            ]

        except json.JSONDecodeError:
            raise ValueError("Invalid JSON format in providers file")

    def nextHttpProvider(self) -> None:
        """Switch to the next HTTP provider in the list."""
        if not self.httpProviders:
            raise IndexError("No HTTP providers available")

        self.httpProviderIndex = (self.httpProviderIndex + 1) % len(self.httpProviders)

    def getTx(
        self,
        to: str,
        value: int = 0,
        data: Optional[bytes] = None,
        nonce: Optional[int] = None,
        gas: Optional[int] = None,
        gasPrice: Optional[int] = None,
        gasPriceMultiplier: float = 1.0,
    ) -> Dict[str, Union[str, int, bytes]]:
        """
        Create a transaction dictionary with the specified parameters.

        Args:
            to: Destination address
            value: Amount of ETH to send (in wei)
            data: Transaction data
            nonce: Transaction nonce (if None, will use 1 for testing)
            gas: Gas limit
            gasPrice: Gas price in wei
            gasPriceMultiplier: Multiplier for gas price

        Returns:
            Dictionary containing transaction parameters
        """
        tx = {
            "from": self.account_address,
            "to": to,
            "nonce": nonce if nonce is not None else 1,
            "value": value,
        }

        if data is not None:
            tx["data"] = data

        if gas is not None:
            tx["gas"] = gas

        if gasPrice is not None:
            final_gas_price = int(gasPrice * gasPriceMultiplier)
            tx["gasPrice"] = final_gas_price

        return tx
