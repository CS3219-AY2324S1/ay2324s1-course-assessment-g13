#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "postgres" --dbname "users" <<-EOSQL
    DROP DATABASE IF EXISTS users;
    CREATE DATABASE users;
EOSQL