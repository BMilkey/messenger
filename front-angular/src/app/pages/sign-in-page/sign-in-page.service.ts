import { Injectable } from '@angular/core';
import {FormBuilder, FormControl, FormGroup, Validators} from "@angular/forms";
import {apis} from "../../api/api";
import {Observable} from "rxjs";
import {RegisterBody, SignInBody} from "../../api/api-interfaces";

@Injectable({
  providedIn: 'root'
})
export class SignInPageService {
  registerForm = new FormGroup({
    login: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required),
    name: new FormControl('', Validators.required)
  });

  signInForm = new FormGroup({
    login: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required)
  })

  constructor(private apis: apis) { }

  signUp(login: string, password: string, name: string) {
    const body: RegisterBody = {
      login: login,
      password: password,
      name: name,
    }

    this.apis.registerUser(body).subscribe();
  }

  signIn(login: string, password: string) {
    const body: SignInBody = {
      login: login,
      password: password,
    }

    this.apis.signIn(body).subscribe();
  }
}
