# ipfs-retriever
This is an application that retrieves files from IPFS network.
It use `go-ipfs` as a library.

Run this to increase the buffer memory, it is required by `go-ipfs`
```
sudo sysctl -w net.core.rmem_max=2500000
```

# Usage
```
# syntax
# go run main.go [list of CIDs]
go run main.go QmW83RqcKdLUsEEwjWQiz9E6jpjfaFcyNKYVpLtqgra5Hm
```
Check downloaded files at `OUTPUT_BASE_PATH` which is set in `.env` file