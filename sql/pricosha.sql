-- Drop old pricosha database and create new one for use
-- IMPORTANT: save old data from pricosha database if necessary
drop database pricosha;
create database pricosha;
use pricosha;

SELECT "------------------------- Adding Tables -------------------------" as "";
-- Populate database with tables.
SELECT "Adding Person Table" as "";
create table Person
    (email  varchar(64),
    password char(64) not null,
    f_name varchar(32),
    l_name varchar(32),
    primary key (email)
    );

SELECT "Adding Content_Item Table" as "";
/* With the addition of format (0 = regular item, 1 = poll):
    If format = 1:
        file_name = question asked in poll
*/
create table Content_Item
    (item_id  int not null AUTO_INCREMENT,
    poster_email varchar(64),
    file_path  varchar(64),
    file_name  varchar(64),
    post_time timestamp,
    is_pub boolean,
    format int DEFAULT 0,
    location varchar(128),
    primary key (item_id),
    foreign key (poster_email) references Person(email)
        on delete set null
    );

SELECT "Adding Friend_Group Table" as "";
create table Friend_Group
    (fg_name varchar(32),
    owner_email varchar(64),
    description varchar(64),
    primary key (fg_name, owner_email),
    foreign key (owner_email) references Person(email)
        on delete cascade
    );

SELECT "Adding Tag Table" as "";
create table Tag
    (tagger_email varchar(64),
    tagged_email varchar(64),
    item_id int,
    status boolean,
    tag_time timestamp,
    primary key (tagger_email, tagged_email, item_id),
    foreign key (tagger_email) references Person(email)
        on delete cascade,
    foreign key (tagged_email) references Person(email)
        on delete cascade,
    foreign key (item_id) references Content_Item(item_id)
        on delete cascade
    );

SELECT "Adding Comment Table" as "";
create table Comment
    (email varchar(64),
    item_id int,
    comment_time timestamp,
    body varchar(64),
    primary key (email, item_id, comment_time),
    foreign key (email) references Person(email)
        on delete cascade,
    foreign key (item_id) references Content_Item(item_id)
        on delete cascade
    );

SELECT "Adding Rate Table" as "";
create table Rate
    (email varchar(64),
    item_id int,
    rate_time timestamp,
    emoji varchar(20) CHARACTER SET utf8mb4,
    primary key (email, item_id),
    foreign key (email) references Person(email)
        on delete cascade,
    foreign key (item_id) references Content_Item(item_id)
        on delete cascade
    );

SELECT "Adding Share Table" as "";
create table Share
    (fg_name varchar(32),
    owner_email varchar(64),
    item_id int,
    primary key (fg_name, owner_email, item_id),
    foreign key (fg_name, owner_email) references Friend_Group(fg_name, owner_email)
        on delete cascade,
    foreign key (item_id) references Content_Item(item_id)
        on delete cascade
    );

SELECT "Adding Belong Table" as "";
create table Belong
    (member_email varchar(64),
    fg_name varchar(32),
    owner_email varchar(64),
    role int DEFAULT 2,
    primary key (member_email, fg_name, owner_email),
    foreign key (member_email) references Person(email)
        on delete cascade,
    foreign key (fg_name, owner_email) references Friend_Group(fg_name, owner_email)
        on delete cascade
    );

SELECT "Adding Vote Table" as "";
create table Vote
    (voter_email varchar(64),
    item_id int,
    choice varchar(64),
    primary key (voter_email, item_id),
    foreign key (voter_email) references Person(email)
        on delete cascade,
    foreign key (item_id) references Content_Item(item_id)
        on delete cascade
    );

SELECT "------------------------- Adding Data -------------------------" as "";
-- Insert dummy data
-- Adds Persons
SELECT "Adding Persons" as "";
INSERT INTO Person
    (email, password, f_name, l_name)
