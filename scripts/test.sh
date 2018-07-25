#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

go get -u github.com/axw/gocov/gocov
go get -u github.com/AlekSi/gocov-xml
go get -u github.com/jstemmer/go-junit-report
go get -u gopkg.in/alecthomas/gometalinter.v2

export GOROOT=$(go env GOROOT)
FILE="${FILE:-0}"

echo
echo "Running gometalinter"
gometalinter.v2 -i > /dev/null 2>&1
if [ ${FILE} -eq 0 ]; then
	gometalinter.v2 --vendor --enable-all --cyclo-over=25 --disable=lll --disable=safesql --disable=dupl --disable=gochecknoinits --disable=gochecknoglobals --deadline=80s --checkstyle --json ./...
else
	gometalinter.v2 --vendor --enable-all --cyclo-over=25 --disable=lll --disable=safesql --disable=dupl --disable=gochecknoinits --disable=gochecknoglobals --deadline=80s --checkstyle ./... | tee /dev/tty > checkstyle-report.xml
fi

echo
export CGO_ENABLED=0
echo "Running tests:"
if [ ${FILE} -eq 0 ]; then
  go test -v ./...
else
  go test -v ./... | tee /dev/tty | go-junit-report > junit-report.xml
fi

echo
echo "Testing coverage"

go test -covermode=count -coverprofile=profile.cov ./...
go tool cover -func profile.cov

if [ ${FILE} -eq 0 ]; then
  rm -f profile.cov
else
  gocov convert profile.cov | gocov-xml > coverage.xml
fi
echo
