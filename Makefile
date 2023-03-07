.PHONY: wk docker-build-wk docker-run-wk

all: api wk chromium

run-all:
	docker compose -f ./deployment/docker-compose.yml up -d --build

wk: docker-build-wk

docker-build-wk:
	docker build -t kaewsai-wkhtmltopdf-app -f ./deployment/Dockerfile.wkhtmltopdf-app .

docker-run-wk:
	docker run -d --env-file ./.env-wk kaewsai-wkhtmltopdf-app

api: docker-build-api

docker-build-api:
	docker build -t kaewsai-pdf-api -f ./deployment/Dockerfile.api .

docker-run-api:
	docker run -d -p 8080:8080 --env-file ./.env-api kaewsai-pdf-api

chromium: docker-build-chromium

docker-build-chromium:
	docker build -t kaewsai-chromium-app -f ./deployment/Dockerfile.chromium-app .

docker-run-chromium:
	docker run -d --env-file ./.env-chromium kaewsai-chromium-app

plot-dependency:
	docker build -t gen-diagram -f ./script/Dockerfile.diagram .
	docker run -v $(PWD):/app gen-diagram -s -novendor -o github.com/mattanapol,./cmd ./cmd/api | dot -Tpng -o ./docs/diagrams/api.png
	docker run -v $(PWD):/app gen-diagram -s -novendor -o github.com/mattanapol,./cmd ./cmd/chromium-app | dot -Tpng -o ./docs/diagrams/chromium-app.png
	docker run -v $(PWD):/app gen-diagram -s -novendor -o github.com/mattanapol,./cmd ./cmd/wkhtmltopdf-app | dot -Tpng -o ./docs/diagrams/wkhtmltopdf-app.png
