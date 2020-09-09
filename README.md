## RESTful API using golang

<br>

[![Go Report Card](https://www.goreportcard.com/badge/github.com/YashMeh/golang-restAPI)](https://www.goreportcard.com/badge/github.com/YashMeh/golang-restAPI)

<br>
This is a very simple REST api, made entirely for the sake of learning coding patterns in golang.Any suggestions are welcome :smiley:

### Endpoints

| Method | Endpoint     | Description                        |
| ------ | ------------ | ---------------------------------- |
| GET    | /api/books   | returns the list of all books      |
| GET    | /api/book/id | returns the book with the given id |
| POST   | /api/books   | creates a book                     |
| PUT    | /api/book/id | updates the book with the given id |
| DELETE | /api/book/id | deletes the book with the given id |

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

- Start the API using

```
go build -o restapi && ./restapi
```

- Run the tests

```
go test
```

### Todo

- [x] Include route tests
- [ ] Include integration tests
- [x] Dockerize
- [ ] Generate go-report
- [ ] Create CI

### References

- [Clean Architecture by Angad Sharma](https://github.com/L04DB4L4NC3R/clean-architecture-sample)

- [Full stack epic part - 1](https://github.com/morris-ribs/musicstore)
