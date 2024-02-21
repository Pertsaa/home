BINARY_NAME=home

run:
	./bin/$(BINARY_NAME)

dev:
	./tailwindcss -i style/input.css -o static/css/styles.css
	go build -o bin/$(BINARY_NAME) .

prod:
	./tailwindcss -i style/input.css -o static/css/styles.css --minify
	go build -o bin/$(BINARY_NAME) .
