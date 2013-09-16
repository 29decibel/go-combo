NO_COLOR	= \033[0m
COLOR	 	= \033[32;01m
SUCCESS_COLOR	= \033[35;01m


all: win linux mac
	@echo "Done!"

mac:
	@echo "$(SUCCESS_COLOR)Build for Mac ...$(NO_COLOR)"
	@go build -o ./bin/gocombo.mac ./server/main.go

linux:
	@echo "$(SUCCESS_COLOR)Build for Linux ...$(NO_COLOR)"
	@GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o ./bin/gocombo.linux ./server/main.go

win:
	@echo "$(SUCCESS_COLOR)Build for Windows ...$(NO_COLOR)"
	@GOOS=windows GOARCH=386 go build -o ./bin/gocombo.exe ./server/main.go
