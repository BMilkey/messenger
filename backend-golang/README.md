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

	var request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

**Response:**

If the login and password are valid, the response will be a JSON object with the following field:

	c.JSON(http.StatusOK, gin.H{
		"user_id": auth.User_id,
		"auth_token": auth.Auth_token,
	})

If the login and password are invalid, the response will be a JSON object with the following field:

- `error`: A description of the error that occurred (string)

## POST /post/auth/register_user/

This endpoint is used to register a new user.

**Request Body:**

The request body should be a JSON object with the following fields:

	var request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

**Response:**

If the user is successfully registered, the response will be a JSON object with the following field:

	c.JSON(http.StatusOK, gin.H{
		"user_id": newUserId,
		"auth_token": authToken,
	})

If there is an error during the registration process, the response will be a JSON object with the following field:

- `error`: A description of the error that occurred (string)

## POST /post/auth/user_by_id/

This endpoint is used to retrieve the user ID associated with a given user ID.

**Request Body:**

The request body should be a JSON object with the following field:

	var request struct {
		UserID string `json:"user_id"`
		Auth_token string `json:"auth_token"`
	}

**Response:**

If the user ID is valid, the response will be a JSON object with the following fields:

	c.JSON(http.StatusOK, gin.H{
		"id":          user.Id,
		"name":        user.Name,
		"link":        user.Link,
		"about":       user.About,
		"last_online": user.Last_connection,
		"image_id":    user.Image_id,
	})

If the user ID is invalid, the response will be a JSON object with the following field:

- `error`: A description of the error that occurred (string)

## POST /post/chat/chat_ids_by_user_id/

This endpoint is used to retrieve the chat IDs associated with a given user ID.

**Request Body:**

The request body should be a JSON object with the following field:

	var request struct {
		UserID string `json:"user_id"`
		Auth_token string `json:"auth_token"`
	}

**Response:**

If the user ID is valid, the response will be a JSON object with the following field:

	c.JSON(http.StatusOK, gin.H{
		"chat_ids": chat_ids})

If the user ID is invalid, the response will be a JSON object with the following field:

- `error`: A description of the error that occurred (string)

## POST /post/chat/user_ids_by_chat_id/

This endpoint is used to retrieve the user IDs associated with a given chat ID.

**Request Body:**

The request body should be a JSON object with the following field:

	var request struct {
		ChatID string `json:"chat_id"`
		Auth_token string `json:"auth_token"`
	}

**Response:**

If the chat ID is valid, the response will be a JSON object with the following field:

	c.JSON(http.StatusOK, gin.H{"user_ids": userIDs})

If the chat ID is invalid, the response will be a JSON object with the following field:

- `error`: A description of the error that occurred (string)
