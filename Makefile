.PHONY: gen
gen: 
	swagger generate server -A football-api -t api/ -f swagger/swagger.yaml

.PHONY: server
serve: 
	go run api/cmd/football-api-server/main.go

.PHONY: build
build:
	go build -v -o ./bin/server ./api/cmd/football-api-server/main.go

.PHONY: vet
vet: 
	go vet -v ./api/... 

.PHONY: clean
clean:
	rm -rf ./bin/server

.PHONY: dcup
dcup:
	docker-compose up 

.PHONY: dcdown
dcdown: 
	docker-compose down

.PHONY: dcexe
dcexe: 
	docker-compose exec mongodb bash


