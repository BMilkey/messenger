# How to run
```
\messenger\backend-golang> go run main.go
```
or
```
> messenger.exe
```


# How to request
get_chats (user_id):
Send a GET request to http://localhost:8080/get_chats/{user_id}, replacing {user_id} with an actual user ID.
Example: http://localhost:8080/get_chats/123

get_messages(chat_id):
Send a GET request to http://localhost:8080/get_messages/{chat_id}, replacing {chat_id} with an actual chat ID.
Example: http://localhost:8080/get_messages/456

get_chat_participants(chat_id):
Send a GET request to http://localhost:8080/get_chat_participants/{chat_id}, replacing {chat_id} with an actual chat ID.
Example: http://localhost:8080/get_chat_participants/789

get_images(list of image_id):
Send a POST request to http://localhost:8080/get_images with a JSON body containing the list of image IDs.
Example request body: {"image_ids": ["1", "2", "3"]}

get_files(list of file_id):
Send a POST request to http://localhost:8080/get_files with a JSON body containing the list of file IDs.
Example request body: {"file_ids": ["4", "5", "6"]}