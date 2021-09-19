# circuit-breaker-go

### How to play

1. Run 3rd-api
`go run cmd/3rd-api/main.go`
2. Run server
`go run cmd/server/main.go `
3. Server get data from 3rd-api
```
curl --location --request GET 'http://localhost:8080' 
```
4. Toggle 3rd-api to return 500
```
curl --location --request POST 'http://localhost:3000/toggle'
```
5. Try to get data from 3rd-api from server
