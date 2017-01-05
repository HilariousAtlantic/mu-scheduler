# Schedule-Buddy
A site to simplify the scheduling process, specifically at Miami University.

## Installation
Install [Go](https://golang.org/), [Glide](https://glide.sh/), [PostgreSQL](https://www.postgresql.org), and [Node](https://nodejs.org/en/)
```
brew install go glide postgres node
```
Then run these commands to start postgres and create a new user
```
brew services start postgresql
createuser schedule_buddy
```
Add these lines to your `.bashrc`
```shell
export PATH=$PATH:/usr/local/opt/go/libexec/bin
export GOPATH=$HOME/golang
export GOROOT=/usr/local/opt/go/libexec
```
Run these commands from the root directory
```
npm run setup
npm run serve
```

## Database
Create the database
```
go run *.go createdb
```
Delete the database
```
go run *.go dropdb
```
Or in one step
```
go run *.go resetdb
```
To interact with the database, run
```
psql -U schedule_buddy
```
Then you may run queries like so:
```
schedule_buddy#=> SELECT * FROM courses;

```
