.PHONY: all docker-build docker-run

all: docker-build

docker-build:
	docker build -t kaewsai-wkhtmltopdf-app -f ./deployment/Dockerfile.wkhtmltopdf-app .

docker-run:
	docker run -v "$(pwd)/output":/output kaewsai-wkhtmltopdf-app