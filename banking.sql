CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "document_number" bigint UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "operation_types" (
  "id" bigserial PRIMARY KEY,
  "description" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "account_id" int NOT NULL,
  "operation_type_id" int NOT NULL,
  "amount" bigint NOT NULL,
  "event_date" timestamp DEFAULT (now())
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("operation_type_id") REFERENCES "operation_types" ("id");

