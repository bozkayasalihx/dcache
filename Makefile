build: 
	go build -o bin/sbcache

run: build 
	bin/sbcache -a :3000 -la :4000

runtest: 
	go run test/tester.go 