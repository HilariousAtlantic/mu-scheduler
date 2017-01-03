# Schedule-Buddy
A site to simplify the scheduling process, specifically at Miami University.

# Run
Install Go

    brew install go

Install [Glide](https://glide.sh/), a Go dependency manager

    curl https://glide.sh/get | sh

Install Node

    brew install node

Add this to your `.bashrc`
```shell
export PATH=$PATH:/usr/local/opt/go/libexec/bin
export GOPATH=$HOME/golang
export GOROOT=/usr/local/opt/go/libexec
```
then from the root directory

    npm install
    npm run build
    glide install
    go run *.go
