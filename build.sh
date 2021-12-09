echo "building binary for windows amd64"
GOOS=windows GOARCH=amd64 go build -o ./bin/lessonend-win-amd64.exe main.go
echo "building binary for darwin amd64"
GOOS=darwin GOARCH=amd64 go build -o ./bin/lessonend-mac-amd64 main.go
echo "building binary for linux amd64"
GOOS=linux GOARCH=amd64 go build -o ./bin/lessonend-linux-amd64 main.go
echo "building binary for windows arm"
GOOS=windows GOARCH=arm go build -o ./bin/lessonend-win-arm.exe main.go
echo "building binary for darwin arm64"
GOOS=darwin GOARCH=arm64 go build -o ./bin/lessonend-mac-arm64 main.go
echo "building binary for linux arm"
GOOS=linux GOARCH=arm go build -o ./bin/lessonend-linux-arm main.go
