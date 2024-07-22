package models

import (
	"time"
)

// Define models

type Auth struct {
	Login_hash    string    `json:"login_hash"`
	Password_hash string    `json:"password_hash"`
	Email         string    `json:"email"`
	User_id       string    `json:"user_id"`
	Auth_token    string    `json:"auth_token"`
	Auth_expires  time.Time `json:"auth_expires"`
}

type ChatInfo struct {
	Id          string    `json:"id"`
	Link        string    `json:"link"`
	Title       string    `json:"title"`
	User_id     string    `json:"user_id"`
	Create_time time.Time `json:"create_time"`
	About       string    `json:"about"`
	Image_id    string    `json:"image_id"`
}

type Message struct {
	Id           string    `json:"id"`
	Chat_id      string    `json:"chat_id"`
	User_id      string    `json:"user_id"`
	Create_time  time.Time `json:"create_time"`
	Text         string    `json:"text"`
	Reply_msg_id string    `json:"reply_msg_id"`
}

type User struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Link        string    `json:"link"`
	About       string    `json:"about"`
	Last_online time.Time `json:"last_online"`
	Image_id    string    `json:"image_id"`
}

type Image struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	// Size defining number of bytes
	Size uint   `json:"size"`
	Path string `json:"path"`
}

type File struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	// size defining number of bytes
	Size uint   `json:"size"`
	Path string `json:"path"`
}

type SignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Auth_token  string    `json:"auth_token"`
	Name        string    `json:"name"`
	Link        string    `json:"link"`
	About       string    `json:"about"`
	Last_online time.Time `json:"last_online"`
	Image_id    string    `json:"image_id"`
}

type SignUpRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type CreateChatRequest struct {
	Auth_token  string   `json:"auth_token"`
	Title       string   `json:"title"`
	Users_links []string `json:"users_links"`
}

type ChatUsers struct {
	Chat_id string `json:"chat_id"`
	Users   []User `json:"users"`
}

type Chats struct {
	Chats []ChatInfo `json:"chats"`
}

type ByTokenRequest struct {
	Auth_token string `json:"auth_token"`
}

type UsersByNameRequest struct {
	Name       string `json:"name"`
	Auth_token string `json:"auth_token"`
}

type Users struct {
	Users []User `json:"users"`
}

type ByChatIdRequest struct {
	Chat_id    string `json:"chat_id"`
	Auth_token string `json:"auth_token"`
}

type MessagesByChatIdRequest struct {
	Chat_id    string     `json:"chat_id"`
	From_date  *time.Time `json:"from_date"`
	To_date    *time.Time `json:"to_date"`
	Auth_token string     `json:"auth_token"`
}

type Messages struct {
	Messages []Message `json:"messages"`
}

type CreateMessageRequest struct {
	Chat_id      string `json:"chat_id"`
	Text         string `json:"text"`
	Auth_token   string `json:"auth_token"`
	Reply_msg_id string `json:"reply_msg_id"`
}

type CreateMessageResponse struct {
	Message   Message `json:"message"`
	Reply_msg Message `json:"reply_msg"`
}

type AddUserToChatRequest struct {
	Chat_id    string `json:"chat_id"`
	User_link  string `json:"user_link"`
	Auth_token string `json:"auth_token"`
}

type ChangeUserInfoRequest struct {
	Auth_token string `json:"auth_token"`
	New_name   string `json:"new_name"`
	New_link   string `json:"new_link"`
	New_about  string `json:"new_about"`
	New_image  string `json:"new_image"`
}
