import { Component } from '@angular/core';
import {SignInPageService} from "./sign-in-page.service";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {apis} from "../../api/api";
import {Router, Routes} from "@angular/router";
import {ChatPageService} from "../chat-page/chat-page.service";

@Component({
  selector: 'app-sign-in-page',
  standalone: true,
  imports: [
    FormsModule,
    ReactiveFormsModule,
  ],
  templateUrl: './sign-in-page.component.html',
  styleUrl: './sign-in-page.component.scss'
})
export class SignInPageComponent {
  pageState = 'menu';
  regForm = this.formService.registerForm;
  signInForm = this.formService.signInForm;

  constructor(private formService: SignInPageService, private router: Router, private chatService: ChatPageService) {}

  register() {
    const data = this.regForm.getRawValue();
    if (data.login !== '' && data.password !== '' && data.name !== '' && data.name !== null && data.login !== null && data.password !== null) {
      this.formService.signUp(data.login, data.password, data.name);
      this.pageState = 'menu';
      this.chatService.getToken();
      this.router.navigateByUrl('/chat');
    } else {
      alert("You've written some shit:( Try again");
    }
  }

  signIn() {
    const data = this.signInForm.getRawValue();
    if (data.login !== '' && data.password !== '' && data.login !== null && data.password !== null) {
      this.formService.signIn(data.login, data.password);
      this.pageState = 'menu';
      this.chatService.getToken();
      this.router.navigateByUrl('/chat');
    } else {
      alert("You've written some shit:( Try again");
    }
  }
}
