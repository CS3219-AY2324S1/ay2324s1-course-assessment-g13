#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "postgres" --dbname "apigateway" <<-EOSQL
    DROP DATABASE IF EXISTS apigateway;
    CREATE DATABASE apigateway;
EOSQL
