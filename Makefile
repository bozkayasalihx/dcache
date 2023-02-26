build: 
	go build -o bin/sbcache

run: build 
	bin/sbcache -a :3001 -la :4001

runtest: 
	go run test/tester.go 