VALUES
    ("AA@nyu.edu", SHA2("AA",256), "Ann", "Anderson"),
    ("BB@nyu.edu", SHA2("BB",256),"Bob", "Baker"),
    ("CC@nyu.edu", SHA2("CC",256), "Cathy", "Chang"),
    ("DD@nyu.edu", SHA2("DD",256), "David", "Davidson"),
    ("EE@nyu.edu", SHA2("EE",256), "Ellen", "Ellenberg"),
    ("FF@nyu.edu", SHA2("FF",256), "Fred", "Fox"),
    ("GG@nyu.edu", SHA2("GG",256), "Gina", "Gupta"),
    ("HH@nyu.edu", SHA2("HH",256), "Helen", "Harper"),
    ("HMH@nyu.edu", SHA2("HH",256), "Helen", "Harper");
    -- ("NotBB@nyu.edu", SHA2("BB",256), "Bob", "Baker");

-- Adds Friend_Groups
SELECT "Adding Friend_Groups" as "";
INSERT INTO Friend_Group
    (fg_name, owner_email, description)
VALUES
    ("family", "AA@nyu.edu", "Ann's Family"),
    ("family", "BB@nyu.edu", "Bob's Family"),
    ("roommates", "AA@nyu.edu", "Ann's Roommates"),
    ("nomembersnomods", "AA@nyu.edu", "Test case!");

-- Adds Owners to Friend Groups in Belong
SELECT "Adding Owners to Belong" as "";
INSERT INTO Belong
    (member_email, fg_name, owner_email, role)
VALUES
    ("AA@nyu.edu", "family", "AA@nyu.edu", 0),
    ("BB@nyu.edu", "family", "BB@nyu.edu", 0),
    ("AA@nyu.edu", "roommates", "AA@nyu.edu", 0),
    ("AA@nyu.edu", "nomembersnomods", "AA@nyu.edu", 0);

-- Adds Members to Friend Groups in Belong
SELECT "Adding Members to Belong" as "";
INSERT INTO Belong
    (member_email, fg_name, owner_email)
VALUES
    ("CC@nyu.edu", "family", "AA@nyu.edu"),
    ("DD@nyu.edu", "family", "AA@nyu.edu"),
    ("EE@nyu.edu", "family", "AA@nyu.edu"),
    ("FF@nyu.edu", "family", "BB@nyu.edu"),
    ("EE@nyu.edu", "family", "BB@nyu.edu"),
    ("GG@nyu.edu", "roommates", "AA@nyu.edu"),
    ("HH@nyu.edu", "roommates", "AA@nyu.edu");

-- Adds Mods to Friend Groups in Belong
SELECT "Adding Mods to Belong" as "";
INSERT INTO Belong
    (member_email, fg_name, owner_email, role)
VALUES
    ("BB@nyu.edu", "family", "AA@nyu.edu", 1);

-- Content_Items should be added or shared ONLY in this section

-- Adds Content_Items (1, 2, 3, 4, 5)
SELECT "Adding Content_Items" as "";
INSERT INTO Content_Item
    (poster_email, file_path, file_name, post_time, is_pub, location)
VALUES
    ("AA@nyu.edu", "/Photos/Animals", "Whiskers", "2010-12-01 03:39:01", TRUE, "Astoria"),
    ("AA@nyu.edu", "/Photos/Room21", "leftovers in fridge", "2014-06-10 04:00:30", FALSE, "Astoria"),
    ("BB@nyu.edu", "/Photos/Pets", "Rover", "2017-04-02 07:17:02", FALSE, "East Flatbush"),
    ("CC@nyu.edu", "/Taxes/2009/EpsteinMemes","OPM_Epstein", "2018-12-02 03:12:10", TRUE, "East Flatbush"),
    ("EE@nyu.edu", "no", "no", "2018-12-02 03:12:11", TRUE, "DUMBO");

-- Shares Content_Items
SELECT "Adding Shares of Content_Items" as "";
INSERT INTO Share
    (fg_name, owner_email, item_id)
VALUES
    ("family", "AA@nyu.edu", 1),
    ("roommates", "AA@nyu.edu", 2),
    ("family", "BB@nyu.edu", 3),
    ("family", "AA@nyu.edu", 4),
    ("family", "BB@nyu.edu", 5);

