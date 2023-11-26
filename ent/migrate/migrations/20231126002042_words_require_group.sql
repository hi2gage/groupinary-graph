-- Modify "words" table
ALTER TABLE "words" DROP CONSTRAINT "words_groups_rootWords", ADD COLUMN "group_words" bigint NOT NULL, ADD CONSTRAINT "words_groups_words" FOREIGN KEY ("group_words") REFERENCES "groups" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
