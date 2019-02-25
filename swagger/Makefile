gen:
	@make clean
	swagger generate server -f ./src/swagger.yml -t ./gen --exclude-main

clean:
	rm -rf gen/*

.PHONY: gen
