build:
	mkdir -p bin/
	go build -v -o bin/sesame main.go
install: bin/sesame
	install bin/sesame $(PREFIX)/bin/
clean:
	rm -rf bin/
