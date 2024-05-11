import { Injectable } from '@angular/core';
import {getAllEntities, getEntitiesIds} from "@ngneat/elf-entities";
import {apiRepo} from "../../api/api.repo";
import {BehaviorSubject} from "rxjs";
import {apis} from "../../api/api";
import {FormControl, FormGroup, Validators} from "@angular/forms";

@Injectable({
  providedIn: 'root'
})
export class ChatPageService {
  currentToken = new BehaviorSubject<any>(sessionStorage.getItem('currentToken'));
  chats = new BehaviorSubject<any>('');
  messages = new BehaviorSubject<any>('');
  you = new BehaviorSubject<any>('');

  messageForm = new FormGroup({
    text: new FormControl('', Validators.required),
  });

  newChatForm = new FormGroup({
    title: new FormControl('', Validators.required),
    name: new FormControl('', Validators.required),
  });

  constructor(public apiRepo: apiRepo, private api: apis) { }

  getToken() {
    // console.log(this.apiRepo.usersStore.query(getEntitiesIds()))
    let token = this.apiRepo.usersStore.query(getEntitiesIds()).at(-1);
    if (token !== undefined) {
      sessionStorage.setItem('currentToken', token);
      this.currentToken.next(token);

      let users = this.apiRepo.usersStore.query(getAllEntities());
      let you = users.find(item => item.id === this.currentToken.value);
      if (you !== undefined) {
        this.you.next(you);
      }
    }
  }

  createChat() {
    if (this.newChatForm.getRawValue().title !== '' && this.newChatForm.getRawValue().name !== '') {
      this.getUsersByName();
      let friend = this.apiRepo.usersStore.query(getAllEntities()).find(item => item.name === this.newChatForm.getRawValue().name);
      let body = {
        auth_token: this.currentToken.value,
        title: this.newChatForm.getRawValue().title,
        users_links: friend?.link
      }
      this.api.createChatGetUsers(body).subscribe();
      this.getChats();
    }
  }

  getChats() {
    if(this.currentToken.value === Array(0) || sessionStorage.getItem('currentToken') !== null && this.currentToken.value !== undefined && this.currentToken.value !== null && this.currentToken.value !== "") {
      this.api.getChatsByToken({auth_token: this.currentToken.value}).subscribe();
      let data = this.apiRepo.chatsStore.query(getAllEntities());
      this.chats.next(data);
    }
  }

  sendMessage(message: any) {
    this.api.sendMessage(message).subscribe();
  }

  getMessages(id: string) {
    this.api.getMessagesByChatId({chat_id: id, auth_token: this.currentToken.value}).subscribe();
    let data = this.apiRepo.messagesStore.query(getAllEntities()).filter(item => item.chat_id === id);
    this.messages.next(data);
    console.log(data)
  }

  getUsersByName() {
    if (this.newChatForm.getRawValue().name !== '') {
      let name = this.newChatForm.getRawValue().name;
      let body = {
        name: name,
        auth_token: this.currentToken.value
      }
      this.api.getUsersByName(body).subscribe();
    }
  }
}
