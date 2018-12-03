-- Drop old pricosha database and create new one for use
-- IMPORTANT: save old data from pricosha database if necessary
drop database pricosha;
create database pricosha;
use pricosha;

-- Populate database with tables.
create table Person
    (email  varchar(64),
    password char(64) not null,
    f_name varchar(32),
    l_name varchar(32),
    primary key (email)
    );

create table Content_Item
    (item_id  int not null AUTO_INCREMENT,
    poster_email varchar(64),
    file_path  varchar(64),
    file_name  varchar(64),
    post_time timestamp,
    is_pub boolean,
    primary key (item_id),
    foreign key (poster_email) references Person(email)
        on delete set null
    );


create table Friend_Group
    (fg_name varchar(32),
    owner_email varchar(64),
    description varchar(64),
    primary key (fg_name, owner_email),
    foreign key (owner_email) references Person(email)
        on delete cascade
    );


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


create table Belong
    (member_email varchar(64),
    fg_name varchar(32),
    owner_email varchar(64),
    primary key (member_email, fg_name, owner_email),
    foreign key (member_email) references Person(email)
        on delete cascade,
    foreign key (fg_name, owner_email) references Friend_Group(fg_name, owner_email)
        on delete cascade
    );

-- Insert dummy data
INSERT INTO Person VALUES ("AA@nyu.edu", SHA2("AA",256), "Ann", "Anderson");
INSERT INTO Person VALUES ("BB@nyu.edu", SHA2("BB",256),"Bob", "Baker");
INSERT INTO Person VALUES ("CC@nyu.edu", SHA2("CC",256), "Cathy", "Chang");
INSERT INTO Person VALUES ("DD@nyu.edu", SHA2("DD",256), "David", "Davidson");
INSERT INTO Person VALUES ("EE@nyu.edu", SHA2("EE",256), "Ellen", "Ellenberg");
INSERT INTO Person VALUES ("FF@nyu.edu", SHA2("FF",256), "Fred", "Fox");
INSERT INTO Person VALUES ("GG@nyu.edu", SHA2("GG",256), "Gina", "Gupta");
INSERT INTO Person VALUES ("HH@nyu.edu", SHA2("HH",256), "Helen", "Harper");
INSERT INTO Friend_Group VALUES ("family", "AA@nyu.edu", "Ann's Family");
INSERT INTO Belong VALUES ("AA@nyu.edu", "family", "AA@nyu.edu");
INSERT INTO Belong VALUES ("CC@nyu.edu", "family", "AA@nyu.edu");
INSERT INTO Belong VALUES ("DD@nyu.edu", "family", "AA@nyu.edu");
INSERT INTO Belong VALUES ("EE@nyu.edu", "family", "AA@nyu.edu");
INSERT INTO Friend_Group VALUES ("family", "BB@nyu.edu", "Bob's Family");
INSERT INTO Belong VALUES ("BB@nyu.edu", "family", "BB@nyu.edu");
INSERT INTO Belong VALUES ("FF@nyu.edu", "family", "BB@nyu.edu");
INSERT INTO Belong VALUES ("EE@nyu.edu", "family", "BB@nyu.edu");
INSERT INTO Friend_Group VALUES ("roommates", "AA@nyu.edu", "Ann's Roommates");
INSERT INTO Belong VALUES ("AA@nyu.edu", "roommates", "AA@nyu.edu");
INSERT INTO Belong VALUES ("GG@nyu.edu", "roommates", "AA@nyu.edu");
INSERT INTO Belong VALUES ("HH@nyu.edu", "roommates", "AA@nyu.edu");
INSERT INTO Content_Item VALUES (1, "AA@nyu.edu", "/Photos/Animals", "Whiskers", "2010-12-01 03:39:01", TRUE);
INSERT INTO Share VALUES ("family", "AA@nyu.edu", 1);
INSERT INTO Content_Item VALUES (2, "AA@nyu.edu", "/Photos/Room21", "leftovers in fridge", "2014-06-10 04:00:30", FALSE);
INSERT INTO Share VALUES ("roommates", "AA@nyu.edu", 2);
INSERT INTO Content_Item VALUES (3, "BB@nyu.edu", "/Photos/Pets", "Rover", "2017-04-02 07:17:02", FALSE);
INSERT INTO Share VALUES ("family", "BB@nyu.edu", 3);
INSERT INTO Content_Item VALUES (4, "CC@nyu.edu", "/Taxes/2009/EpsteinMemes","OPM_Epstein", "2018-12-02 03:12:10", TRUE);
INSERT INTO Share VALUES ("family", "AA@nyu.edu", 4);
INSERT INTO Tag VALUES ("AA@nyu.edu", "GG@nyu.edu", 2, TRUE, "2018-11-21 05:10:30");
INSERT INTO Tag VALUES ("DD@nyu.edu", "CC@nyu.edu", 4, FALSE, "2018-09-18 03:12:30");
INSERT INTO Tag VALUES ("BB@nyu.edu", "FF@nyu.edu", 3, TRUE, "2018-10-27 09:22:30");
INSERT INTO Rate VALUES ("EE@nyu.edu", 1, "2018-11-27 09:22:30", "0x1f61a");
INSERT INTO Rate VALUES ("HH@nyu.edu", 2, "2018-07-17 04:22:30", "0x1f61a");
INSERT INTO Rate VALUES ("CC@nyu.edu", 4, "2018-03-23 12:22:30", "0x1f61a");
INSERT INTO content_item
  (poster_email, file_path, file_name, post_time, is_pub)
VALUES
  ("HH@nyu.edu", "/home/data/pie.jpg", "Pie", NOW(), 1),
  ("HH@nyu.edu", "/home/data/turkey.jpg", "Turkey", NOW(), 1),
  ("HH@nyu.edu", "/home/data/mashed.jpg", "Mashed Potatoes", NOW(), 1)
