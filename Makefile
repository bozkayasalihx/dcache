testrun:
	go run test/test.go

build:
	go build -o bin/dcache

run: build
	./bin/dcache

runfollower: build
	./bin/dcache --listenAddr :4000  --leaderAddr  :3000

