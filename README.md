## RESTful API using golang

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

### Todo

[ ] Include tests
[ ] Dockerize
[ ] Generate go-report
