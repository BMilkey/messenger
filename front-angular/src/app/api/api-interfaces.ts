import { Time } from "@angular/common";
import { Timestamp } from "rxjs";

export interface RegisterBody {
  login: string;
  password: string;
  name: string;
}

export interface SignInBody {
  login: string;
  password: string;
}

export interface Chat {
  id: string;
  link: string;
  title: string;
  user_id: string;
  create_time: Date;
  about: string;
  image_id: string;
}

export interface User {
  id: string;
  name: string;
  link: string;
  about: string;
  last_connection: string;
  image_id: string;
}

export interface Message {
  id: string;
  chat_id: string;
  user_id: string;
  create_time: string;
  text: string;
  reply_msg_id: string;
}
