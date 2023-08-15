CREATE DATABASE transfers;

\c transfers;

CREATE TABLE "user" (
    "ID" INT NOT NULL,
    "Nome" VARCHAR(20) NOT NULL,
    "Saldo" INT NOT NULL,
    CONSTRAINT "PK_User" PRIMARY KEY ("ID")
);