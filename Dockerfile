FROM golang:alpine

WORKDIR /go/src/github.com/justym/football-api
COPY . .

ENV GO111MODULE=on
RUN go mod download
#RUN go build -v -o ./bin/api ./api/cmd/football-api-server/main.go

EXPOSE 8000
CMD [ "go","run","./api/cmd/football-api-server/main.go","--port=8000"]







