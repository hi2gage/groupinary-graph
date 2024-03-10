-- atlas:txtar

-- checks/destructive.sql --
-- atlas:assert DS103
SELECT NOT EXISTS (SELECT 1 FROM "words" WHERE "group_root_words" IS NOT NULL) AS "is_empty";

-- migration.sql --
-- Modify "words" table
ALTER TABLE "words" DROP COLUMN "group_root_words";
