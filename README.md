# Mailing

Technologies used :
- Golang
- JWT
- RabbitMQ
- Rest

## Set up project
Set up environment variables

Create a .env file and add following vars with your data

First there are vars for db
Then vars for RabbitMQ

```
USERDB=root
PASSDB=root
IPDB=db:3306
NAMEDB=mailing

HOSTMAIL=smtp.mailtrap.io
PORTMAIL=587
USERMAIL=ed8a72da8c6188
PASSMAIL=595bc17392bcb0
FROMMAIL=353cbdd366-e809e4@inbox.mailtrap.io
TOMAIL=353cbdd366-e809e4@inbox.mailtrap.io
```

```sh
$ docker-compose up --build
```

Replace "root" (it's username) and "root" (it's password) by your own login

```sh
$ docker exec -i db sh -c 'exec mysql -uroot -p"root"' < ./migrations/mailing.sql
```

Check API host
```sh
$ docker ps
```


## API Routes


	r.HandleFunc("/campaigns", controllers.CreateCampaign).Methods("POST")
	r.HandleFunc("/campaigns/{id}", controllers.GetCampaign).Methods("GET")
	r.HandleFunc("/broadcasts", controllers.CreateBroadcast).Methods("POST")
	r.HandleFunc("/broadcasts", controllers.GetBroadcast).Queries("name", "{name}").Methods("GET")
	r.HandleFunc("/broadcasts/subscriber", controllers.AddSubscriber).Methods("POST")
	r.HandleFunc("/broadcasts/subscriber", controllers.DeleteSubscriber).Methods("DELETE")

```
Login route
POST   /api/login
```

```
Add a new campaign
POST   /api/campaigns
```

```
Get a specific campaign
GET    /api/campaigns/{id}
```

```
Add a new broadcast
POST    /api/broadcasts              
```

```
Get a specific broadcast (by name)
GET    /api/broadcasts/{name}
```

```http
Add a new subscriber
POST    /broadcasts/subscriber       
```

```http
Delete a subscriber
DELETE   /broadcasts/subscriber        
```

## Unit tests
You have to add your unit tests by using `_test` suffix for example `auth_test.go`allow us to test `auth.go` file

## Next Step
- Add all unit tests
- Make some refactoring
- Include token in each request
- Add some routes like DELETE /broadcast/{id}, POST /client

## Tutorials sources

Authentification : https://www.youtube.com/watch?v=-Scg9INymBs&t=906s  
DB use & creation : https://tutorialedge.net/golang/golang-mysql-tutorial/  
Hashing : https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go  
