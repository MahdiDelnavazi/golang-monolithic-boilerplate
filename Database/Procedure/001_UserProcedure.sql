
CREATE OR REPLACE FUNCTION checkUserExist(UserName1 varchar(20))
    RETURNS TABLE
            (
                UserName varchar
            )
AS
$$
BEGIN

RETURN QUERY
SELECT "UserName" FROM "User" where "UserName" = UserName1;

END
$$
LANGUAGE 'plpgsql';

-- --------------------------------------------------------


CREATE OR REPLACE FUNCTION getUser(UserName1 varchar(20))
    RETURNS TABLE
            (
                Id uuid,
                UserName varchar
            )
AS
$$
BEGIN

    RETURN QUERY
        SELECT "Id" , "UserName" FROM "User" where "UserName" = UserName1;

END
$$
    LANGUAGE 'plpgsql';

-- --------------------------------------------------------


CREATE OR REPLACE FUNCTION loginUser(UserName1 varchar(20))
    RETURNS TABLE
            (
                Id uuid,
                UserName varchar,
                Password varchar
            )
AS
$$
BEGIN

    RETURN QUERY
        SELECT "Id" , "UserName" , "Password" FROM "User" where "UserName" = UserName1;

END
$$
    LANGUAGE 'plpgsql';

-- --------------------------------------------------------

CREATE OR REPLACE FUNCTION newuser(UserName1 varchar, Password1 varchar, TicketCount1 integer)
    RETURNS TABLE
            (
                UserName varchar
            )
AS
$$
BEGIN

RETURN QUERY
    INSERT INTO "User" ("UserName", "Password" ,  "TicketCount")
            VALUES (UserName1,Password1, TicketCount1)
            RETURNING "User"."UserName";

END
$$
LANGUAGE 'plpgsql';




