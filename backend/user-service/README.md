# User Service

## Development

1. Create a `.env` file referencing from the `.env.sample` in the current directory

2. Run `docker compose up --build` in the current directory

3. In VScode, go to `File` tab, select `Add Folder to Workspace...` and add `/user-service` folder to workspace 

## Commands

To check what is inside the database, run `docker exec -it postgres psql -U <POSTGRES_USER> <POSTGRES_DB>`
