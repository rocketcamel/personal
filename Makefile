TAILWIND_INPUT = static/css/input.css
TAILWIND_OUTPUT = static/css/output.css
BUILD_BIN = tmp/bin/main
BUILD_CMD = go build -o $(BUILD_BIN) cmd/main.go

templ-dev:
	templ generate --watch --proxy="localhost:3000"

templ-prod:
	templ generate

tailwind-dev:
	tailwindcss -i $(TAILWIND_INPUT) -o $(TAILWIND_OUTPUT) --watch

tailwind-prod:
	tailwindcss -i $(TAILWIND_INPUT) -o $(TAILWIND_OUTPUT)

esbuild-dev:
	bun bundle

esbuild-prod:
	bun bundle --prod

server:
	air \
		--build.cmd "$(BUILD_CMD)" \
		--build.bin "$(BUILD_BIN)" \
		--build.delay "100" \
		--build.exclude_dir "node_modules" \
		--build.include_ext "go" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true

dev:
	make -j4 tailwind-dev templ-dev esbuild-dev server

prod:
	make tailwind-prod templ-prod esbuild-prod
	$(BUILD_CMD)

