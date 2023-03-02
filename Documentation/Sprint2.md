# Sprint 2 Backend

## Primary Updates
We decided to implement a more "morally" constructive initiative within our program. The platform will still be used as a means of "gambling", of sorts, but it will have the added component of stock market connectivity. This way, users can develop skill sets while competing against each other, and also have the opportunity to gain experience in trading. This, we felt, allowed for an inherently more constructive and educative experience.

### Stock Market Connectivity
To facilitate trades from the eColosseum platform, we made use of Alpaca's Trade API. For the sake of testing, we assumed that the user already has an API key with Alpaca, but in the future, we will implement a way for users to create their own API keys. The user will then be able to connect their Alpaca account to the eColosseum platform, and then be able to trade stocks from the platform.
Implementation of the GUI for trading in eColosseum is simple: we will include input boxes to specify the stock, number of shares, price, and type of order (market or limit).
There should ideally be a way to view the user's current portfolio, but this is not a priority for the time being.
There should also be error handling for invalid stock ticker inputs, but for now we will assume the input ticker is the intended one. 

### Database API
We will be using a SQLite database to store player information. We need persistence for scores, "cash balance", and basic user information. This database will be seperate from that of the user's stock portfolio. 

Documentation for the E-Colosseum can be found at
    * [openai.yaml](/backend/openai.yaml)
    * [Swagger](https://app.swaggerhub.com/apis/b-cheek/E-Colosseum-API/1.0.0)
    * Right here!

**User API**
========

REST API for managing user data

Version: 1.0.0

All rights reserved

http://apache.org/licenses/LICENSE-2.0.html

Methods
-------

\[ Jump to [Models](#__Models) \]

# Default

<details><summary><code>get /users</code></summary>

[Up](#__Methods)

    get /users

Get all users (usersGet)

### Return type

array\[[UserResponse](#UserResponse)\]

### Example data

Content-Type: application/json

    [ {
      "birthday" : "birthday",
      "password" : "password",
      "id" : 0,
      "username" : "username"
    }, {
      "birthday" : "birthday",
      "password" : "password",
      "id" : 0,
      "username" : "username"
    } ]

### Produces

This API call produces the following media types according to the Accept request header; the media type will be conveyed by the Content-Type response header.

*   `application/json`

### Responses

#### 200

OK

### Example data

Content-Type: users

    [{"id":1,"username":"alice","password":"password","birthday":"1990-01-01T00:00:00.000Z"},{"id":2,"username":"bob","password":"password","birthday":"1995-01-01T00:00:00.000Z"}]

#### 500

Internal server error [Error](#Error)

### Example data

Content-Type: error

    {"error":{"code":500,"message":"Internal server error"}}
</details>
* * *
<details><summary><code>delete /users/{id}</code></summary>

[Up](#__Methods)

    delete /users/{id}

Delete a user by ID (usersIdDelete)

### Path parameters

id (required)

Path Parameter — ID of the user to delete

### Return type

[Message](#Message)

### Example data

Content-Type: application/json

    {
      "message" : "message"
    }

### Produces

This API call produces the following media types according to the Accept request header; the media type will be conveyed by the Content-Type response header.

*   `application/json`

### Responses

#### 200

OK [Message](#Message)

### Example data

Content-Type: message

    {"message":"User deleted successfully"}

#### 404

User not found [Error](#Error)

### Example data

Content-Type: error

    {"error":{"code":404,"message":"User not foudn"}}

#### 500

Internal server error [Error](#Error)

### Example data

Content-Type: error

    {"error":{"code":500,"message":"Internal server error"}}
</details>
* * *
<details><summary><code>get /users/{id}</code></summary>

[Up](#__Methods)

    get /users/{id}

Get a user by ID (usersIdGet)

### Path parameters

id (required)

Path Parameter — ID of the user to retrieve

### Return type

[UserResponse](#UserResponse)

### Example data

Content-Type: application/json

    {
      "birthday" : "birthday",
      "password" : "password",
      "id" : 0,
      "username" : "username"
    }

### Produces

This API call produces the following media types according to the Accept request header; the media type will be conveyed by the Content-Type response header.

*   `application/json`

### Responses

#### 200

OK [UserResponse](#UserResponse)

### Example data

Content-Type: user

    {"id":1,"username":"alice","password":"password","birthday":"1990-01-01T00:00:00.000Z"}

#### 404

User not found [Error](#Error)

### Example data

Content-Type: error

    {"error":{"code":404,"message":"User not found"}}

#### 500

Internal server error [Error](#Error)

### Example data

Content-Type: error

    {"error":{"code":500,"message":"Internal server error"}}
</details>
* * *
<details><summary><code>patch /users/{id}</code></summary>

[Up](#__Methods)

    patch /users/{id}

Update a user by ID (usersIdPatch)

### Path parameters

id (required)

Path Parameter — ID of the user to update

### Consumes

This API call consumes the following media types via the Content-Type request header:

*   `application/json`

### Request body

body [User](#User) (required)

Body Parameter —

example: `{ "value" : { "username" : "bob", "password" : "newPass" } }`

### Return type

[UserResponse](#UserResponse)

### Example data

Content-Type: application/json

    {
      "birthday" : "birthday",
      "password" : "password",
      "id" : 0,
      "username" : "username"
    }

### Produces

This API call produces the following media types according to the Accept request header; the media type will be conveyed by the Content-Type response header.

*   `application/json`

### Responses

#### 200

Successfully updated the user [UserResponse](#UserResponse)

### Example data

Content-Type: user

    {"id":1,"username":"alice","password":"password","birthday":"1990-01-01T00:00:00.000Z"}

#### 400

Invalid request payload [Error](#Error)

### Example data

Content-Type: error

    {"error":{"code":400,"message":"Invalid request payload"}}

#### 404

User not found [Error](#Error)

### Example data

Content-Type: error

    {"error":{"code":404,"message":"User not found"}}

#### 500

Internal Server Error [Error](#Error)

### Example data

Content-Type: error

    {"error":{"code":500,"message":"Internal server error"}}
</details>
* * *
<details><summary><code>put /users/id</code></summary>

[Up](#__Methods)

    put /users/{id}

Update a user by ID (usersIdPut)

### Path parameters

id (required)

Path Parameter — ID of the user to update

### Consumes

This API call consumes the following media types via the Content-Type request header:

*   `application/json`

### Request body

body [User](#User) (required)

Body Parameter — User to update

example: `{ "value" : { "username" : "Bob", "password" : "newPass", "birthday" : "2000-06-06T00:00:00.000Z" } }`

### Return type

[UserResponse](#UserResponse)

### Example data

Content-Type: application/json

    {
      "birthday" : "birthday",
      "password" : "password",
      "id" : 0,
      "username" : "username"
    }

### Produces

This API call produces the following media types according to the Accept request header; the media type will be conveyed by the Content-Type response header.

*   `application/json`

### Responses

#### 200

OK [UserResponse](#UserResponse)

### Example data

Content-Type: user

    {"id":1,"username":"alice","password":"password","birthday":"1990-01-01T00:00:00.000Z"}

#### 404

User not found [Error](#Error)

### Example data

Content-Type: error

    {"error":{"code":404,"message":"User not found"}}

#### 500

Internal server error [Error](#Error)

### Example data

Content-Type: error

    {"error":{"code":500,"message":"Internal server error"}}
</details>
* * *

<details><summary><code>post /users</code></summary>

[Up](#__Methods)

    post /users

Create a new user (usersPost)

### Consumes

This API call consumes the following media types via the Content-Type request header:

*   `application/json`

### Request body

body [User](#User) (required)

Body Parameter — User to create

example: `{ "value" : { "username" : "alice", "password" : "password", "birthday" : "1990-01-01T00:00:00.000Z" } }`

### Return type

[UserResponse](#UserResponse)

### Example data

Content-Type: application/json

    {
      "birthday" : "birthday",
      "password" : "password",
      "id" : 0,
      "username" : "username"
    }

### Produces

This API call produces the following media types according to the Accept request header; the media type will be conveyed by the Content-Type response header.

*   `application/json`

### Responses

#### 201

Created [UserResponse](#UserResponse)

### Example data

Content-Type: user

    {"id":1,"username":"alice","password":"password","birthday":"1990-01-01T00:00:00.000Z"}

#### 400

Invalid request payload [Error](#Error)

### Example data

Content-Type: error

    {"error":{"code":400,"message":"Invalid request payload"}}

#### 500

Internal server error [Error](#Error)

### Example data

Content-Type: error

    {"error":{"code":500,"message":"Internal server error"}}
</details>
* * *

Models
------

\[ Jump to [Methods](#__Methods) \]

### `Error`

* error: [Error\_error](#Error_error)
    * code: [Integer](#integer)
    * message: [String](#string)
    * details (optional): [Object](#object)

### `Message` 

message: [String](#string)

### `User` 

* username (optional): [String](#string)
* password (optional): [String](#string)
* birthday (optional): [String](#string)

### `UserResponse`

* id: [Integer](#integer)
* username: [String](#string)
* password: [String](#string)
* birthday: [String](#string)


# Sprint 2 Frontend

## Primary Updates
To support the functionality of maintaining stock market trading, our opening page gives users the option to login or signup. There is also a prompt for the user to play as guest, if the user chooses to compete in cognitive brain games without the pressure of trading. After the user chooses one of the three options, they can access the homepage. This is where the user can access a variety of brain games and augment their understanding of stock market trading.

## Unit Tests / Cypress Tests
