#!/bin/bash

go get github.com/labstack/echo
go get github.com/lib/pq
go get github.com/labstack/echo/middleware
go get github.com/jmoiron/sqlx
go build -o start-server server/*.go
./start-server
