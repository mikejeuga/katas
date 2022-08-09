repo=$(shell basename "`pwd`")
gopher:
	@git init
	@go mod init github.com/mikejeuga/$(repo)
	@go mod tidy
	@go install github.com/alecthomas/assert/v2
	@go install github.com/matryer/moq@latest


t: test
test:
	@go test -v ./...


c: commit
commit:
	@git commit -am "$m"
	@git pull --rebase
	git push
