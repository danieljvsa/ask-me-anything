CREATE TABLE "rooms" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "messages" (
  "id" bigserial PRIMARY KEY,
  "message" text,
  "user_id" bigserial NOT NULL,
  "parent_id" varchar,
  "likes_count" bigint DEFAULT 0,
  "answered" bool DEFAULT false,
  "room_id" bigserial NOT NULL,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz DEFAULT 'now()'
);

COMMENT ON COLUMN "messages"."message" IS 'Content of the post';

ALTER TABLE "rooms" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");
