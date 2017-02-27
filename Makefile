#!make
include envfile
export $(shell sed 's/=.*//' envfile)

run: build
	./bin/doiici

build:
	go build -o bin/doiici

install:
	glide install
