build:
	go build -o solve 

test:
	go test ./... 

run: build
	./solve

clean:
	rm solve