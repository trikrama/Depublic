BEGIN;

CREATE TABLE IF NOT EXISTS "public"."events" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "image" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "description" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "category" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "location" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "price" bigint NOT NULL,
    "date_start" timestamptz NOT NULL,
    "date_end" timestamptz NOT NULL,
    "quantity" bigint NOT NULL,
    "status" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamptz
);

COMMIT;