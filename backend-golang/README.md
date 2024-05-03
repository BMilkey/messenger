# TODO
- Firebase auth
- Main stuff selects/gets
- Add select/get for last N messages in chat

# How to run
```
\messenger\backend-golang> go run main.go
```
or
```
> messenger.exe
```


# How to request
## POST /post/auth/user_id_by_auth/

This endpoint is used to retrieve the user ID associated with a given login and password.

**Request Body:**

The request body should be a JSON object with the following fields:

- `login`: The login of the user (string)
- `password`: The password of the user (string)

**Response:**

If the login and password are valid, the response will be a JSON object with the following field:

- `user_id`: The ID of the user (string)

If the login and password are invalid, the response will be a JSON object with the following field:

- `error`: A description of the error that occurred (string)

## POST /post/auth/register_user/

This endpoint is used to register a new user.

**Request Body:**

The request body should be a JSON object with the following fields:

- `login`: The login of the user (string)
- `password`: The password of the user (string)
- `name`: The name of the user (string)

**Response:**

If the user is successfully registered, the response will be a JSON object with the following field:

- `user_id`: The ID of the user (string)

If there is an error during the registration process, the response will be a JSON object with the following field:

- `error`: A description of the error that occurred (string)

## POST /post/auth/login_user/

This endpoint is used to retrieve the user ID associated with a given user ID.

**Request Body:**

The request body should be a JSON object with the following field:

- `user_id`: The ID of the user (string)

**Response:**

If the user ID is valid, the response will be a JSON object with the following fields:

- `login`: The login of the user (string)
- `password`: The password of the user (string)

If the user ID is invalid, the response will be a JSON object with the following field:

- `error`: A description of the error that occurred (string)

## POST /post/auth/logout_user/

This endpoint is used to retrieve the chat IDs associated with a given user ID.

**Request Body:**

The request body should be a JSON object with the following field:

- `user_id`: The ID of the user (string)

**Response:**

If the user ID is valid, the response will be a JSON object with the following field:

- `chat_ids`: An array of chat IDs (array of strings)

If the user ID is invalid, the response will be a JSON object with the following field:

- `error`: A description of the error that occurred (string)

## POST /post/chat/chat_ids_by_user_id/

This endpoint is used to retrieve the chat IDs associated with a given user ID.

**Request Body:**

The request body should be a JSON object with the following field:

- `user_id`: The ID of the user (string)

**Response:**

If the user ID is valid, the response will be a JSON object with the following field:

- `chat_ids`: An array of chat IDs (array of strings)

If the user ID is invalid, the response will be a JSON object with the following field:

- `error`: A description of the error that occurred (string)

## POST /post/chat/user_ids_by_chat_id/

This endpoint is used to retrieve the user IDs associated with a given chat ID.

**Request Body:**

The request body should be a JSON object with the following field:

- `chat_id`: The ID of the chat (string)

**Response:**

If the chat ID is valid, the response will be a JSON object with the following field:

- `user_ids`: An array of user IDs (array of strings)

If the chat ID is invalid, the response will be a JSON object with the following field:

- `error`: A description of the error that occurred (string)