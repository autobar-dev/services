dev: 
	mprocs \
		"make auth-dev" \
		"make currency-dev" \
		"make email-dev" \
		"make emailtemplate-dev" \
		"make module-dev" \
		"make product-dev" \
		"make realtime-dev" \
		"make user-dev" \
		"make wallet-dev"

auth-dev:
	(cd auth/ && go run main.go)

currency-dev:
	(cd currency/ && cargo run)

email-dev:
	(cd email/ && go run main.go)

emailtemplate-dev:
	(cd emailtemplate/ && npm run start-dev)

module-dev:
	(cd module/ && go run main.go)

product-dev:
	(cd product/ && go run main.go)

realtime-dev:
	(cd realtime/ && go run main.go)

user-dev:
	(cd user/ && go run main.go)

wallet-dev:
	(cd wallet/ && go run main.go)
