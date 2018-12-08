# PriCoSha
The objective of this project was to design a social-media-like website system that allows users to post and interact with content and other users via a frontend interface while maintaining changes on a backend SQL-based server (mySQL, in this case). 

It was authored by: Madeleine Nicolas (mcpnicolas), Jayson Isaac (jki127), Andrea Vasquez (amvasquez), Anthony Taldone (at3089), and Graeme Ferguson (gqo)
## Getting Started
### Prerequisites
* Golang interpreter and source code

* Go MySQL Driver (Found at: https://github.com/go-sql-driver/mysql)

### Building
The project directory, hereby referred to as `/pricosha`, must be located in the `/src` folder of your GOPATH. Afterwards, navigate to /frontend with `/pricosha` and type the following to build to binary:

```bash
make
```
To clean the binary from your directory, type:

```bash
make clean
```

### Running
In the /frontend folder, type the following to run the prebuilt binary:

```bash
./frontend
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

### Defriend
**Author:** Andrea Vasquez

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
Assuming you are within `/pricosha`:

    /backend

        /profile.go

    /frontend

        /profile_handler.go

        /add_bio_handler.go

    /web

        /template

            profile.html

**Images**

Put images here.

### Add Comments
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
Assuming you are within `/pricosha`:

    /backend

        /profile.go

    /frontend

        /profile_handler.go

        /add_bio_handler.go

    /web

        /template

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
Assuming you are within `/pricosha`:

    /backend

        /profile.go

    /frontend

        /profile_handler.go

        /add_bio_handler.go

    /web

        /template

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
UPDATE Person SET bio=? WHERE email=?
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
Assuming you are within `/pricosha`:

    /backend

        /profile.go

    /frontend

        /profile_handler.go

        /add_bio_handler.go

    /web

        /template

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
Assuming you are within `/pricosha`:

    /backend

        /profile.go

    /frontend

        /profile_handler.go

        /add_bio_handler.go

    /web

        /template

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
Assuming you are within `/pricosha`:

    /backend

        /profile.go

    /frontend

        /profile_handler.go

        /add_bio_handler.go

    /web

        /template

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
Assuming you are within `/pricosha`:

    /backend

        /profile.go

    /frontend

        /profile_handler.go

        /add_bio_handler.go

    /web

        /template

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
Assuming you are within `/pricosha`:

    /backend

        /profile.go

    /frontend

        /profile_handler.go

        /add_bio_handler.go

    /web

        /template

            profile.html

**Images**

Put images here.