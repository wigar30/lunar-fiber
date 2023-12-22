migrate:
	go run ./pkg/migrations/main.go -method=up

migrate-down:
	go run ./pkg/migrations/main.go -method=down

migrate-fresh:
	go run ./pkg/migrations/main.go -method=fresh

migrate-seed:
	go run ./pkg/seeds/main.go

wire:
	cd internal/app && go generate

dev:
	. ${HOME}/.nvm/nvm.sh && nvm use 18 && npx nodemon --exec go run ./cmd/app/main.go --signal SIGTERM