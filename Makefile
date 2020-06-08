gen: 
	swagger generate server -A football-api -t api/ -f swagger/swagger.yaml

serve: 
	go run api/cmd/football-api-server/main.go