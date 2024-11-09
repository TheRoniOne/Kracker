-- Create "sessions" table
CREATE TABLE "sessions" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "expires_at" timestamptz NOT NULL, "user_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- Create index "idx_sessions_expires_at" to table: "sessions"
CREATE INDEX "idx_sessions_expires_at" ON "sessions" ("expires_at");
-- Create index "idx_sessions_id" to table: "sessions"
CREATE INDEX "idx_sessions_id" ON "sessions" ("id");
