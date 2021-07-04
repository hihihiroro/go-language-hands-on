CREATE TABLE "posts" (
    "id" integer PRIMARY KEY AUTOINCREMENT,
    "created_at" datetime,
    "updated_at" datetime,
    "deleted_at" datetime,
    "address" varchar(255),
    "message" varchar(255),
    "user_id" integer,
    "group_id" integer
);
