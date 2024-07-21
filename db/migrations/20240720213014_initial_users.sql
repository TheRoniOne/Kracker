-- Create "users" table
CREATE TABLE "users" ("id" bigserial NOT NULL, "username" text NOT NULL, "email" text NOT NULL, "salted_hash" text NOT NULL, "firstname" text NOT NULL, "lastname" text NOT NULL, PRIMARY KEY ("id"));
-- Create index "idx_id" to table: "users"
CREATE INDEX "idx_id" ON "users" ("id");
