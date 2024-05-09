import {AfterViewChecked, AfterViewInit, Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {ChatPageService} from "./chat-page.service";
import {apis} from "../../api/api";

@Component({
  selector: 'app-chat-page',
  standalone: true,
  imports: [
  ],
  templateUrl: './chat-page.component.html',
  styleUrl: './chat-page.component.scss'
})
export class ChatPageComponent implements OnInit, AfterViewChecked {
  dialog: HTMLDialogElement | undefined;
  template: Element | undefined;
  dialogInsides : Element | undefined;
  open = true;

  constructor(public chatService: ChatPageService) {}

  ngOnInit() {
    this.dialog = document.createElement('dialog');
    this.dialog.className = 'dialog-container__dialog';
    this.dialogInsides = document.getElementsByClassName('dialog-container__dialog__dialog-insides')[0];
    this.dialog.appendChild(this.dialogInsides);
    this.template = document.getElementsByClassName('dialog-container')[0];
  }

  ngAfterViewChecked() {
    this.chatService.getToken();
  }

  openModal() {
    if (this.dialog !== undefined && this.template !== undefined) {
      this.template.appendChild(this.dialog);
      this.dialog.showModal();
      this.open = false;
    }
    console.log(this.dialog);
  }

  closeModal() {
    if(this.dialog !== undefined) {
      this.dialog.close();
    }
  }

  protected readonly apis = apis;
}
