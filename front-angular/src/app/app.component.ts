import { Component } from '@angular/core';
import {RouterLink, RouterOutlet} from '@angular/router';
import {HttpClient} from "@angular/common/http";
import {take, tap} from "rxjs";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, RouterLink],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  auth: boolean = false;
  response: any;

  constructor(private http: HttpClient) { }

  registerUserRequest() {
    const url = `http://147.45.70.245:80/post/chat/chat_ids_by_user_id/`;

    return this.http.post(url, {user_id: 'kHBrjINqoIRPuG3ACxf5XFtQdhj1'}).pipe(take(1), tap((response) => this.setData(response)));
  }

  registerUser() {
    this.registerUserRequest().subscribe();
  }

  setData(data: any) {
    this.response = data;
    console.log(this.response);
  }
}
