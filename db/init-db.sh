#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE TABLE IF NOT EXISTS visits (
        id SERIAL PRIMARY KEY,
        count INTEGER NOT NULL
    );

    INSERT INTO visits (count) VALUES (0);
EOSQL

