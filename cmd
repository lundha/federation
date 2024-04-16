#!/bin/bash

function generate() {
    # contracts | suppliers | users
    if [ -z "$1" ]; then
        echo "Usage: generate [contracts|suppliers|users]"
        return
    fi
	go get github.com/99designs/gqlgen@v0.17.44
    cd $1 && go run github.com/99designs/gqlgen
}

function dev() {
    function cleanup {
        echo "trap. Cleaning up..."
        kill "$ACCOUNTS_PID"
        kill "$PRODUCTS_PID"
        kill "$REVIEWS_PID"
    }
    trap cleanup EXIT

    go build -o /tmp/srv-contracts ./contracts
    go build -o /tmp/srv-suppliers ./suppliers
    go build -o /tmp/srv-users ./users

    /tmp/srv-contracts &
    ACCOUNTS_PID=$!

    /tmp/srv-suppliers &
    PRODUCTS_PID=$!

    /tmp/srv-users &
    REVIEWS_PID=$!

    sleep 1

    node gateway/index.js
}

$@
