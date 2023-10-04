[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/6BOvYMwN)
# Group 13 

Follow the steps below to set up the entire project using **Docker Compose**.

1. Turn on [Docker Desktop](https://www.docker.com/products/docker-desktop/)
2. Copy the `.env.sample` file into a `.env` and fill in the environment variables with the required secrets
3. Change directory to the `scripts` folder and run `chmod +rwx <script>` for both scripts
4. Ensure that all the repositories have their environment setup as well by following the respective `README.md`'s
5. Once that is ensured, simply run `docker compose up --build`
6. Go to `http://localhost:3000` and interact with Peerprep!

## For development

For development, we are using `Dockerfile.dev` for all backend microservices to
allow for hot-reload using air. Since we do not employ multi-stage builds for 
development environments, the image size is usually much larger than it should be.

## For production [WIP]

For production, we are using `Dockerfile.prod` for all services. These files employ
a multistage build process which ensures the final image is lean and only contains
the minimum amount of files to containerise the service.  

## FAQ / Troubleshooting
1. Running `docker compose up --build` returns a "database service unhealthy" error and terminates.
    > Simply rerun the command and this issue should go away.
