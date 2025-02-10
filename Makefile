APP_NAME=go-webapp

.PHONY: install-deps
install-deps:
	npm install

.PHONY: tailwind-watch
tailwind-watch:
	npx @tailwindcss/cli -i ./web/static/css/index.css -o ./web/static/css/output.css --watch

.PHONY: tailwind-build
tailwind-build:
	npx @tailwindcss/cli -i ./web/static/css/index.css -o ./web/static/css/style.min.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch
	
.PHONY: dev
dev:
	@echo "Building application..."
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go
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
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go


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