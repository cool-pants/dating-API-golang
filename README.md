# Stumble RESTful API

## Description

A RESTful API has been made using Golang and the Gin framework as the base.
The functionalities of the API backend are:

- Can convert JSON data into SQL data then stored into an RDBMS(Postgres, in this case)
- Allows operators to find all matches (A likes B and B likes A) within all registered users
- Allows operators to find matched users for a particular user
- Allows operators to find All users within a distance _'k'_ from an User _'X'_
- Allows operators to query users by their name or part of their name

## API Documentation

[Postman Documentation](https://documenter.getpostman.com/view/11790517/UyxnCjNZ)
