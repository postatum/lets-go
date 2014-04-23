MAIN=main.go

install:
	go install views
	go install utils
	go install api

run: install
	go run $(MAIN)
