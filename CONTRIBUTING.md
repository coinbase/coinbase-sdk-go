# Contributing Guide

## Development

### Go Version

Developing in this repository requires Go 1.21 or higher.

### Set-up

Clone the repo by running:

```bash
git clone git@github.com:coinbase/coinbase-sdk-go.git
```

### Linting

To autocorrect all lint errors, run:

```bash
make lint-fix
```

To detect all lint errors, run:

```bash
make lint
```

### Testing

To run all tests, run:

```bash
make test
```

### Generating Documentation

To generate documentation from the TypeDoc comments, run:

```bash
make docs
```
