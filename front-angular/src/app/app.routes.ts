import { Routes } from '@angular/router';
import {SettingsPageComponent} from "./pages/settings-page/settings-page.component";
import {ChatPageComponent} from "./pages/chat-page/chat-page.component";
import {SignInPageComponent} from "./pages/sign-in-page/sign-in-page.component";

export const routes: Routes = [
  {path: '', redirectTo: '/auth', pathMatch: 'full'},
  {path: 'auth', component: SignInPageComponent},
  {path: 'chat', component: ChatPageComponent},
  {path: 'settings', component: SettingsPageComponent},
  {path: '**', redirectTo: '/auth', pathMatch: 'full'},
];
