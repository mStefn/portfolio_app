#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE TABLE IF NOT EXISTS visit_counter (
        id SERIAL PRIMARY KEY,
        visits INT DEFAULT 0
    );
    INSERT INTO visit_counter (id, visits) VALUES (1, 0) ON CONFLICT (id) DO NOTHING;
EOSQL

