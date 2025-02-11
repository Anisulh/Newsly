.PHONY: install-deps
install-deps:
	@echo "Installing dependencies..."
	go mod tidy
	cd ./web && npm install
	go install github.com/a-h/templ/cmd/templ@latest

.PHONY: tailwind-watch
tailwind-watch:
	npx @tailwindcss/cli -i ./web/static/css/index.css -o ./web/static/css/output.css --watch

.PHONY: tailwind-build
tailwind-build:
	@echo "Building tailwindcss..."
	npx @tailwindcss/cli -i ./web/static/css/index.css -o ./web/static/css/style.min.css --minify

.PHONY: templ-generate
templ-generate:
	@echo "Generating templates..."
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch
	
.PHONY: dev
dev:
	@echo "Building application..."
	go build -o ./tmp/$(APP_NAME) ./cmd/main.go
	@echo "Starting development watchers..."
	# Run air in the background
	air & \
	# Run templ watcher in the background
	templ generate --watch & \
	# Run tailwind watcher in the background
	npx @tailwindcss/cli -i ./web/static/css/index.css -o ./web/static/css/output.css --watch & \
	# Wait for all background processes to exit
	wait

.PHONY: build
build: install-deps
	make tailwind-build
	make templ-generate
	@echo "Building application..."
	go build -ldflags "-X main.Environment=production" -o ./bin ./cmd/main.go


.PHONY: prod
prod: build
	@echo "Starting production application..."
	./bin/$(APP_NAME)

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	  go test -race -v -timeout 30s ./...