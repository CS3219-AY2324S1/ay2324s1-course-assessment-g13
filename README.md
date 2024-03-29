[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/6BOvYMwN)
# Group 13 

Follow the steps below to set up the entire project using **Docker Compose**.

1. Turn on [Docker Desktop](https://www.docker.com/products/docker-desktop/)
2. Copy the `.env.sample` file into a `.env` and fill in the environment variables with the required secrets (Only the `.env` at root directory is needed, .env in subdirectories are deprecated)
3. Change directory to the `scripts` folder and run `chmod +rwx <script>` for both scripts
4. Ensure that all the directories have their environment variables setup as well by following the respective `README.md`'s in the directories
5. For TA reviewing this assignment or **production** environment, simply run `docker compose -f docker-compose.prod.yaml up --build`. For **development**, run `docker compose up --build`
6. Go to `http://localhost:3000` and interact with Peerprep!

## For development

For development, we are using `Dockerfile.dev` for all backend microservices to
allow for hot-reload using air. Since we do not employ multi-stage builds for 
development environments, the image size is usually much larger than it should be.

## For production 

For production, we are using `Dockerfile.prod` for all services. These files employ
a multistage build process which ensures the final image is lean and only contains
the minimum amount of files to containerise the service.  

## Google Cloud CLI helpful commands
**To use kubectl with a specific project**
```shell
gcloud container clusters get-credentials ${CLUSTER_NAME} \
--project=${PROJECT} \
--region=${REGION}
```
For the purpose of this project, an example of the above command would be
```shell
gcloud container clusters get-credentials peerpreps --project=smooth-research-397708 --region=asia-southeast1
```

Now, you can run commands like:
- `kubectl get nodes`: Gets the nodes available in your cluster
- `kubectl get deployments`: Track status of deployments
- `kubectl get pod`: See pods created by deployment
- `kubectl get services`: Get all services

**Gets the project ID**
```shell
gcloud config get-value project
```

## FAQ / Troubleshooting
1. Running `docker compose up --build` returns a "database service unhealthy" error and terminates.
    > Simply rerun the command and this issue should go away. The reason for this is because the script
   > is designed to terminate on error, and database existing is considered an error.
