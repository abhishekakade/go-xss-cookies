run: bin/xss-store-cookies
	@PATH="$(PWD)/bin:$(PATH)" heroku local

bin/xss-store-cookies: main.go
	go build -o bin/xss-store-cookies main.go