#！/bin/bash

GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`
gofmt -s -w ${GOFILES}
git add * && git commit -a  -m  'update' && git push origin master
