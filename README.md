This is an example API project written Go and Fiber Framework. Test codes has been written for all controller functions.
Environment variables are taken from config.yml.

# Used Tech Stacks

1. Fiber Framework
2. Go Channel for evicting cache
3. Postgresql
4. Redis for caching
5. Air for running the app continuously

## Running the app on localhost

```bash
$ go run . dev
```

## Running the app with Docker

```bash
$ docker-compose up -d
```

## Running the app in Vagrant

```bash
$ vagrant up && vagrant ssh
$ cd /vagrant
$ docker-compose up -d
```