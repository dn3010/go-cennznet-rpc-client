# Go CENNZnet RPC Client (GCRPC)

[![License: Apache v2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc Reference](https://godoc.org/github.com/centrifuge/go-substrate-rpc-client?status.svg)](https://godoc.org/github.com/centrifuge/go-substrate-rpc-client)
[![Build Status](https://travis-ci.com/centrifuge/go-substrate-rpc-client.svg?branch=master)](https://travis-ci.com/centrifuge/go-substrate-rpc-client)
[![codecov](https://codecov.io/gh/centrifuge/go-substrate-rpc-client/branch/master/graph/badge.svg)](https://codecov.io/gh/centrifuge/go-substrate-rpc-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/centrifuge/go-substrate-rpc-client)](https://goreportcard.com/report/github.com/centrifuge/go-substrate-rpc-client)

CENNZnet RPC client in Go.
This is a fork of [go-substrate-rpc-client](https://github.com/centrifuge/go-substrate-rpc-client) which adds support for making CENNZnet transactions
and is modelled after [cennznet/api.js](https://github.com/cennznet/api.js).

## State

This package is feature complete, but it is relatively new and might still contain bugs. We advice to use it with caution in production. It comes without any warranties, please refer to LICENCE for details.

## Requirements
Substrate Key Management requires `subkey` to be present in your PATH: https://substrate.dev/docs/en/knowledgebase/integrate/subkey

The `subkey` recommended version: https://github.com/paritytech/substrate/releases/tag/v2.0.0-rc6 

## Documentation & Usage Examples

Please refer to https://godoc.org/github.com/centrifuge/go-substrate-rpc-client

## Contributing

1. Install dependencies by running `make` followed by `make install`
1. Run tests `make test`
1. Lint `make lint` (you can use `make lint-fix` to automatically fix issues)

## Run tests in a Docker container against the CENNZnet Docker image

1. Run the docker container `make test-dockerized`

## Run tests locally against the CENNZnet Docker image

1. Start the Substrate Default Docker image: `make run-cennznet-docker`
1. In another terminal, run the tests against that image: `make test`
1. Visit https://polkadot.js.org/apps for inspection

## Run tests locally against any CENNZnet endpoint

1. Set the endpoint: `export RPC_URL="http://example.com:9933"`
1. Run the tests `make test`
