create table person
    (email  varchar(20),
    password char(64),
    f_name varchar(30),
    l_name varchar(30),
    primary key (email)
    );

create table content_item
    (item_id  int not null AUTO_INCREMENT,
    poster_email varchar(20),
    file_path  varchar(30),
    file_name  varchar(30),
    post_time timestamp,
    is_pub boolean,
    primary key (item_id),
    foreign key (poster_email) references person(email)
    );


create table friend_group
    (fg_name varchar(20),
    owner_email varchar(20),
    description varchar(50),
    primary key (fg_name, owner_email),
    foreign key (owner_email) references person(email) on delete cascade
    );


create table tags
    (tag_email varchar(20),
    tagged_email varchar(20),
    item_id int,
    status varchar(10),
    tag_time timestamp,
    primary key (tag_email, tagged_email, item_id),
    foreign key (tag_email) references person(email) on delete cascade,
    foreign key (tagged_email) references person(email) on delete cascade,
    foreign key (item_id) references content_item(item_id) on delete cascade
    );

create table rates
    (email varchar(20),
    item_id int,
    rate_time timestamp,
    emojii varchar(20) CHARACTER SET utf8mb4,
    primary key (email, item_id),
    foreign key (email) references person(email) on delete cascade,
    foreign key (item_id) references content_item(item_id) on delete cascade
    );


create table share
    (fg_name varchar(20),
    owner_email varchar(20),
    item_id int,
    primary key (fg_name, owner_email, item_id),
    foreign key (fg_name, owner_email) references friend_group(fg_name, owner_email) on delete cascade,
    foreign key (item_id) references content_item(item_id) on delete cascade
    );


create table belong
    (email varchar(20),
    fg_name varchar(20),
    owner_email varchar(20),
    primary key (email, fg_name, owner_email),
    foreign key (email) references person(email) on delete cascade,
    foreign key (fg_name, owner_email) references friend_group(fg_name, owner_email) on delete cascade
    );

-- Part B: INSERTS

INSERT INTO person VALUES ("AA@nyu.edu", SHA2("AA",256), "Ann", "Anderson");
INSERT INTO person VALUES ("BB@nyu.edu", SHA2("BB",256),"Bob", "Baker");
INSERT INTO person VALUES ("CC@nyu.edu", SHA2("CC",256), "Cathy", "Chang");
INSERT INTO person VALUES ("DD@nyu.edu", SHA2("DD",256), "David", "Davidson");
INSERT INTO person VALUES ("EE@nyu.edu", SHA2("EE",256), "Ellen", "Ellenberg");
INSERT INTO person VALUES ("FF@nyu.edu", SHA2("FF",256), "Fred", "Fox");
INSERT INTO person VALUES ("GG@nyu.edu", SHA2("GG",256), "Gina", "Gupta");
INSERT INTO person VALUES ("HH@nyu.edu", SHA2("HH",256), "Helen", "Harper");
INSERT INTO friend_group VALUES ("family", "AA@nyu.edu", "Ann's Family");
INSERT INTO belong VALUES ("AA@nyu.edu", "family", "AA@nyu.edu");
INSERT INTO belong VALUES ("CC@nyu.edu", "family", "AA@nyu.edu");
INSERT INTO belong VALUES ("DD@nyu.edu", "family", "AA@nyu.edu");
INSERT INTO belong VALUES ("EE@nyu.edu", "family", "AA@nyu.edu");
INSERT INTO friend_group VALUES ("family", "BB@nyu.edu", "Bob's Family");
INSERT INTO belong VALUES ("BB@nyu.edu", "family", "BB@nyu.edu");
INSERT INTO belong VALUES ("FF@nyu.edu", "family", "BB@nyu.edu");
INSERT INTO belong VALUES ("EE@nyu.edu", "family", "BB@nyu.edu");
INSERT INTO friend_group VALUES ("roommates", "AA@nyu.edu", "Ann's Roommates");
INSERT INTO belong VALUES ("AA@nyu.edu", "roommates", "AA@nyu.edu");
INSERT INTO belong VALUES ("GG@nyu.edu", "roommates", "AA@nyu.edu");
INSERT INTO belong VALUES ("HH@nyu.edu", "roommates", "AA@nyu.edu");
INSERT INTO content_item(poster_email, file_name, is_pub) values ("AA@nyu.edu", "Whiskers", FALSE);
INSERT INTO share VALUES ("family", "AA@nyu.edu", 1);
INSERT INTO content_item(poster_email, file_name, is_pub) values ("AA@nyu.edu", "leftovers in fridge", FALSE);
INSERT INTO share VALUES ("roommates", "AA@nyu.edu", 2);
INSERT INTO content_item(poster_email, file_name, is_pub) values ("BB@nyu.edu", "Rover", FALSE);
INSERT INTO share VALUES ("family", "BB@nyu.edu", 3);
INSERT INTO content_item
  (poster_email, file_path, file_name, post_time, is_pub)
VALUES
  ("HH@nyu.edu", "/home/data/pie.jpg", "Pie", NOW(), 1),
  ("HH@nyu.edu", "/home/data/turkey.jpg", "Turkey", NOW(), 1),
  ("HH@nyu.edu", "/home/data/mashed.jpg", "Mashed Potatoes", NOW(), 1)
