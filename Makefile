compile:
	go build -o barley main.go

run:
	go run main.go

linux:
	GOOS=linux GOARCH=amd64 go build -o barley main.go

darwin:
	GOOS=darwin GOARCH=amd64 go build -o barley main.go

brew:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/brew/barley main.go
	GOOS=darwin GOARCH=arm64 go build -o ./bin/brew/barley-silicon main.go
	wait
	tar -czf ./bin/brew/barley.tar.gz ./bin/brew/barley
	tar -czf ./bin/brew/barley-silicon.tar.gz ./bin/brew/barley-silicon

win:
	GOOS=windows GOARCH=amd64 go build -o barley main.go

linux-arm:
	GOOS=linux GOARCH=arm go build -o barley main.go

darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o barley main.go

win-arm:
	GOOS=windows GOARCH=arm go build -o barley main.go

build:
	echo "building binary for windows amd64"
	GOOS=windows GOARCH=amd64 go build -o ./bin/barley-win-amd64.exe main.go
	echo "building binary for darwin amd64"
	GOOS=darwin GOARCH=amd64 go build -o ./bin/barley-mac-amd64 main.go
	echo "building binary for linux amd64"
	GOOS=linux GOARCH=amd64 go build -o ./bin/barley-linux-amd64 main.go
	echo "building binary for windows arm"
	GOOS=windows GOARCH=arm go build -o ./bin/arm/barley-win-arm.exe main.go
	echo "building binary for darwin arm64"
	GOOS=darwin GOARCH=arm64 go build -o ./bin/arm/barley-mac-arm64 main.go
	echo "building binary for linux arm"
	GOOS=linux GOARCH=arm go build -o ./bin/arm/barley-linux-arm main.go

clean:
	rm -rf bin

install:
	sudo go build -o /usr/bin/barley main.go
