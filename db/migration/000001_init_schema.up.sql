CREATE TABLE "word" (
  "id" bigserial PRIMARY KEY,
  "word" varchar NOT NULL
);

CREATE TABLE "usages" (
  "id" bigserial PRIMARY KEY,
  "word_id" bigint NOT NULL,
  "description" varchar NOT NULL
);

CREATE TABLE "example" (
  "id" bigserial PRIMARY KEY,
  "usage_id" bigint NOT NULL,
  "sentence" varchar NOT NULL
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "latest_visit" timestamptz NOT NULL DEFAULT (now()),
  "study_amount" int NOT NULL DEFAULT 0,
  "study_goal" int NOT NULL DEFAULT 10,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "writing" (
  "id" bigserial PRIMARY KEY,
  "writing" varchar NOT NULL,
  "usage_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "usages" ADD FOREIGN KEY ("word_id") REFERENCES "word" ("id");

ALTER TABLE "example" ADD FOREIGN KEY ("usage_id") REFERENCES "usages" ("id");

ALTER TABLE "writing" ADD FOREIGN KEY ("usage_id") REFERENCES "usages" ("id");

ALTER TABLE "writing" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "word" ("word");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "writing" ("user_id");

CREATE INDEX ON "writing" ("user_id", "usage_id");
