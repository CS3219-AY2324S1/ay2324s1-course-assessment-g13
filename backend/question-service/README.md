### Context
Since we're not publishing our docker image on the docker registry, each developer has to build
their own docker image from scratch before running it locally. Here are the steps for you to build
your own docker image and then run it on port 8080. 
### Running Docker on your machine

 1. To build your own docker image, run the following command 
    ```bash
    docker build --tag question-service .
    ```

 2. Alternatively, build your image using multistage builds to build a leaner image binary 
     ```bash
<<<<<<< HEAD
     docker build -t question-service:multistage -f Dockerfile.stage.stage.stage.multistage .
=======
     docker build -t question-service:multistage -f Dockerfile.prod .
>>>>>>> 7bb1470 (Set up containerisation for question-service)
     ```

 3. Once the image is finished building, run the following command to verify
    ```bash
    docker image ls
    ```
    
 4. Next, run the image with port 8080 published on our local network
     ```bash
     docker run --publish 8080:8080 question-service
     ```
    To run in detached mode, simply add a `--detach` or `-d` flag while running `docker run`.
    To view running detached containers, run `docker ps` to see a list of containers running on your machine. 
    <br><br>

 5. To remove any docker images, run the following command:
    ```bash
    docker rm <container_name>
    ```
    
Set your `.env` file based on the `.env.sample` for MONGODB_URI variable. For local testing, this should default to
"mongodb://mongodb:27017"