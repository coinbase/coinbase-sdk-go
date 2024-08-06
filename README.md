# Coinbase Go SDK


The Coinbase Go SDK enables the simple integration of crypto into your app. By calling Coinbase's Platform APIs, the SDK allows you to provision crypto wallets, send crypto into/out of those wallets, track wallet balances, and trade crypto from one asset into another.

The SDK currently supports Customer-custodied Wallets on the Base Sepolia test network.

**NOTE: The Coinbase SDK is currently in Alpha. The SDK:**

- **may make backwards-incompatible changes between releases**
- **should not be used on Mainnet (i.e. with real funds)**

## Documentation

- [Platform API Documentation](https://docs.cdp.coinbase.com/platform-apis/docs/welcome)

## Requirements

The Coinbase server-side SDK requires Go version 1.21 or higher. To view your currently installed versions of Go, run the following from the command-line:

```bash
go version
```

## Installation

Install the latest version:

- [Go latest](https://go.dev/doc/install)

#### You can install the SDK as follows
```bash
go install github.com/coinbase/coinbase-sdk-go
```

