# Database statements


# -- Crete User model in database 
CREATE TABLE users(
   ID   INT              NOT NULL,
   name VARCHAR (100)     NOT NULL,
   email  VARCHAR (50) ,
   status    BOOLEAN,       
   PRIMARY KEY (ID)
);