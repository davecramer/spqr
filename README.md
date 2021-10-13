[![Go](https://github.com/pg-sharding/spqr/actions/workflows/build.yaml/badge.svg)](https://github.com/pg-sharding/spqr/actions/workflows/build.yaml)
[![Go](https://github.com/pg-sharding/spqr/actions/workflows/tests.yaml/badge.svg)](https://github.com/pg-sharding/spqr/actions/workflows/tests.yaml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pg-sharding/spqr)
![Go Report](https://goreportcard.com/badge/github.com/pg-sharding/spqr)
![Lines of code](https://img.shields.io/tokei/lines/github/pg-sharding/spqr)

# Stateless Postgres Query Router

## Development

How to build

```
make
make build
```

Check it out

```
make run

.......


root@spqr_client_1_1:/# connect.sh
psql (13.4 (Ubuntu 13.4-1.pgdg20.04+1), server lolkekcheburek)
Type "help" for help.

db1=?> \q
root@spqr_client_1_1:/#

```
