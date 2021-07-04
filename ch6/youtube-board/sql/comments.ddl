CREATE TABLE "comments" (
    "id" integer PRIMARY KEY AUTOINCREMENT,
    "created_at" datetime,
    "updated_at" datetime,
    "deleted_at" datetime,
    "user_id" integer,
    "post_id"  integer,
    "message" varchar(255)
);
