#### Context
This api-gateway will be single entry point for all other microservices, like user service and question service. 
All endpoints will pass through the api-gateway for token validation before routing to the other services.
This api-gateway also has its own database and implemented Authentication for Peerpreps users. 

#### Setting Up
1. Set up the environment variables as shown from `.env.sample`
    - Make sure that all the service that you want to use has its env file configured.
2. Create a docker network called `peerpreps-backend` in `/backend`
> `docker network create peerpreps-backend`
3. Change directory to `/api-gateway` and run the containter 
> `docker compose up --build`
4. Go to the other service that you want to access to
     - For example, if you want to access the question service, cd to `/question-service` and run
    > `docker compose up --build`

That would be all, and you should be able to use api-gateway to access other end points in other services.
