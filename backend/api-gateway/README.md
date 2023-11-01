# Question Service

## Setting up environment

1. Create a `.env` file referencing from the `.env.sample` in the current directory

2. Here's an example of the `.env` file:

```
PGUSER="postgres"
POSTGRES_PASSWORD="postgres"
POSTGRES_DB="apigateway"

ACCESS_TOKEN_SECRET_KEY=super_secret_key
REFRESH_TOKEN_SECRET_KEY=super_secret_key

USER_SERVICE_URL=http://user-service:8080
QUESTION_SERVICE_URL=http://question-service:8080
COLLAB_SERVICE_URL=http://collaboration-service:8080
MATCHING_SERVICE_URL=http://matching-producer:8080
AGW_URL=http://localhost:1234
FRONTEND_URL=http://localhost:3000

GITHUB_CLIENT_ID=<details omitted>
GITHUB_CLIENT_SECRET=<details omitted>
```

For `GITHUB_CLIENT_ID` and `GITHUB_CLIENT_SECRET`, please refer to the assignment submission folder.
