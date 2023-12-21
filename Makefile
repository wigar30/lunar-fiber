migrate:
	go run ./package/migrations/main.go -method=up

migrate-down:
	go run ./package/migrations/main.go -method=down

migrate-fresh:
	go run ./package/migrations/main.go -method=fresh

migrate-seed:
	go run ./package/seeds/main.go

wire:
	cd internal/app/config && go generate

dev:
	. ${HOME}/.nvm/nvm.sh && nvm use 18 && npx nodemon --exec go run ./cmd/app/main.go --signal SIGTERM