build: clean
	go build
	./modest-go

clean:
	rm -f modest-go

run:
	LD_LIBRARY_PATH="./modest/lib" \
	go run main.go
