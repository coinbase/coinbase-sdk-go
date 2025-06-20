# Coinbase Go SDK Changelog

## [0.0.28] - 2025-06-04

### Removed

- Removed support for testnet Holesky.

## [0.0.27] - 2025-05-09

### Added

- Add support for validator top-ups.

## [0.0.26] - 2025-05-06

### Added

- Add support for execution layer consolidations post Pectra.

## [0.0.25] - 2025-04-25

### Added

- Add helper staking operation option functions for onchain billing such as `WithFeeRecipientAddress` and `WithRewardSplitterPlanID`.

## [0.0.24] - 2025-04-18

### Added

- Add network constant for `Hoodi`.

## [0.0.23] - 2025-04-15

### Added

- Add support for both execution and consensus layer based unstaking.
- Add user facing examples for all new unstaking scenarios.

## [0.0.22] - 2025-03-27

### Added

- Add `ExecutionLayerWithdrawalsOptionBuilder` and `WithExecutionLayerWithdrawals` as an option to allow for native ETH execution layer withdrawals as defined in https://eips.ethereum.org/EIPS/eip-7002.
- Add `Hoodi` network support.

## [0.0.21] - 2025-02-28

### Changed

- Improved precision of `FromAtomicAmount` and `ToAtomicAmount` calculations. Fixed an issue where the amount "32006467556000000000" wei was interpreted as "32.006467555999999999" due to lossy floating point math. Updated precision to 128 bits to resolve this.

## [0.0.20] - 2025-02-27

### Added

- Add `WithdrawalCredentials` function to the `Validator` resource to help users get the withdrawl credentials for a given validator.

## [0.0.19] - 2025-02-11

### Added

- Add `FeeRecipientAddress` and `ForwardedFeeRecipientAddress` functions to the `Validator` resource to help users get the fee_recipient  and forwarded_fee_recipient addresses for a given validator.

## [0.0.18] - 2025-02-10

### Added

- Add `GetPendingClaimableBalance` function to help users get the pending claimable balance for a given address.

## [0.0.17] - 2025-02-06

### Added

- Update auto-generated client code to help relax decoding of server returned jsons to allow new fields to be added without breaking the SDK.

## [0.0.16] - 2025-02-05

### Added

- Add getters for `Validator` object to expose more data to users.
- Add test cases for `Validator` object.
- Update SDK with latest generated client code.

## [0.0.15] - 2024-12-02

### Added

- Add `WithIntegratorContractAddress` to allow for setting an explicit integrator contract address for Shared ETH staking.

### Fixed

- Better error handling between the SDK and the API.

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

### Changed

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
