test:
	@gom test

dev:
	@gom build

build:
	GOOS=linux GOARCH=amd64 gom build

dist:
	@rake dist

