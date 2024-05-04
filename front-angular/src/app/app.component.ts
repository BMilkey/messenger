import {AfterViewChecked, Component, DoCheck, OnChanges, OnInit} from '@angular/core';
import {Router, RouterLink, RouterOutlet} from '@angular/router';
import {HttpClient} from "@angular/common/http";
import {take, tap} from "rxjs";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, RouterLink],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements DoCheck{
  auth: boolean = false;
  href: string = '';

  constructor(private router: Router) {}

  ngDoCheck() {
    this.href = this.router.url;
    console.log(this.href);
  }
}
