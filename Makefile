BINARY_NAME=promova-news

build:
	go build -o ${BINARY_NAME} ./src/main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ./${BINARY_NAME}

test:
	go test ./src/... -v


