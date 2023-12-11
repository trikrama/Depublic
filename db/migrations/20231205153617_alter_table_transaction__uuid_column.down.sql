BEGIN;


-- Buat kolom baru dengan SERIAL
ALTER TABLE "public"."transactions" ADD COLUMN new_id SERIAL PRIMARY KEY;

-- Isi kolom baru dengan nilai dari kolom "id" yang diubah ke SERIAL
UPDATE "public"."transactions" SET new_id = id;

-- Hapus kolom "id" yang baru
ALTER TABLE "public"."transactions" DROP COLUMN id;

-- Ubah nama kolom "new_id" menjadi "id"
ALTER TABLE "public"."transactions" RENAME COLUMN new_id TO id;

COMMIT;