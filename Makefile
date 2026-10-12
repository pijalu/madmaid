.PHONY: all install build dev start

all: build

install:
	npm install

build:
	npm run build
	go build -o madmaid .

dev:
	npm run dev

start: build
	./madmaid
