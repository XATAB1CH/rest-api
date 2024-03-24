CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   email VARCHAR NOT NULL UNIQUE,
   encrypted_password VARCHAR NOT NULL); 