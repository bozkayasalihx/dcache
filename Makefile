build: 
	go build -o bin/sbcache

run: build 
	bin/sbcache

runtest: 
	go run test/tester.go