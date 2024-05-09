import {AfterViewChecked, Component, DoCheck, OnChanges, OnInit} from '@angular/core';
import {Router, RouterLink, RouterOutlet} from '@angular/router';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, RouterLink],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements DoCheck{
  href: string = '';

  constructor(private router: Router) {}

  ngDoCheck() {
    this.href = this.router.url;
  }
}