-- Adds Content_Items in Past 24 Hours (6, 7, 8)
SELECT "Adding Content_Items in Past 24 Hours" as "";
INSERT INTO Content_Item
    (poster_email, file_path, file_name, post_time, is_pub)
VALUES
    ("HH@nyu.edu", "/home/data/pie.jpg", "Pie", NOW(), 1),
    ("HH@nyu.edu", "/home/data/turkey.jpg", "Turkey", NOW(), 1),
    ("HH@nyu.edu", "/home/data/mashed.jpg", "Mashed Potatoes", NOW(), 1);

-- Add Poll Content_Items in Past 24 Hours (9, 10, 11)
SELECT "Adding Poll Content_Items in Past 24 Hours" as "";
INSERT INTO Content_Item
    (poster_email, file_path, file_name, post_time, is_pub, format)
VALUES
    ("BB@nyu.edu", "nothing", "Do you like apples?", NOW(), 1, 1),
    ("CC@nyu.edu", "nothing", "Best yoghurt?", NOW(), 0, 1),
    ("CC@nyu.edu", "nothing", "What should I do with my Friday night?", NOW(), 0, 1);

-- Shares Polls
SELECT "Adding Shares of Polls" as "";
INSERT INTO Share
    (fg_name, owner_email, item_id)
VALUES
    ("family", "AA@nyu.edu", 9),
    ("family", "AA@nyu.edu", 10),
    ("family", "AA@nyu.edu", 11);

-- END OF CONTENT_ITEM SHARING AND ADDING SECTION

-- Add Votes
SELECT "Adding Votes to Polls" as "";
INSERT INTO Vote
    (voter_email, item_id, choice)
VALUES
    ("EE@nyu.edu", 10, "TEST CASE FOR CLEANUP"),
    -- Above insert is for testing the CleanUp() function
    ("BB@nyu.edu", 9, "I LOVE APPLES"),
    ("AA@nyu.edu", 9, "Yes"),
    ("CC@nyu.edu", 9, "YES"),
    ("CC@nyu.edu", 11, "Party!"),
    ("BB@nyu.edu", 11, "Sleep!"),
    ("AA@nyu.edu", 11, "Work on databases :(");

-- Adds Tags
SELECT "Adding Tags" as "";
INSERT INTO Tag
    (tagger_email, tagged_email, item_id, status, tag_time)
VALUES
    ("EE@nyu.edu", "BB@nyu.edu", 10, TRUE, NOW()),
    ("BB@nyu.edu", "EE@nyu.edu", 10, TRUE, NOW()),
    -- Above 2 inserts are for testing the CleanUp() function
    ("AA@nyu.edu", "GG@nyu.edu", 2, TRUE, "2018-11-21 05:10:30"),
    ("DD@nyu.edu", "CC@nyu.edu", 4, FALSE, "2018-09-18 03:12:30"),
    ("BB@nyu.edu", "FF@nyu.edu", 3, TRUE, "2018-10-27 09:22:30");

-- Adds Comments
SELECT "Adding Comments" as "";
INSERT INTO Comment
    (email, item_id, comment_time, body)
VALUES
    ("EE@nyu.edu", 10, NOW(), "TEST CASE FOR CLEANUP"),
    -- Above insert is for testing the CleanUp() function
    ("EE@nyu.edu", 1, "2018-11-27 09:22:30", "this is a comment"),
    ("CC@nyu.edu", 4, "2018-03-23 12:22:30", "loveitttt"),
    ("HH@nyu.edu", 2, "2018-07-17 04:22:30", "yummy");

-- Adds Ratings
SELECT "Adding Rates" as "";
INSERT INTO Rate
    (email, item_id, rate_time, emoji)
VALUES
    ("EE@nyu.edu", 10, NOW(), "👍"),
    -- Above insert is for testing the CleanUp() function
    ("EE@nyu.edu", 1, "2018-11-27 09:22:30", "👍"),
    ("CC@nyu.edu", 4, "2018-03-23 12:22:30", "👍"),
    ("HH@nyu.edu", 2, "2018-07-17 04:22:30", "👍");
