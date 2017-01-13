# mu-scheduler

A web application designed to streamline the course registration process at Miami University - Oxford.

## Installation

Install [Go](https://golang.org/), [PostgreSQL](https://www.postgresql.org), and [Node](https://nodejs.org/)
```
brew install go postgres node
```
Add these lines to your `.bashrc`
```shell
export PATH=$PATH:/usr/local/opt/go/libexec/bin
export GOPATH=$HOME/golang
export GOROOT=/usr/local/opt/go/libexec
```
Run these commands from the project root directory
```
npm run database
npm run setup
npm run serve
```

## Server

To build the server (command must be ran from the project root directory)
```
npm run build-server
```

## Database

To create the database (command must be ran from the project root directory)
```
npm run createdb
```
To delete the database (command must be ran from the project root directory)
```
npm run deletedb
```
To delete the database and create a new one (command must be ran from the project root directory)
```
npm run resetdb
```
To interact with the database
```
npm run opendb
```
