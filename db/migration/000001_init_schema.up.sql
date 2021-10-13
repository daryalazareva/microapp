CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "encrypted_password" varchar NOT NULL
);

CREATE INDEX ON "users" ("email");