### How to Run The Project
First start mysql server with docker:
```bash
docker-compose up -d
```

finally run the server: 

```bash
go run server.go
```
Now navigate to https://localhost:8080 you can see graphiql playground and query the graphql server.
