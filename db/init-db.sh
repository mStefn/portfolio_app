#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE TABLE IF NOT EXISTS visit_counter (
        id SERIAL PRIMARY KEY,
        visits INTEGER NOT NULL
    );

    INSERT INTO visit_counter (visits) VALUES (0);
EOSQL

