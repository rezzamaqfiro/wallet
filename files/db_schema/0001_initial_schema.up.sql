CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "orders" (
  "invoice" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "total_payment" bigint DEFAULT 0,
  "user_id" uuid DEFAULT uuid_generate_v4(),
  "status" text,
  "created_at" timestamptz,
  "created_by" text,
  "updated_at" timestamptz,
  "updated_by" text,
  "deleted_at" timestamptz,
  "deleted_by" text,
  "deleted" bool
);

CREATE TABLE "users" (
  "user_id"     uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name"        varchar(150),
  "balance"     bigint DEFAULT 0,
  "created_at" timestamptz,
  "created_by" text,
  "updated_at" timestamptz,
  "updated_by" text,
  "deleted_at" timestamptz,
  "deleted_by" text,
  "deleted" bool
);
