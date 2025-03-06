CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    username VARCHAR(250),
    email VARCHAR(250) UNIQUE,
    created TIMESTAMPTZ,
    updated TIMESTAMPTZ,

    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    CONSTRAINT pkey_tbl PRIMARY KEY (id)
);