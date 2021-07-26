-- CREATE TYPE "habit_type" AS ENUM (
--   'daily',
--   'weekday',
--   'weekend'
-- );

CREATE TYPE "status" AS ENUM (
  'todo',
  'in_progress',
  'done'
);

CREATE TABLE "user" (
  "id" bigint,
  "username" varchar,
  "password" varchar
);

CREATE TABLE "habit" (
  "id" bigint PRIMARY KEY,
  "user_id" bigint,
  "days" varchar,
  "description" varchar,
  "status" enum
);

ALTER TABLE "habit" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
