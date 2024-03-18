-- Create "todos" table
CREATE TABLE "public"."todos" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "title" text NULL,
  "status" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_todos_deleted_at" to table: "todos"
CREATE INDEX "idx_todos_deleted_at" ON "public"."todos" ("deleted_at");
