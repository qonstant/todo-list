CREATE TABLE IF NOT EXISTS "tasks" (
    "id" bigserial PRIMARY KEY,
    "title" varchar(200) NOT NULL,
    "active_at" timestamptz NOT NULL,
    "done" boolean NOT NULL DEFAULT (false),
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);