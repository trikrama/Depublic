BEGIN;

CREATE TABLE IF NOT EXISTS "public"."blogs" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "image" VARCHAR(255) NOT NULL,
    "date" VARCHAR(255) NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "body" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ(6) NOT NULL,
    "updated_at" TIMESTAMPTZ(6) NOT NULL,
    "deleted_at" TIMESTAMPTZ(6)
);

COMMIT;