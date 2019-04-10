.PHONY: build clean deploy

build:
	echo "start make build"
	ls bin
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	ls bin
	echo "ending make build"
	
clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
