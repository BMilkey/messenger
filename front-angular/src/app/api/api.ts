import {take, tap} from "rxjs";
import {HttpClient} from "@angular/common/http";
import {RegisterBody, SignInBody} from "./api-interfaces";
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
}
