CREATE OR REPLACE FUNCTION newticket(UserId1 uuid, Subject1 varchar, Message1 varchar, Image1 varchar, Like1 boolean)
    RETURNS TABLE
            (
                Subject varchar,
                Message text,
                Image varchar,
                Like boolean
            )

AS
$$
BEGIN

    RETURN QUERY
        INSERT INTO "Ticket" ("UserId", "Subject", "Message", "Image", "Like")
            VALUES (UserId1, Subject1, Message1, Image1, Like1)
            RETURNING "Subject","Message","Image","Like";

END
$$
    LANGUAGE 'plpgsql';

