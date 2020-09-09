FROM golang:1.12.0-alpine3.9
# create a working directory
WORKDIR /go/src/app
# add source code
COPY . .
# expose port
EXPOSE 8000
# run main.go
CMD ["go", "run", "main.go"]