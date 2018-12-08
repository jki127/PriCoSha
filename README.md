# PriCoSha
The objective of this project was to design a social-media-like website system that allows users to post and interact with content and other users via a frontend interface while maintaining changes on a backend SQL-based server (mySQL, in this case). 

It was authored by: Madeleine Nicolas (mcpnicolas), Jayson Isaac (jki127), Andrea Vasquez (amvasquez), Anthony Taldone (at3089), and Graeme Ferguson (gqo)
## Getting Started
### Prerequisites
* Golang interpreter and source code

* Go MySQL Driver (Found at: https://github.com/go-sql-driver/mysql)

### Building
The project directory, hereby referred to as /pricosha, must be located in the /src folder of your GOPATH. Afterwards, navigate to /frontend with /pricosha and type the following to build to binary:

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
### Defriend
| Author | Description | Schema Changes | Queries | Source Code Location | Images |
| ------ | ----------- | -------------- | ------- | -------------------- | ------ |
| Anthony Taldone | User navigates to page that displays information relevant to themselves as well as their own personal information which includes: Name, Bio, Friend Group Membership, Friend Group Ownership, and a list of friends. It also allows for management of pending tags, adding bios, viewing content items, commenting and viewing said comments on content items, and viewing ratings on content items. This feature was designed to add personalized and social element to the PriCoSha platform with ease of access to major features. | Addition of Bio 