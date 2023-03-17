CREATE TABLE IF NOT EXISTS "accounts" (
  "id" bigserial PRIMARY KEY,
  "document_number" bigint NOT NULL UNIQUE,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "operation_types" (
  "id" bigserial PRIMARY KEY,
  "multiplier" int NOT NULL,
  "description" varchar(50) NOT NULL UNIQUE,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "transactions" (
  "id" bigserial PRIMARY KEY,
  "account_id" int NOT NULL,
  "operation_type_id" int NOT NULL,
  "amount" FLOAT NOT NULL,
  "event_date" timestamp DEFAULT (now()),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("operation_type_id") REFERENCES "operation_types" ("id");

