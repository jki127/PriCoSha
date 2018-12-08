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

Put images here.

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

ex. Add bio information to DB

```sql
UPDATE Person SET bio=? WHERE email=?
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

Put images here.

### Location Data
**Author:** Anthony Taldone 

**Description:**

Description text

**Why this feature?**

Selling text

**Schema Changes:**

ex. Addition of bio field to Person

**Queries:**

ex. Add bio information to DB

```sql
UPDATE Person SET bio=? WHERE email=?
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

Put images here.

### Folders
**Author:** Anthony Taldone 

**Description:**

Description text

**Why this feature?**

text

**Schema Changes:**

ex. Addition of bio field to Person

**Queries:**

ex. Add bio information to DB

```sql
UPDATE Person SET bio=? WHERE email=?
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

Put images here.

### User Privileges
**Author:** Anthony Taldone 

**Description:**

Description text

**Why this feature?**

text

**Schema Changes:**

ex. Addition of bio field to Person

**Queries:**

ex. Add bio information to DB

```sql
UPDATE Person SET bio=? WHERE email=?
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

Put images here.

### Content Item Deletion
**Author:** Anthony Taldone 

**Description:**

Description text

**Why this feature?**

text

**Schema Changes:**

ex. Addition of bio field to Person

**Queries:**

ex. Add bio information to DB

```sql
UPDATE Person SET bio=? WHERE email=?
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

Put images here.

### Poll Content Item
**Author:** Anthony Taldone 

**Description:**

Description text

**Why this feature?**

text

**Schema Changes:**

ex. Addition of bio field to Person

**Queries:**

ex. Add bio information to DB

```sql
UPDATE Person SET bio=? WHERE email=?
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

Put images here.