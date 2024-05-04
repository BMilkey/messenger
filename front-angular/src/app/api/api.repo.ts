import {createStore} from '@ngneat/elf';
import {addEntities, getAllEntities, withEntities} from "@ngneat/elf-entities";
import {Injectable} from "@angular/core";

interface User {
  id: string;
  name?: string;
  link?: string;
  about?: string;
  last_online?: string;
  image_id?: string;
}

@Injectable({
  providedIn: 'root'
})
export class apiRepo {
  chatStore = createStore({name: 'chat'}, withEntities<User>());

  setToken(data: any) {
    if (data.hasOwnProperty('auth_token')) {
      this.chatStore.update(addEntities({id: data.auth_token, name: data.name, link: data.link, about: data.about, last_online: data.last_online, image_id: data.image_id}));
    }
  }

  setUser(data: any) {
    if (data.hasOwnProperty('auth_token')) {
      this.chatStore.update(addEntities({id: data.auth_token, name: data.name, link: data.link, about: data.about, last_online: data.last_online, image_id: data.image_id}));
    }
    console.log(this.chatStore.query(getAllEntities()));
    console.log(typeof data.last_online);
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

