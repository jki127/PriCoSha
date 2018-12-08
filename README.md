# PriCoSha
The objective of this project was to design a social-media-like website system that allows users to post and interact with content and other users via a frontend interface while maintaining changes on a backend SQL-based server (mySQL, in this case). 

It was authored by: Madeleine Nicolas (mcpnicolas), Jayson Isaac (jki127), Andrea Vasquez (amvasquez), Anthony Taldone (at3089), and Graeme Ferguson (gqo)
## Getting Started
### Prerequisites
* Golang interpreter and source code

* Go MySQL Driver (Found at: https://github.com/go-sql-driver/mysql)

### Building
The project directory, hereby referred to as `pricosha/`, must be located in the `src/` folder of your GOPATH. Afterwards, navigate to `frontend/` with `pricosha/` and type the following to build to binary:

```bash
make
```
To clean the binary from your directory, type:

```bash
make clean
```

### Running
In the `frontend/` folder, type the following to run the prebuilt binary:

```bash
.frontend/
```

If you want to build and run at the same time, type:

```bash
make run
```

## Base Features
The base features were accomplished according to specifications in the original homework document but are listed here for ease of access:

* View public content

* Login

* View shared content items and info about them

* Manage tags

* Post a content item

* Tag a content item

* Add friend

## Extra Features
All appear of question marks in query sections represent a point at which our prepared statement was passed information from the user.

Database is hereby referred to as DB.

All features require the following files:

pricosha/

    backend/

        backend.go
    
    frontend/

        frontend.go
    
    sql/

        pricosha.sql

All source files descriptions are given within the scope of the `pricosha/` folder.

### Defriend
**Author:** Andrea Vasquez

**Description:**

User navigates to Friend Group page where they see their respective Friend Groups that they own and belong to. If the user owns a group or has respective privileges, next to the friend group information is the add friend and delete friend button. When the user clicks on delete friend, it directs to a new page where the user can type the username of the friend group they want to delete. If the user belongs in the friend group, they get deleted; otherwise an error message pops up that you cannot delete a friend that doesn’t already belong in the friend group.

**Why this feature?**

This is a good feature to add as it allows the owner of the friend group to control who they want in their friend group. It is relevant as the application already contains a feature to add a new member to a friend group, therefore it make logical sense to have a counteractive feature to remove members from the same friend group. With the removal of a friend it is important to tackle what information relevant to the deleted friend gets removes as well (Tags, Comments, & Rates)

**Schema Changes:**

None.

**Queries:**

Deletes row from Belong table

```sql
DELETE FROM Belong
WHERE member_email=?
AND fg_name=?
AND owner_email=?
```

Subquery seen in following queries can be assumed to be within the queries as the following name: `VALID_SUB_QUERY`

Finds all item_ids a user can view

```sql
SELECT item_id
FROM Content_Item
WHERE item_id IN (
    SELECT item_id FROM Share
    WHERE (fg_name, owner_email) IN (
        SELECT fg_name, owner_email
        FROM Belong
        WHERE member_email=?
    )
) OR (is_pub = 1 AND post_time > DATE_SUB(NOW(), INTERVAL 24 HOUR))
OR poster_email=?
```

Clears invalid Tag rows

```sql
DELETE FROM Tag
WHERE item_id NOT IN (
    VALID_SUB_QUERY
)
AND
(tagger_email=?
OR
tagged_email=?)
```

Clears invalid Rate rows

```sql
DELETE FROM Rate
WHERE item_id NOT IN (
    VALID_SUB_QUERY
)
AND email=?
```

Clears invalid Vote rows

```sql
DELETE FROM Vote
WHERE item_id NOT IN (
    VALID_SUB_QUERY
)
AND voter_email=?
```

Clears invalid Comment rows

```sql
DELETE FROM Comment
WHERE item_id NOT IN (
    VALID_SUB_QUERY
)
AND email=?
```

**Source Files**

    backend/

        remove_hanging.go
        
        add_friend_related.go

        friend_group.go

    frontend/

        delete_friend_handler.go

        form_delete_friend_handler.go

        friend_group_handler.go

    web/

        template/

            delete_friend.html

            friend_groups.html


**Images**

Viewing friend groups as GG@nyu.edu

![](https://i.imgur.com/YrJGbdy.png)

Viewing tag of private Content_Item in Friend Group as GG@nyu.edu

![](https://i.imgur.com/iqjd0ul.png)

Deleting GG@nyu.edu from Friend Group

![](https://i.imgur.com/rA82A8o.png)

Viewing lack of friend group as GG@nyu.edu

![](https://i.imgur.com/zazV7Xh.png)

Viewing removed tag of GG@nyu.edu

![](https://i.imgur.com/ZKVPz0Z.png)

### Add Comments
**Author:** Madeleine Nicolas

**Description:**

User navigates to the page for any content item that is visible to that person. On each content item page, there is a button to add a text comment to that content item. After submitting the comment, the page reloads and displays all comments on that item to the user.

**Why this feature?**

This is a good feature to add because it enables users to interact with other users on the PriCoSha platform by commenting on each other’s content items. Unlike rating, comments allow users to give personal, customized feedback on a content item.

**Schema Changes:**

Created a new table Comment with primary keys:

    * Email of Commenter (references Person(email))
    
    * Content_Item ID (references Content_Item(item_id))

    * Timestamp of Comment


**Queries:**

Find comments on Content_Item

```sql
SELECT Comment.email, comment_time, body, f_name, l_name 
FROM Comment JOIN Person ON Comment.email=Person.email 
WHERE item_id=?
ORDER BY comment_time DESC
```

Insert Comment row into DB with primary key and body

```sql
INSERT INTO Comment (email, item_id, comment_time, body)
VALUES (?, ?, ?, ?)
```

**Source Files**

    backend/

        add_comment.go

        content_item.go

    frontend/

        add_comment_handler.go

        content_item_handler.go

        post_item_handler.go

    web/

        template/

            content_item.html

**Images**

Viewing comments

![](https://i.imgur.com/lFG4K1p.png)

Writing a comment

![](https://i.imgur.com/BDoKf2O.png)

Viewing written comments

![](https://i.imgur.com/3vxzrOI.png)

### Location Data
**Author:** Jayson Isaac

**Description:**
Content Items can be filtered by Location

**Why this feature?**
It’s useful to be able to filter files by their location such as images.

**Schema Changes:**

ex. Addition of location attribute to Content_Item table

**Queries:**
Get locations of user’s content items as well as the number of content items that the user has from each location
```sql
SELECT location, count(item_id) FROM Content_Item
WHERE (item_id IN (
	-- All item ids shared in a user's friendgroups
	SELECT item_id FROM Share
	WHERE (fg_name, owner_email) IN (
		-- All friend groups the user belongs to
		SELECT fg_name, owner_email FROM Belong
		WHERE member_email= ?
	)
)  OR (is_pub = 1 AND post_time > DATE_SUB(NOW(), INTERVAL 24 HOUR))
OR poster_email = ?)
AND location IS NOT NULL
GROUP BY location
```

Get all the content items that a user has access to from one location
```sql
SELECT item_id, poster_email, file_path, file_name, post_time
FROM Content_Item
WHERE location = ? AND (item_id IN (
	-- All item ids shared in a user's friendgroups
	SELECT item_id FROM Share
	WHERE (fg_name, owner_email) IN (
		-- All friend groups the user belongs to
		SELECT fg_name, owner_email FROM Belong
		WHERE member_email = ?
	)
)  OR (is_pub = 1 AND post_time > DATE_SUB(NOW(), INTERVAL 24 HOUR))
OR poster_email = ?)
ORDER BY Content_Item.post_time DESC
```

**Source Files**

    backend/
        /content_location.go

    frontend/
        /content_location_handler.go

    web/
        template/
            content_location.html
	    main.html

**Images**

Put images here.

### Profile Page
**Author:** Anthony Taldone 

**Description:**

User navigates to page that displays information relevant to themselves as well as their own personal information which includes: Name, Bio, Friend Group Membership, Friend Group Ownership, and a list of friends. It also allows for management of pending tags, adding bios, viewing content items, commenting and viewing said comments on content items, and viewing ratings on content items. 

**Why this feature?**

This feature was designed to add personalized and social element to the PriCoSha platform with ease of access to major features.

**Schema Changes:**

Addition of bio field to Person

**Queries:**

All appear of question marks represents a point at which our prepared statement was passed information from the user.

Add bio information to database (hereby referred to as DB):

```sql
UPDATE Person
SET bio=? 
WHERE email=?
```

Get list of friends that are members of groups that you own

```sql
SELECT DISTINCT member_email, f_name, l_name
FROM Belong JOIN Person
ON Belong.member_email = Person.email
WHERE Belong.owner_email=?
AND Belong.member_email!=?
```

Get list of friends that are owners of groups of which you are a member

```sql
SELECT DISTINCT owner_email, f_name, l_name
FROM Belong JOIN Person
ON Belong.member_email = Person.email
WHERE member_email=?
AND owner_email!=?
```

Get user's personal information

```sql
SELECT f_name, l_name, bio
FROM Person
WHERE email=?
```

**Source Files**

    backend/

        /profile.go

    frontend/

        /profile_handler.go

        /add_bio_handler.go

    web/

        template/

            profile.html

**Images**

Complete profile page

![](https://i.imgur.com/UCn54ra.png)

Viewing bio

![](https://i.imgur.com/SrCbzfa.png)

Writing bio

![](https://i.imgur.com/uhi4fVK.png)

Seeing pending tag request

![](https://i.imgur.com/yqFK95c.png)

Viewing/adding comments

![](https://i.imgur.com/Z3Iz9qC.png)

Viewing owned friend groups

![](https://i.imgur.com/EsjJiqH.png)

Viewing friend groups which you are a member of

![](https://i.imgur.com/oQwfN2M.png)

Viewing friends list

![](https://i.imgur.com/nabSGY6.png)

### Folders

**Author:** Jayson Isaac

**Description:**

Content Items can be categorized by a common folder name

**Why this feature?**
It’s useful to be able to view common files on one page

**Schema Changes:**

ex. Addition of Folder and Include tables. 

**Queries:**
Get all folders a user has
```sql
SELECT folder_name FROM Folder WHERE email =?
```

Get all content items in a folder
```sql
SELECT item_id, poster_email, file_path, file_name, post_time FROM Include
NATURAL JOIN Content_Item
WHERE folder_name = ? AND email = ?
```

Create Folder
```sql
INSERT INTO Folder (folder_name, email)
VALUES (?, ?)
```

Get Content not in a specified folder (so the user can add those items)
```sql
SELECT item_id, poster_email, file_path, file_name
FROM Content_Item
WHERE (item_id IN (
	-- All item ids shared in a user's friendgroups
	SELECT item_id FROM Share
	WHERE (fg_name, owner_email) IN (
		-- All friend groups the user belongs to
		SELECT fg_name, owner_email FROM Belong
		WHERE member_email=?
	)
)  OR (is_pub = 1 AND post_time > DATE_SUB(NOW(), INTERVAL 24 HOUR))
OR poster_email=?) AND
item_id NOT IN (
	SELECT item_id FROM Include
	NATURAL JOIN Content_Item
	WHERE folder_name = ? AND email = ?
)
ORDER BY Content_Item.post_time DESC
```

Add item to folder
```
INSERT INTO Include (folder_name, email, item_id)
VALUES (?, ?, ?)
```
**Source Files**

    backend/
        /content_folder.go

    frontend/
        /content_folder_handler.go

    web/
        template/
            content_folder.html
	    main.html


### User Privileges
**Author:** Graeme Ferguson

**Description:**

This feature adds three distinct roles for FriendGroup members. These roles are as follows: member which can post content, comment, rate, and tag; mod which can invite new members, ban/kick members, delete posts, and everything a member can do; admin which can delete the group, promote members to mod and vice versa, rename the group, give away admin privilege to another member, and everything a mod can do. Notably, there can only be one group admin. The first group admin is the owner. Ownership of the group transfers with transfer of admin privileges.

**Why this feature?**

Most social media platforms integrate some form of moderation and user privileges into themselves. It is a necessary feature in an online platform where users can behave outside of what owners of groups would prefer.

**Schema Changes:**

Add a role field to Belong 

**Queries:**

Find users at certain role in FriendGroup

```sql
SELECT member_email
FROM Belong
WHERE fg_name=?
AND owner_email=?
AND role=?
```

Find user's role in FriendGroup

```sql
SELECT role
FROM Belong
WHERE fg_name=?
AND owner_email=?
AND member_email=?
```

Update user's role in FriendGroup

```sql
UPDATE Belong
SET role=?
WHERE member_email=?
AND fg_name=?
AND owner_email=?
```

Find Friend Groups user can unshare item from

```sql
SELECT fg_name, owner_email
FROM Share NATURAL JOIN Friend_Group NATURAL JOIN Belong
-- Check if the user is the original poster of the Content_Item
WHERE (member_email IN (
    SELECT poster_email
    FROM Content_Item
    WHERE item_id=?
)
-- Or if the user has mod privileges over the Shared Content_Item
OR role < 2)
AND member_email =?
AND item_id = ?
```

Delete row from share based on primary key

```sql
DELETE FROM Share
WHERE fg_name=?
AND owner_email=?
AND item_id=?
```

Renaming a friend group is the following queries based on creating a new friend group with a new name (and old description), updating all rows in Belong and Share to that new friend group, and deleting the old friend group

```sql
SELECT description
FROM Friend_Group
WHERE fg_name=?
AND owner_email=?
```

```sql
INSERT INTO Friend_Group
(fg_name, owner_email, description)
VALUES
(?, ?, ?)
```

```sql
UPDATE Belong
SET fg_name=?
WHERE fg_name=?
AND owner_email=?
```

```sql
UPDATE Share
SET fg_name=?
WHERE fg_name=?
AND owner_email=?
```

```sql
DELETE FROM Friend_Group
WHERE fg_name=?
AND owner_email=?
```

SwapOwner consists of checking if the new owner exists, checking that that new owner does not already own a group with the same name as the one being swapped to them, creating a new group with that owner (and old description), updating all belong and share rows tied to that group, and deleting the old group. Any repeated quries from Rename have been left out for the sake of brevity.

```sql
SELECT email
FROM Person
WHERE email=?
```

```sql
SELECT fg_name
FROM Friend_Group
Where fg_name=?
AND owner_email=?
```

```sql
UPDATE Belong
SET owner_email=?
WHERE fg_name=?
AND owner_email=?
```

```sql
UPDATE Share
SET owner_email=?
WHERE fg_name=?
AND owner_email=?
```

Delete Friend Group specified by primary key

```sql
DELETE FROM Friend_Group
WHERE fg_name=?
AND owner_email=?
```

**Source Files**

    backend/

        friend_group.go

        manage_privileges.go

    frontend/

        add_friend_handler.go

        change_owner_handler.go

        change_privilege_handler.go

        delete_friend_handler.go

        delete_group_handler.go

        form_add_friend_handler.go

        form_delete_friend_handler.go

        friend_group_handler.go

        manage_privileges_handler.go

        profile_handler.go

        rename_group_handler.go

        unshare_handler.go

    web/

        template/

            content_item.html

            friend_groups.html

            manage_privileges.html

            profile.html

**Images**

BB@nyu.edu is an admin of his family group

![](https://i.imgur.com/TvGNJVo.png)

BB@nyu.edu is a mod of AA@nyu.edu's family group

![](https://i.imgur.com/GUpl52T.png)

BB can manage privileges, add, and delete friends of both groups

![](https://i.imgur.com/PqgLu8v.png)

BB can unshare posts from both groups

![](https://i.imgur.com/xovqfZe.png)

BB can promote/demote members in his group

![](https://i.imgur.com/ieRAhOP.png)

BB can rename his group

![](https://i.imgur.com/CyE0LnO.png)

![](https://i.imgur.com/n0lJWHp.png)

BB can delete his group

![](https://i.imgur.com/ZQezisR.png)

BB can swap ownership of the group to FF@nyu.edu

![](https://i.imgur.com/ZwMN8JW.png)

### Content Item Deletion
**Author:** Madeleine Nicolas

**Description:**

This feature adds the ability for user's to delete Content_Items they have posted.

**Why this feature?**

Removing posts is a standard feature of any social media platform. If you can post content, you should be able to remove it.

**Schema Changes:**

None.

**Queries:**

Verify that user is Content_Item poster by selecting that row and checking if it exists

```sql
SELECT poster_email 
FROM Content_Item 
WHERE item_id = ?
```

Delete Content_Item from DB

```sql
DELETE FROM Content_Item
WHERE item_id=?
```

**Source Files**

    backend/

        delete_content.go

    frontend/

        content_item_handler.go

        delete_item_handler.go

    web/

        template/

            content_item.html

**Images**

Selecting delete

![](https://i.imgur.com/BQoeFVf.png)

Being alerted that delete will occur

![](https://i.imgur.com/UCjdMEI.png)

### Poll Content Item
**Author:** Graeme Ferguson 

**Description:**

This feature adds a new type of Content_Item called Poll that allows users to vote on options (one vote per user per poll) which is then displayed on the frontend as Content_Item information with interactive buttons to vote differently or for the first time.

**Why this feature?**

Many social media platforms feature polls (such as Messenger). They are a useful way for users to interact and choose various things or simply list and aggregate opinions.

**Schema Changes:**

* Add format field to Content_Item

* Created a new table Vote with primary keys:

    * Email of Voter (references Person(email))
    
    * Content_Item ID (references Content_Item(item_id))

    * Choice

**Queries:**

Find if item is poll by selecting it only if it is a poll


```sql
SELECT item_id
FROM Content_Item
WHERE format=1
AND item_id=?
```

Find all choices and the votes for said choices for certain poll

```sql
SELECT choice, COUNT(*) as vote_count
FROM Vote
WHERE item_id=?
GROUP BY choice
ORDER BY vote_count DESC
```

Find if a vote has been cast by selecting the vote if it exists

```sql
SELECT item_id
FROM Vote
WHERE voter_email=?
AND item_id=?
```

Update voter's choice if new choice is chosen

```sql
UPDATE Vote
SET choice=?
WHERE voter_email=?
AND item_id=?
```

If voter has not voted on poll before, insert Vote row with primary key

```sql
INSERT INTO Vote
(voter_email, item_id, choice)
VALUES
(?, ?, ?)
```

**Source Files**

    backend/

        content_item.go

        polls.go

        remove_hanging.go

    frontend/

        add_vote_handler.go

    web/

        template/

            content_item.html

            main.html

**Images**

View identifying poll info on main page

![](https://i.imgur.com/8tZm6Ms.png)

View detailed poll info on content item page

![](https://i.imgur.com/3f1lvZ8.png)

Click option button to vote on pre-existing option

![](https://i.imgur.com/8MGTFRd.png)

Write in option to vote differently

![](https://i.imgur.com/3Ton1Ly.png)

![](https://i.imgur.com/jYs80Wd.png)

## Base Feature Contribution
View Shared Content - Jayson
Manage Tags - Anthony
Post a Content Item - Maddie
Tag a Content Item - Graeme
Add Friend - Andrea
