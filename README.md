## Learning Objectives

- Build a CRUD API app using Go
- Gain some familiarity with AWS Cloud Platform

## Overview
### API Endpoints

HTTP verbs, or methods, should be used in compliance with their definitions under the [HTTP/1.1](http://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html) standard.
The action taken on the representation will be contextual to the media type being worked on and its current state. Here's an example of how HTTP verbs map to create, read, update, delete operations in a particular context:

| HTTP Endpoints                        | ResponseCode            | DESC            |
| -----------                           | --------------- | --------------- |
| [GET: /user/[id]](#get-userid)        | 200(Content Found) / 204(No Content Found) | Find user by id |
| [POST: /user](#post-user)             | 201 (Successfully Created)    | Create user     |

### GET: /user/[id]

```bash
curl --location --request GET 'http://localhost:8080/user/123'
```

Response body:

    {
        "id":"124",
        "name":"MyName",
        "email":"myname@email.com"
    }

### POST: /user

```bash
curl --location --request POST 'http://goDynamoLB-459636926.us-west-1.elb.amazonaws.com/user' \
--header 'Content-Type: application/json' \
--data-raw '{   
    "id": "124",
    "name": "as",
    "email": "as"
}'
```

Request body:

    {   
        "id": "124",
        "name": "MyName",
        "email": "myname@email.com"
    }


## Architecture Overview

## Deployment Strategies