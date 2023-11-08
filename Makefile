.PHONY: push_images build_images images

# Build all microservice images with the v2 tag based on prod Dockerfiles
build_images:
	docker build -f ./frontend/Dockerfile.prod -t gcr.io/smooth-research-397708/frontend:v2 ./frontend
	docker build -f ./backend/api-gateway/Dockerfile.prod -t gcr.io/smooth-research-397708/api-gateway:v2 ./backend/api-gateway
	docker build -f ./backend/collaboration-service/Dockerfile.prod -t gcr.io/smooth-research-397708/collaboration-service:v2 ./backend/collaboration-service
	docker build -f ./backend/matching-service/consumer/Dockerfile.prod -t gcr.io/smooth-research-397708/matching-consumer:v2 ./backend/matching-service/consumer
	docker build -f ./backend/matching-service/producer/Dockerfile.prod -t gcr.io/smooth-research-397708/matching-producer:v2 ./backend/matching-service/producer
	docker build -f ./backend/question-service/Dockerfile.prod -t gcr.io/smooth-research-397708/question-service:v2 ./backend/question-service
	docker build -f ./backend/user-service/Dockerfile.prod -t gcr.io/smooth-research-397708/user-service:v2	 ./backend/user-service

# Pushes all newly built microservice images onto GCR with the v2 tag
push_images:
	docker push gcr.io/smooth-research-397708/frontend:v2
	docker push gcr.io/smooth-research-397708/api-gateway:v2
	docker push gcr.io/smooth-research-397708/collaboration-service:v2
	docker push gcr.io/smooth-research-397708/matching-consumer:v2
	docker push gcr.io/smooth-research-397708/matching-producer:v2
	docker push gcr.io/smooth-research-397708/question-service:v2
	docker push gcr.io/smooth-research-397708/user-service:v2

# Requires gcloud credential, so run:
# 1. gcloud auth login
# 2. gcloud auth configure-docker
# Builds all microservice images, then pushes them to GCR
images: build_images | push_images