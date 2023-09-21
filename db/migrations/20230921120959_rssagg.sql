-- migrate:up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    name TEXT NOT NULL
);

-- migrate:down

DROP TABLE users;
