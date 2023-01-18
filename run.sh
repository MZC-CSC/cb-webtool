#!/bin/bash
echo "Setup config"
source ./conf/setup.env

#Run server
echo "start server:1235"
go run main.go
