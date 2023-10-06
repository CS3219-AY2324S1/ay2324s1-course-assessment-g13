# User Service

## Setting up environment

1. Create a `.env` file referencing from the `.env.sample` in the current directory

2. Here's an example of the `.env` file:

```
PGUSER="postgres"
POSTGRES_PASSWORD="postgres"
POSTGRES_DB="users"
```

## Commands

To check what is inside the database, run `docker exec -it postgres psql -U <POSTGRES_USER> <POSTGRES_DB>`
