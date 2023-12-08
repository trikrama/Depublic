BEGIN;

-- Create transactions table
CREATE TABLE IF NOT EXISTS "public"."transactions" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "user_id" INT NOT NULL,
    "event_id" INT NOT NULL,
    "transaction_status" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "quantity" INT NOT NULL,
    "total" INT NOT NULL,
    "created_at" timestamptz(6) NOT NULL,
    "updated_at" timestamptz(6) NOT NULL,
    "deleted_at" timestamptz(6)
);

COMMIT;