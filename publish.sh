#!/bin/sh
MACHINE=reapk
BIN="biubon"
FILE="controllers models views $BIN"
REMOTE_DIR=/home/ali/demo/iand

[ -f main.go ] || (echo "main.go not found" && exit)

echo "go build -o biubon main.go"
go build -o biubon main.go

gcutil ssh $MACHINE sudo supervisorctl stop $BIN:*

gcutil push $MACHINE $FILE $REMOTE_DIR

gcutil ssh $MACHINE sudo supervisorctl restart $BIN:*
