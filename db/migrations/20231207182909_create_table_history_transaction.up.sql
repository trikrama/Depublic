BEGIN;

CREATE TABLE IF NOT EXISTS "public"."history_transactions" (
    "history_id" SERIAL NOT NULL PRIMARY KEY,
    "transaction_id" UUID NOT NULL,
    "user_id" INT NOT NULL,
    "action" VARCHAR(50) NOT NULL,
    "timestamp" TIMESTAMP(6) NOT NULL,
    "name_event" VARCHAR(255) NOT NULL,
    "quantity" INT NOT NULL,
    "total" INT NOT NULL
);

COMMIT;