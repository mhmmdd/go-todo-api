# Todo Api Written with Go

### Air for live reloading
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

## Running the app on localhost
```bash
$ go run . dev
```

## Running the app
```bash
$ docker-compose up -d
```

## Running the app in Vagrant
```bash
$ vagrant up && vagrant ssh
$ cd /vagrant
$ docker-compose up -d
```