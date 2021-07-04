CREATE TABLE "groups" (
    "id" integer PRIMARY KEY AUTOINCREMENT,
    "created_at" datetime,
    "updated_at" datetime,
    "deleted_at" datetime,
    "user_id" integer,
    "name" varchar(255),
    "message" varchar(255)
);
