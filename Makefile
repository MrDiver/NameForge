cli: cmd/cli/main.go
	go build -ldflags "-w -s" -o bin/nameforge cmd/cli/main.go

.PHONY: clean

clean:
	rm -r bin
