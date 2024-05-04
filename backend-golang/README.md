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
"auth_token": authToken
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

- /post/chat/users_by_chat_id/
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
