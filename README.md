# Benchmarking with `wrk` and `ghz`

This guide provides steps to set up and use `wrk` for HTTP benchmarking and `ghz` for gRPC benchmarking.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Setting Up wrk](#setting-up-wrk)
- [Setting Up ghz](#setting-up-ghz)
- [Running wrk](#running-wrk)
- [Running ghz](#running-ghz)

## Prerequisites

- Ensure Git is installed on your machine.
- Ensure you have a working Go environment (Go 1.11 or higher is recommended).

## Setting Up `wrk`

### For macOS (using Homebrew):
```bash
brew install wrk
```

### For Linux:
```bash
git clone https://github.com/wg/wrk.git
cd wrk
make
sudo cp wrk /usr/local/bin
```

## Setting Up `ghz`

```bash
go get -u github.com/bojand/ghz/cmd/ghz
```

Ensure `$GOPATH/bin` is in your system's `PATH` so you can run `ghz` from any location.

## Running `wrk`

```bash
wrk -t12 -c400 -d30s http://YOUR_ENDPOINT_URL
```

For POST requests, create a Lua script with request details:
```bash
wrk -t12 -c400 -d30s -s path_to_lua_script.lua http://YOUR_ENDPOINT_URL
```

## Running `ghz`

```bash
ghz --insecure \
    --proto /path/to/your/protofile.proto \
    --call your.package.Service/YourRPCMethod \
    --concurrency 50 \
    --total 2000 \
    YOUR_SERVER_ADDRESS:YOUR_PORT
```
