CREATE TABLE "users" (
    "id" BIGSERIAL NOT NULL,
    "username" TEXT NOT NULL UNIQUE,
    "email" TEXT NOT NULL UNIQUE,
    "salted_hash" TEXT NOT NULL,
    "firstname" TEXT NOT NULL,
    "lastname" TEXT NOT NULL,
    "is_admin" BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY ("id")
);

CREATE INDEX "idx_users_id" ON "users" ("id");

CREATE TABLE "sessions" (
    "id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "expires_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    "user_id" BIGINT NOT NULL,
    PRIMARY KEY ("id")
);

CREATE INDEX "idx_sessions_id" ON "sessions" ("id");
CREATE INDEX "idx_sessions_expires_at" ON "sessions" ("expires_at");
