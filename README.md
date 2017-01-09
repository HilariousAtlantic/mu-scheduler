# Schedule-Buddy
A site to simplify the scheduling process, specifically at Miami University.

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
Run these commands from the root directory
```
npm run database
npm run setup
npm run serve
```
**Note:** `npm run database` starts PostgreSQL, and should not be ran again unless PostgreSQL is stopped.

## Usage

To rebuild the server
```
npm run build-server
```
Or to build and run the server directly
```
go run server/*.go <args>
```

## Database
(in main directory)

To create the database
```
./sb createdb
```
To delete the database
```
./sb dropdb
```
To delete the database and create a new one
```
./sb resetdb
```
To interact with the database
```
psql -U schedule_buddy
```
To run a query
```
schedule_buddy#=> SELECT * FROM courses;
```
