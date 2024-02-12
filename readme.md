# Gault

Gault is a CLI app that provides functionality similar to HashiCorp's Vault, with simplified features. It allows users to store and retrieve key-value pairs securely.

## Installation

To install Gault, use the following command:

`go install github.com/mohieey/gault@latest`

## Usage

### Set a Value

To set a value in Gault, use the following command:

`gault set <key> <value> -k=<encryption key>`

Replace <key> and <value> with the key and value you want to set respectively. Provide an encryption key using the -k flag.

### Get a Value

To retrieve a value from Gault, use the following command:

`gault get <key> -k=<encryption key>`

Replace <key> with the key of the value you want to retrieve. Provide the encryption key using the -k flag.

### Clear All Values

To clear all values stored in Gault, use the following command:

`gault clear`

_Note: Gault stores the values in a file called `.gault` in the home directory and `gault clear` removes that file._

## Contributing

If you'd like to contribute to Gault, please fork the repository and submit a pull request. We welcome contributions!

## License

This project is licensed under the MIT License - see the LICENSE file for details.
