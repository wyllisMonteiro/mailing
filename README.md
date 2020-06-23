# Creation of API using JWT

Technologies used :
- Golang
- JWT

## Set up project
**Run client**
```sh
$ cd client
$ go run main.go
```

**Run server**
```sh
$ cd server
$ go run main.go
```

**Use postman**
- Put GET http://localhost:9000
- Go in Headers
- Put "Token" as KEY and put token generated on http://localhost:9001 as VALUE
- In Body you should have "Super Secret Information"

Tutorial followed for authentification : https://www.youtube.com/watch?v=-Scg9INymBs&t=906s