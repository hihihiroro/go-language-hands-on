CREATE TABLE "users" (
    "id" integer PRIMARY KEY AUTOINCREMENT,
    "created_at" datetime,
    "updated_at" datetime,
    "deleted_at" datetime,
    "account" varchar(255),
    "name" varchar(255),
    "password" varchar(255), 
    "message" varchar(255)
);
