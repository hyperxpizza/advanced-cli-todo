build:
	go build -o bin/todo ./cmd/advanced-cli-todo/main.go
run:
	go run ./cmd/advanced-cli-todo/main.go --loglevel=debug --config=/home/hyperxpizza/dev/golang/advanced-cli-todo/configs/config.yml --mode=default
run_web:
	go run ./cmd/advanced-cli-todo/main.go --loglevel=debug --config=/home/hyperxpizza/dev/golang/advanced-cli-todo/configs/config.yml --mode=web
run_cli:
	go run ./cmd/advanced-cli-todo/main.go --loglevel=debug --config=/home/hyperxpizza/dev/golang/advanced-cli-todo/configs/config.yml --mode=cli