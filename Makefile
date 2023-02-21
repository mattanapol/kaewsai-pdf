.PHONY: wk docker-build-wk docker-run-wk

wk: docker-build-wk

docker-build-wk:
	docker build -t kaewsai-wkhtmltopdf-app -f ./deployment/Dockerfile.wkhtmltopdf-app .

docker-run-wk:
	docker run -v "$(pwd)/output":/output kaewsai-wkhtmltopdf-app

api: docker-build-api

docker-build-api:
	docker build -t kaewsai-pdf-api -f ./deployment/Dockerfile.api .

docker-run-api:
	docker run -p 8080:8080 kaewsai-pdf-api