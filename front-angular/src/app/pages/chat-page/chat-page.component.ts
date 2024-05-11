import {
  AfterContentChecked,
  AfterViewChecked,
  AfterViewInit,
  ChangeDetectionStrategy,
  Component,
  DoCheck,
  OnInit
} from '@angular/core';
import {ChatPageService} from "./chat-page.service";
import {getAllEntities} from "@ngneat/elf-entities";
import {NgClass} from "@angular/common";
import {demandCommandFailureMessage} from "@angular/cli/src/command-builder/utilities/command";
import {ReactiveFormsModule} from "@angular/forms";

@Component({
  selector: 'app-chat-page',
  standalone: true,
  imports: [
    NgClass,
    ReactiveFormsModule
  ],
  templateUrl: './chat-page.component.html',
  styleUrl: './chat-page.component.scss',
  changeDetection: ChangeDetectionStrategy.Default
})
export class ChatPageComponent implements OnInit, AfterViewChecked, AfterViewInit, DoCheck {
  dialog: HTMLDialogElement | undefined;
  template: Element | undefined;
  dialogInsides : Element | undefined;
  open = true;
  chats: {id: string}[] = [];
  left = false;
  right = false;
  chat:any;

  constructor(public chatService: ChatPageService) {}

  ngOnInit() {
    this.dialog = document.createElement('dialog');
    this.dialog.className = 'dialog-container__dialog';
    this.dialogInsides = document.getElementsByClassName('dialog-container__dialog__dialog-insides')[0];
    this.dialog.appendChild(this.dialogInsides);
    this.template = document.getElementsByClassName('dialog-container')[0];
    // this.chatService.getToken();
  }

  ngDoCheck() {
  }

  ngAfterViewInit() {
    this.chatService.getToken();
  }

  ngAfterViewChecked() {
    if (!sessionStorage.getItem('currentToken')) {
      this.chatService.getToken();
    }
    if (this.chats === null && this.chatService.chats.value === '' && this.chatService.currentToken.value !== undefined && this.chatService.currentToken.value !== "" || this.chats.length === 0) {
      this.chatService.getChats();
      this.chats = this.chatService.chats.value;
    }
    console.log(this.chatService.messages.value)
    if (this.chat !== null && this.chat !== undefined && this.chatService.messages.value === '') {
      this.chatService.getMessages(this.chat.id);
    }
  }

  openModal() {
    if (this.dialog !== undefined && this.template !== undefined) {
      this.template.appendChild(this.dialog);
      this.dialog.showModal();
      this.open = false;
    }
  }

  closeModal() {
    if(this.dialog !== undefined) {
      this.dialog.close();
      this.open = true;
    }
  }

  hideLeft() {
    this.left = !this.left;
  }

  hideRight() {
    this.right = !this.right;
  }

  selectChat(chatId: string) {
    this.chat = this.chats.find(item => item.id === chatId);
    this.chatService.getMessages(this.chat.id);
  }

  sendMessage() {
    console.log(this.chat)
    if (this.chat !== undefined) {
      let text = this.chatService.messageForm.getRawValue();
      this.chatService.messageForm.reset();
      const message = {
        chat_id: this.chat.id,
        text: text.text,
        auth_token: this.chatService.currentToken.value,
        reply_msg_id: "fake"
      }
      this.chatService.sendMessage(message);
      this.chatService.getMessages(this.chat.id);
    }
  }
}
