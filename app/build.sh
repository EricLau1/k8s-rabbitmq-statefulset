#!/bin/bash

go build -o sender sender.go
go build -o receiver receiver.go

docker build -t rmq-app:v1 .

docker images
