CREATE TABLE IF NOT EXISTS "client"(
    "id" UUID PRIMARY KEY,
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(13) NOT NULL,
    "driving_license_number" VARCHAR(50) NOT NULL,
    "passport_number" VARCHAR(50) NOT NULL,
    "photo" VARCHAR(50),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "driving_number_given_place" VARCHAR(50) NOT NULL,
    "driving_number_given_date" DATE NOT NULL,
    "driving_number_given_expire" DATE NOT NULL,
    "date_of_birth" TIMESTAMP NOT NULL,
    -- false - not blocked, true - blocked
    "is_blocked" BOOL DEFAULT false,
    "propiska" VARCHAR(50) NOT NULL,
    "passport_pinfl" VARCHAR(14) NOT NULL
);