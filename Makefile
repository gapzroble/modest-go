build: clean
	go build
	./modest-go

clean:
	rm -f modest-go

run:
	go run main.go
