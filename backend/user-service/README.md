# User Service

## Development

1. Create a `.env` file referencing from the `.env.sample` in the current directory

2. Run `docker compose up --build` in the current directory

## Commands

To check what is inside the database, run `docker exec -it postgres psql -U <POSTGRES_USER> <POSTGRES_DB>`

To use pg_admin online, visit `localhost:5050` after running `docker compose up --build`. The email and password is the default in pg-admin environment
