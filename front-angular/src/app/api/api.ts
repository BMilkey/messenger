import {take, tap} from "rxjs";
import {HttpClient} from "@angular/common/http";
import {Chat, ReplyMessage, RegisterBody, SignInBody, ReplyUser, RequestMessage} from "./api-interfaces";
import {Injectable} from "@angular/core";
import {apiRepo} from "./api.repo";

@Injectable({
  providedIn: 'root'
})
export class apis {
  constructor(private http: HttpClient, private repo: apiRepo) {};

  registerUser(list: RegisterBody) {
    const url = `http://147.45.70.245:80/post/auth/register_user`;

    return this.http.post(url, list).pipe(take(1), tap((token) => this.repo.setToken(token)));
  }

  signIn(list : SignInBody) {
    const url = `http://147.45.70.245:80/post/auth/user_by_auth`;

    return this.http.post(url, list).pipe(take(1), tap((data) => this.repo.setUser(data)));
  }

  getUsersByName(list: any) {
    const url = `http://147.45.70.245:80/post/chat/user_by_name`;

    return this.http.post(url, list).pipe(take(1), tap((data) => this.repo.setUser(data)));
  }

  getChatsByToken(token :string) {
    const url = `http://147.45.70.245:80//post/chat/chats_by_token`;

    return this.http.post(url, token).pipe(take(1), tap((data) => this.repo.setChats(data)));
  }

  createChatGetUsers() {
    const url = `http://147.45.70.245:80/post/chat/create_chat_return_users`;

    return this.http.post(url, {}).pipe(take(1), tap((data) => this.repo.setReplyUsers(data)));
  }

  getUsersByChatId(id :string) {
    const url = `http://147.45.70.245:80/post/chat/users_by_chat_id`;

    return this.http.post(url, id).pipe(take(1), tap((data) => this.repo.setReplyUsers(data)));
  }

  getMessagesByChatId(id :string) {
    const url = `http://147.45.70.245:80/post/chat/messages_by_chat_id`;

    return this.http.post(url, id).pipe(take(1), tap((data) => this.repo.setReplyMessages(data)));
  }

  sendMessage(list :RequestMessage) {
    const url = `http://147.45.70.245:80/post/chat/create_message`;

    return this.http.post(url, list).pipe(take(1), tap((data) => this.repo.setReplyOnResponseMessages(data)));
  }

  addUserToChat(list :RequestUser) {
    const url = `http://147.45.70.245:80/post/chat/add_user_to_chat`;

    return this.http.post(url, list).pipe(take(1), tap((data) => this.repo.setReplyUsers(data)));
  }


}
