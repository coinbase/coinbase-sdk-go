# Coinbase Go SDK Changelog

## Unreleased

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
