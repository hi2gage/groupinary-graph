-- Modify "definitions" table
ALTER TABLE "definitions" DROP CONSTRAINT "definitions_words_definitions", ADD CONSTRAINT "definitions_words_definitions" FOREIGN KEY ("word_definitions") REFERENCES "words" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
