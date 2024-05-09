import { Injectable } from '@angular/core';
import {getEntitiesIds} from "@ngneat/elf-entities";
import {apiRepo} from "../../api/api.repo";
import {BehaviorSubject} from "rxjs";
import {apis} from "../../api/api";

@Injectable({
  providedIn: 'root'
})
export class ChatPageService {
  currentToken = new BehaviorSubject<string>('');

  constructor(private apiRepo: apiRepo, private api: apis) { }

  getToken() {
    let token = this.apiRepo.usersStore.query(getEntitiesIds()).at(-1);
    this.currentToken.next(token!);
  }

  createChat() {
    let body = {
      auth_token: this.currentToken.value,
      title: 'test',
      users_links: ['@a6e337e5-18f2-45ac-acf0-fc57e76e07f0']
    }
    this.api.createChatGetUsers(body).subscribe();
    this.api.getChatsByToken({auth_token: this.currentToken.value}).subscribe();
  }
}
