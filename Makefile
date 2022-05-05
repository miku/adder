addercli:
	go build -o addercli cmd/addercli/main.go

clean:
	rm -f addercli
