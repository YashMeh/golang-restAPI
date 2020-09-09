## RESTful API using golang

This is a very simple REST api, made entirely for the sake of learning coding patterns in golang.Any suggestions are welcome :smiley:

### Setup

- Create postgres table using `table.sql`

- Create an `.env` file with the database credentials

```
DBHOST=localhost
DBPORT=5432
DBUSER=parzival
DBPASS=yash123
DBNAME=foo
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
