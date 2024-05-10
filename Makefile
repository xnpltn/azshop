run: 
	go build -o bin/app cmd/main.go && ./bin/app

ts:
	tailwindcss --config configs/tailwind.config.js \
		-i configs/input.css \
		-o assets/css/styles.css \
		-w \
		--minify

templ:
	templ generate -watch -proxy=http://localhost:6969

goose-up:
	cd sql/schema; goose postgres postgresql://postgres:postgres@localhost:5433/azshop up


goose-down:
	cd sql/schema ; goose postgres postgresql://postgres:postgres@localhost:5433/azshop down

sqlc:
	sqlc generate
