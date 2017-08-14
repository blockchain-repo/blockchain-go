# Backend Interfaces

## Structure

- [`changefeed.go`](./changefeed.py): Changefeed-related interfaces
- [`connection.go`](./connection.py): Database connection-related interfaces
- [`query.go`](./query.py): Database query-related interfaces, dispatched through single-dispatch
- [`schema.go`](./schema.py): Database setup and schema-related interfaces, dispatched through
  single-dispatch

Built-in implementations (e.g. [RethinkDB's](./rethinkdb)) are provided in sub-directories and
have their connection type's location exposed as `BACKENDS` in [`connection.go`](./connection.go).