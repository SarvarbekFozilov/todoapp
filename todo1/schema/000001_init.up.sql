CREATE TABLE users
(
  "id"            serial    PRIMARY KEY,
  "username"     varchar(255) not null unique,
  "password_hash" varchar(255) not null,
  "photo"         varchar(255) not null,
  "birthday" date NOT NULL,
  "location" text,
  "created_at" timestamp DEFAULT (now()),
  "created_by" integer,
  "updated_at" timestamp,
  "updated_by" integer,
  "deleted_at" timestamp,
  "deleted_by" integer
)
