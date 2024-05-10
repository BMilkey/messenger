import {createStore} from '@ngneat/elf';
import {addEntities, getAllEntities, getEntity, withEntities} from "@ngneat/elf-entities";
import {Injectable} from "@angular/core";

interface User {
  id: string;
  name?: string;
  link?: string;
  about?: string;
  last_online?: string;
  image_id?: string;
}

interface Chat {
  id: string;
  link: string;
  title: string;
  user_id: string;
  create_time: string;
  about: string;
}

interface Message {
  id: string;
  chat_id: string;
  user_id: string;
  create_time: string;
  text: string;
  reply_msg_id: string;
}

@Injectable({
  providedIn: 'root'
})
export class apiRepo {
  usersStore = createStore({name: 'user'}, withEntities<User>());
  chatsStore = createStore({name: 'chat'}, withEntities<Chat>());
  messagesStore = createStore({name: 'message'}, withEntities<Message>());

  setToken(data: any) {
    if (data.hasOwnProperty('auth_token')) {
      this.usersStore.update(addEntities({id: data.auth_token, name: data.name, link: data.link, about: data.about, last_online: data.last_online, image_id: data.image_id}));
    }
    // console.log(this.usersStore.query(getAllEntities()));
  }

  setUser(data: any) {
    if (data.hasOwnProperty('auth_token')) {
      this.usersStore.update(addEntities({id: data.auth_token, name: data.name, link: data.link, about: data.about, last_online: data.last_online, image_id: data.image_id}));
    }
    console.log(this.usersStore.query(getAllEntities()));
  }

  setChat(data: any) {
    const array = JSON.parse(data["chats"]);
    for(let chat of array) {
      if (!this.chatsStore.query(getEntity(chat.id))) {
        this.chatsStore.update(addEntities({id: chat.id, link: chat.link, title: chat.title, user_id: chat.user_id, create_time: chat.create_time, about: chat.about}));
      }
    }
  }

  setMessages(data: any) {
    const array = JSON.parse(data["messages"]);
    for(let message of array) {
      if (!this.messagesStore.query(getEntity(message.id))) {
        this.messagesStore.update(addEntities({id: message.id, chat_id: message.chat_id, text: message.text, user_id: message.user_id, create_time: message.create_time, reply_msg_id: message.reply_msg_id}));
      }
    }
  }
  /*
  setChats(data: any) {
    if (data.hasOwnProperty('chats')) {
      this.chatStore.update(addEntities(data.chats));
  }
  console.log(this.chatStore.query(getAllEntities()));


  setReplyUsers(data: any) {
    if (data.hasOwnProperty('users')) {
      this.chatStore.update(addEntities(data.users));
    }
    console.log(this.chatStore.query(getAllEntities()));
  }

  setReplyMessages(data: any) {
    if (data.hasOwnProperty('messages')) {
      this.chatStore.update(addEntities(data.messages));
    }
    console.log(this.chatStore.query(getAllEntities()));
  }

  setReplyOnResponseMessages(data: any) {
    if (data.hasOwnProperty('message')) {
      this.chatStore.update(addEntities(data.message));
    }
    if (data.hasOwnProperty('reply_msg')) {
    console.log(this.chatStore.query(getAllEntities()));
  }
  */
}

