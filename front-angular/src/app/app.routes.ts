import { Routes } from '@angular/router';
import {SettingsPageComponent} from "./pages/settings-page/settings-page.component";
import {ChatPageComponent} from "./pages/chat-page/chat-page.component";
import {SignInPageComponent} from "./pages/sign-in-page/sign-in-page.component";
import {authGuard} from "./guards/auth.guard";

export const routes: Routes = [
  {path: '', redirectTo: '/auth', pathMatch: 'full'},
  {path: 'auth', component: SignInPageComponent},
  {path: 'chat', component: ChatPageComponent, canActivate: [authGuard]},
  {path: 'settings', component: SettingsPageComponent, canActivate: [authGuard]},
  {path: '**', redirectTo: '/auth', pathMatch: 'full'},
];
