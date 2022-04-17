CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "User"
(
    "Id"          uuid                DEFAULT uuid_generate_v4(),
    "UserName"    varchar     NOT NULL,
    "Password"    varchar     NOT NULL,
    "TicketCount" integer,
    "CreatedAt"   timestamptz NOT NULL DEFAULT (now()),
    "UpdatedAt"   timestamptz NOT NULL DEFAULT (now()),
    PRIMARY KEY ("Id")
);

CREATE TABLE "Ticket"
(
    "Id"        uuid                 DEFAULT uuid_generate_v4(),
    "UserId"    uuid     NOT NULL,
    "Like"      boolean,
    "Subject"   varchar     NOT NULL,
    "Message"   text,
    "Image"     varchar,
    "CreatedAt" timestamptz NOT NULL DEFAULT (now()),
    PRIMARY KEY ("Id")

);

CREATE TABLE "Reply"
(
    "Id"        uuid                 DEFAULT uuid_generate_v4(),
    "TicketId"  uuid     NOT NULL,
    "UserId"    uuid     NOT NULL,
    "Message"   text,
    "CreatedAt" timestamptz NOT NULL DEFAULT (now()),
    PRIMARY KEY ("Id")
);

ALTER TABLE "Ticket"
    ADD FOREIGN KEY ("UserId") REFERENCES "User" ("Id");

ALTER TABLE "Reply"
    ADD FOREIGN KEY ("TicketId") REFERENCES "Ticket" ("Id");

ALTER TABLE "Reply"
    ADD FOREIGN KEY ("UserId") REFERENCES "User" ("Id");

CREATE INDEX ON "User" ("UserName");

CREATE INDEX ON "Ticket" ("UserId");

CREATE INDEX ON "Reply" ("UserId");



