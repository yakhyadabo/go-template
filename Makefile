default: build

build: 
	go build deploy.go

rss:
	go run deploy.go

install:	
	go install deploy.go
