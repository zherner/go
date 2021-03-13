.PHONY: pinky
.DEFAULT_GOAL := pinky

pinky :
	go build -v ./...

clean :
	rm -rf ./pinky