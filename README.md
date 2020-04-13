# sched

A simple REST API service which was written to explore how web applications can be written in go.
Requires a running postgres, otherwise crashes on start. Assumes a default connection to postgres on localhost, but connectiion string can be supplied on start, see `main.go` and database module.

### Running the project

```bash
go run main.go
```
