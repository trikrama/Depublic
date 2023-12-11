BEGIN;

ALTER TABLE
    "public"."users" 
DROP 
    COLUMN IF EXISTS "number";

COMMIT;
