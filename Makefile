.PHONY: push_images build_images images deploy_env_vars deploy_backend_services deploy_frontend_services deploy_all

# Build all microservice images with the v2 tag based on prod Dockerfiles
build_images:
	docker build --platform linux/x86-64 -f ./frontend/Dockerfile.prod -t gcr.io/smooth-research-397708/frontend:v2 ./frontend
	docker build --platform linux/x86-64 -f ./backend/api-gateway/Dockerfile.prod -t gcr.io/smooth-research-397708/api-gateway:v2 ./backend/api-gateway
	docker build --platform linux/x86-64 -f ./backend/collaboration-service/Dockerfile.prod -t gcr.io/smooth-research-397708/collaboration-service:v2 ./backend/collaboration-service
	docker build --platform linux/x86-64 -f ./backend/matching-service/consumer/Dockerfile.prod -t gcr.io/smooth-research-397708/matching-consumer:v2 ./backend/matching-service/consumer
	docker build --platform linux/x86-64 -f ./backend/matching-service/producer/Dockerfile.prod -t gcr.io/smooth-research-397708/matching-producer:v2 ./backend/matching-service/producer
	docker build --platform linux/x86-64 -f ./backend/question-service/Dockerfile.prod -t gcr.io/smooth-research-397708/question-service:v2 ./backend/question-service
	docker build --platform linux/x86-64 -f ./backend/user-service/Dockerfile.prod -t gcr.io/smooth-research-397708/user-service:v2	 ./backend/user-service

# Pushes all newly built microservice images onto GCR with the v2 tag
push_images:
	docker push gcr.io/smooth-research-397708/frontend:v2
	docker push gcr.io/smooth-research-397708/api-gateway:v2
	docker push gcr.io/smooth-research-397708/collaboration-service:v2
	docker push gcr.io/smooth-research-397708/matching-consumer:v2
	docker push gcr.io/smooth-research-397708/matching-producer:v2
	docker push gcr.io/smooth-research-397708/question-service:v2
	docker push gcr.io/smooth-research-397708/user-service:v2

# Deploys both configmap and secrets into the current k8s cluster
deploy_env_vars:
	kubectl apply -f ./kubernetes/configmap.yaml
	kubectl apply -f ./kubernetes/secrets.yaml

# Deploys all backend microservices to the current k8s cluster
deploy_backend_services:
	kubectl apply -f ./kubernetes/api-gateway.yaml
	kubectl apply -f ./kubernetes/collaboration-service.yaml
	kubectl apply -f ./kubernetes/question-service.yaml
	kubectl apply -f ./kubernetes/matching-consumer.yaml
	kubectl apply -f ./kubernetes/matching-producer.yaml
	kubectl apply -f ./kubernetes/user-service.yaml
	kubectl apply -f ./kubernetes/ingress.yaml

# Deploys all frontend services to the current k8s cluster
deploy_frontend_services:
	kubectl apply -f ./kubernetes/frontend.yaml

# Deploys both frontend and backend microservices to the current k8s cluster
deploy_all: deploy_backend_services | deploy_frontend_services

# Requires gcloud credential, so run:
# 1. gcloud auth login
# 2. gcloud auth configure-docker
# Builds all microservice images, then pushes them to GCR
images: build_images | push_images
