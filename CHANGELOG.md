# Coinbase Go SDK Changelog

## Unreleased

## [0.0.14] - 2024-11-26

### Added

- Add a `Content` and `ToAddressID` method to `Transaction`.

## [0.0.13] - 2024-11-21

### Added

- Update SDK with latest generated client code. 

## [0.0.12] - 2024-09-30

### Added

- Added constants for supported network names
- Added `IsFailedState` and `IsCompleteState` methods to `StakingOperation` to check if the operation is in a failed or complete state

### Changes

- Exposed `IsTerminalState` method on `StakingOperation` to check if the operation is in a terminal state.

## [0.0.11] - 2024-09-23

- Fix a bug where correlation id returned from backend wasn't being relayed to the user
- Add FetchStakingOperation to fetch a staking operation by networkID, addressID and stakingOperationID
- Add ReloadStakingOperation to reload a given staking operation with latest data from the backend

## [0.0.10] - 2024-09-20

- Add Solana signing functionality
- Add and refactor examples.

## [0.0.9] - 2024-09-17

- Update SDK with latest generated client code.

## [0.0.8] - 2024-09-17

- StakingRewards's Date field is now a full timestamp
- StakingBalance's Date field is now a full timestamp
- Empty balance values default to 0

## [0.0.7] - 2024-09-12

- Adds headers to requests with SDK version & language

## [0.0.6] - 2024-09-12

### Added

- Return correlation ID from APIError response

## [0.0.5] - 2024-08-29

### Added

- Allow for use with directly set api key
- Return user facing error type APIError for a server side error

## [0.0.4] - 2024-08-28

### Added

- Add user facing validator statuses

## [0.0.3] - 2024-08-27

### Fixed

- Fixed a bug where we weren't handling the api returned validators properly resulting in an index out of range error.

## [0.0.2] - 2024-08-27

### Fixed

- Fixed a bug where we weren't handling the api returned validators properly resulting in an index out of range error.

## [0.0.1] - 2024-08-23

### Added

Initial release of the Coinbase Go SDK

- Support for Staking on External Wallets
  - Full support for Shared ETH Staking
  - Partial support for Dedicated ETH Staking
    - Only stake is supported, unstake will be coming soon
  - On networks `holesky` and `mainnet`
- Support for getting stakeable balances on External Wallets
- Support for Listing Rewards on External Wallets
