# Contributing to semrel-api

Thank you for your interest in contributing to semrel-api!

## How to Contribute

1. Fork the repository and create a feature branch
2. Make your changes with clear commit messages (Conventional Commits)
3. Add or update tests as appropriate
4. Open a pull request against `main`

## Development Setup

```bash
# Install buf for proto code generation
go install github.com/bufbuild/buf/cmd/buf@latest

# Regenerate proto code after proto changes
buf generate

# Build
go build ./...

# Test
go test ./...
```

## Code of Conduct

This project adheres to the [Code of Conduct](CODE_OF_CONDUCT.md).
All participants are expected to uphold this code.

## License

By contributing, you agree that your contributions will be licensed under the Apache-2.0 License.
