CREATE TYPE "habit_type" AS ENUM (
  'daily',
  'weekday',
  'weekend',
);

CREATE TABLE "user" (
  "id" bigint,
  "username" varchar,
  "password" varchar,
  "habit_id" bigint
);

CREATE TABLE "habit" (
  "id" bigint PRIMARY KEY,
  "habit_type" enum,
  "description" varchar
);

ALTER TABLE "habit" ADD FOREIGN KEY ("id") REFERENCES "user" ("id");
