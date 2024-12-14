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