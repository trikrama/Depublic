BEGIN;

ALTER TABLE 
    "public"."users" 
ADD 
    COLUMN "number" varchar(255) COLLATE "pg_catalog"."default";

COMMIT;