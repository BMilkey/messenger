import {AfterViewChecked, Component, DoCheck, OnChanges, OnInit} from '@angular/core';
import {Router, RouterLink, RouterOutlet} from '@angular/router';
import {NgClass} from "@angular/common";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, RouterLink, NgClass],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements DoCheck{
  href: string = '';
  chat = true;
  settings = false;

  constructor(private router: Router) {}

  ngDoCheck() {
    this.href = this.router.url;
  }

  selectChat() {
    this.chat = true;
    this.settings = false;
  }

  selectSettings() {
    this.chat = false;
    this.settings = true;
  }

  logOut(){
    sessionStorage.clear();
    console.log(sessionStorage);
  }
}
