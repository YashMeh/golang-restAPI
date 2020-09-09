## RESTful API using golang

This is a very simple REST api, made entirely for the sake of learning coding patterns in golang.Any suggestions are welcome :smiley:

### Setup

- Create postgres table using `table.sql`

- Setup environment variables in the root directory using

```
export DBHOST=127.0.0.1
export DBPORT=5432
export DBUSER=parzival
export DBPASS=yash123
export DBNAME=foo
```

### Running :running:

```
go build -o restapi && ./restapi
```

### Testing :rotating_light:

```
go test
```

### Todo

- [x] Include route tests
- [ ] Include integration tests
- [x] Dockerize
- [ ] Generate go-report
- [x] Create CI
