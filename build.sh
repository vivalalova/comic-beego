# bin/zsh
env GOOS=darwin GOARCH=amd64 go build -o comic-go-mac
env GOOS=linux GOARCH=arm go build -o comic-go-linux
env GOOS=windows GOARCH=amd64 go build -o comic-go-win.exe