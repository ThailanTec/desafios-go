CREATE DATABASE urubu;

\c urubu;

CREATE TABLE "user" (
    "ID" INT NOT NULL,
    "nome" VARCHAR(20) NOT NULL,
    "email" VARCHAR(20) NOT NULL,
    "cpf" INT NOT NULL,
    "birthday" VARCHAR(12) NOT NULL,
    "money" int NOT NULL,
    CONSTRAINT "PK_user" PRIMARY KEY ("ID")
);