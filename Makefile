run:
	@npx tailwindcss -i ./static/styles/input.css -o ./static/styles/output.css
	@templ generate
	@echo ----------------------------------------------------
	@go run *.go

