# Creation of API using JWT

Technologies used :
- Golang
- JWT

## Set up project
Set up variable environnement

Create a .env file and add following vars with your data

```
USERDB=root
PASSDB=root
IPDB=db:3306
NAMEDB=mailing
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
Coming soon

Tutorial for authentification : https://www.youtube.com/watch?v=-Scg9INymBs&t=906s
Tutorial for database : https://tutorialedge.net/golang/golang-mysql-tutorial/
Tutorial for hash : https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go