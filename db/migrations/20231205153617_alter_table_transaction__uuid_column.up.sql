BEGIN;

-- Tambahkan EXTENSION "uuid-ossp"
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Buat kolom baru dengan UUID
ALTER TABLE "public"."transactions" ADD COLUMN new_id UUID DEFAULT uuid_generate_v4() NOT NULL;

-- Isi kolom baru dengan nilai UUID yang dihasilkan dari kolom "id" yang ada
UPDATE "public"."transactions" SET new_id = uuid_generate_v4();

-- Tentukan kolom "new_id" sebagai primary key
ALTER TABLE "public"."transactions" DROP CONSTRAINT transactions_pkey;
ALTER TABLE "public"."transactions" ADD PRIMARY KEY (new_id);

-- Hapus kolom "id" yang lama
ALTER TABLE "public"."transactions" DROP COLUMN id;

-- Ubah nama kolom "new_id" menjadi "id"
ALTER TABLE "public"."transactions" RENAME COLUMN new_id TO id;

COMMIT;