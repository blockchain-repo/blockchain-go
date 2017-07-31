# Unichain Go

A high-level description of the files and subdirectories of Unichain-go.

There are four database tables which underpin Unichain: `backlog`, where incoming transactions are held temporarily until they can be consumed; `blocks`, where blocks of transactions are written permanently; `assets`, to store the assets (schema, indexes, queries read/write); and `votes`, where votes are written permanently.  It is the votes in the `votes` table which must be queried to determine block validity and order.

## Files

### [`main.go`](./main.go)

Contains code for the CLI.

## Folders

### [`config`](./config)

Methods for managing the configuration, including loading configuration files, automatically generating the configuration, and keeping the configuration consistent across instances.

### [`log`](./log)

Logging infrastructure.

### [`common`](./common)

crypto and utils.

### [`core`](./core)

functional interface.

### [`models`](./models)

There are three main kinds: Transaction,Block,Vote.

### [`pipelines`](./pipelines)

Structure and implementation of various subprocesses.

### [`backend`](./backend)

Code for building the database connection, creating indexes, and other database setup tasks.

### [`web`](./web)

Web server and API.

### [`vendor`](./vendor)

Golang vendor tool to copy external source code from GOPATH or remote location to the local vendor folder. 

## Others

### [`CHANGELOG`](./CHANGELOG)

### [`CONTRIBUTING.md`](./CONTRIBUTING.md)

### [`LICENSE`](./LICENSE)

### [`README.md`](./README.md)

### [`VERSION`](./VERSION)

### [`.gitignore`](./.gitignore)

```
unichain-go/
├── backend/
│   ├── mongodb/
│   └── rethinkdb/
├── common/
├── config/
├── core/
├── log/
├── main.go
├── models/
├── pipelines/
└── web/
```