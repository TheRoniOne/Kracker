-- Create "users" table
CREATE TABLE "users" ("id" bigserial NOT NULL, "username" text NOT NULL, "email" text NOT NULL, "salted_hash" text NOT NULL, "firstname" text NOT NULL, "lastname" text NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "users_email_key" UNIQUE ("email"), CONSTRAINT "users_username_key" UNIQUE ("username"));
-- Create index "idx_users_id" to table: "users"
CREATE INDEX "idx_users_id" ON "users" ("id");
