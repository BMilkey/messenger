package models

import (
	"time"
)

// Define models

type Auth struct {
	Login_hash string `json:"login_hash"`
	Password_hash string `json:"password_hash"`
	Email string `json:"email"`
	User_id string `json:"user_id"`
	Auth_token string `json:"auth_token"`
	Auth_expires time.Time `json:"auth_expires"`
}

type Chat struct {
	Id          string    `json:"id"`
	Link        string    `json:"link"`
	Title       string    `json:"title"`
	User_id     string    `json:"user_id"`
	Create_time time.Time `json:"create_time"`
	About       string    `json:"about"`
	Image_id    string   `json:"image_id"`
}

type Message struct {
	Id           string    `json:"id"`
	Chat_id      string    `json:"chat_id"`
	User_id      string    `json:"user_id"`
	Create_time  time.Time `json:"create_time"`
	Text         string    `json:"text"`
	Reply_msg_id string   `json:"reply_msg_id"`
}

type User struct {
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	Link            string    `json:"link"`
	About           string    `json:"about"`
	Last_connection time.Time `json:"last_connection"`
	Image_id        string   `json:"image_id"`
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


