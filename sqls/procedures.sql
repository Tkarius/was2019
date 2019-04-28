CREATE PROCEDURE was_groupwork.CreateBusinessAnnouncement 

/*--------------------------------------- 

-- Author        :    Mira Pohjola  

-- Date          :    22.04.2019 

-- Description   :    Adds a business announcement. Runs checks for the provided category to prevent from getting an user inserted with invalid announcement. Adds user if not exists. If exists adds the announcement for the existing user 

-----------------------------------------*/ 

    @category varchar(8), 

    @announcement nvarchar(max), 

    @username varchar(60), 

    @useremail varchar(260), 

    @expiration_date datetime 

    AS 

    BEGIN 

    SET NOCOUNT ON 

        DECLARE 

        @UserID int 

            IF @category IN ('buying', 'selling') 

                BEGIN 

IF NOT EXISTS(SELECT AppUser.id FROM was_groupwork.[User] AppUser WHERE AppUser.email = @useremail) 

                    BEGIN 

                        INSERT INTO was_groupwork.[User] (name, email) VALUES (@username, @useremail); 

                        SET @UserID = SCOPE_IDENTITY(); 

                    END 

                ELSE 

SELECT @UserID = AppUser.id FROM was_groupwork.[User] AppUser WHERE AppUser.email = @useremail; 

                    BEGIN 

INSERT INTO was_groupwork.BusinessAnnouncement(announcement, user_id, category, expiration_date) VALUES 

                        (@announcement, @UserID, @category, @expiration_date); 

                    END 

                END 

    END 

---------------------------------------

CREATE PROCEDURE was_groupwork.SelectBusinessAnnouncements 

/*--------------------------------------- 

-- Author        :    Mira Pohjola 

-- Date          :    22.04.2019 

-- Description   :    Selects unexpired BusinessAnnouncements. 

-----------------------------------------*/ 

AS 

BEGIN 

SET NOCOUNT ON 

SELECT u.name, u.email, ba.announcement, ba.category, ba.expiration_date FROM was_groupwork.BusinessAnnouncement ba, was_groupwork.[User] u WHERE expiration_date > GETDATE() AND ba.user_id = u.id; 

 

END 
