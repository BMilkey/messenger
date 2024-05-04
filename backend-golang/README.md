# Backend 

## TODO

- [ ] Expiring tokens (now they are infinite)

## How to build
Windows 10 
- in .../backend-golang/
```
go mod tidy
go build main.go
```

## How to run
Windows 10
- in .../backend-golang/
```
go run main.go
```
- or
``` 
> ./messenger.exe
```

# API
## register_user
/post/auth/register_user/
### in:
```
        Login    string `json:"login"`

        Password string `json:"password"`

        Name     string `json:"name"`
```
### out:
- good
```
"auth_token": authToken
```
- error
```
"error": error
```

## user_by_auth

- /post/auth/user_by_auth/
### in:
```
        Login    string `json:"login"`

        Password string `json:"password"`
```
### out:
- good
```
	"auth_token":  auth.Auth_token,
	"name":        user.Name,
	"link":        user.Link,
	"about":       user.About,
	"last_online": user.Last_connection,
	"image_id":    user.Image_id,
```
- error
```
"error": error
```

## chats_by_token

- /post/chat/chats_by_token/
### in:
```
        Auth_token string `json:"auth_token"`
```
### out:
- good
```
"chats": string(jsonData),
```
- error
```
"error": error
```

## create_chat_return_users

- /post/chat/create_chat_return_users/
### in:
```
        Auth_token  string   `json:"auth_token"`

        Title       string   `json:"title"`

        Users_links []string `json:"users_links"`
```
### out:
- good
```
        "users": string(jsonData),
```
- error
```
"error": error
```

## users_by_name

- /post/chat/users_by_name/
### in:
```
        Name       string `json:"name"`

        Auth_token string `json:"auth_token"`
```
### out:
- good
```
        "users": string(jsonData),
```
- error
```
"error": error
```

## users_by_chat_id

- /post/chat/users_by_chat_id/
### in:
```
        ChatID     string `json:"chat_id"`

        Auth_token string `json:"auth_token"`
```
### out:
- good
```
        "users": string(jsonData),
```
- error
```
"error": error
```

## messages_by_chat_id

- /post/chat/messages_by_chat_id/
### in:
```
        ChatID     string `json:"chat_id"`

        Auth_token string `json:"auth_token"`
```
### out:
- good
```
        "messages": string(jsonData),
```
- error
```
"error": error
```

## create_message

- /post/chat/create_message/
### in:
```
        Chat_id      string `json:"chat_id"`

        Text         string `json:"text"`

        Auth_token   string `json:"auth_token"`

        Reply_msg_id string `json:"reply_msg_id"`
```
### out:
- good
```
        "message": message,

        "reply_msg": reply_msg,
```
- error
```
"error": error
```

## add_user_to_chat

- /post/chat/add_user_to_chat/
### in:
```
		Chat_id    string `json:"chat_id"`
		User_link  string `json:"user_link"`
		Auth_token string `json:"auth_token"`
```
### out:
- good
```
		"user": user,
```
- error
```
"error": error
```