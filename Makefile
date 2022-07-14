EXEC=wordle

wordle:
	rm -f buid/${EXEC}
	go build -o build/${EXEC} *.go
