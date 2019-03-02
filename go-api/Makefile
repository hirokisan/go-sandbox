install:
	go install ./...

design:
	mkdir -p ./design/${API_NAME} && touch ./design/${API_NAME}/${API_NAME}.go

api:
	goagen bootstrap -d github.com/hirokisan/go-sandbox/go-api/design/${API_NAME} -o ./${API_NAME}

.PHONY: design
