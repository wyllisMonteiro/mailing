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

## API DOC
Coming soon

## Unit tests
Coming soon

Tutorial for authentification : https://www.youtube.com/watch?v=-Scg9INymBs&t=906s
Tutorial for database : https://tutorialedge.net/golang/golang-mysql-tutorial/
Tutorial for hash : https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go