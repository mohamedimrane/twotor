run:
	@npx tailwindcss -i ./static/styles/input.css -o ./static/styles/output.css
	@templ generate
	@sqlc generate
	@echo ----------------------------------------------------
	@go run *.go

migrate:
	@migrate -path migrations -database "sqlite3://database.sqlite" -verbose up

migrate-down:
	@migrate -path migrations -database "sqlite3://database.sqlite" -verbose down
