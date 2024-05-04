import { Component } from '@angular/core';
import {CreateChatModalComponent} from "./modals/create-chat-modal/create-chat-modal.component";

@Component({
  selector: 'app-chat-page',
  standalone: true,
  imports: [
    CreateChatModalComponent
  ],
  templateUrl: './chat-page.component.html',
  styleUrl: './chat-page.component.scss'
})
export class ChatPageComponent {
  openModal = false;
}
