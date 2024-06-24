import {CanActivateFn, Router} from '@angular/router';
import {ChatPageService} from "../pages/chat-page/chat-page.service";
import {inject} from "@angular/core";
import {apiRepo} from "../api/api.repo";
import {getAllEntities} from "@ngneat/elf-entities";
import {SignInPageService} from "../pages/sign-in-page/sign-in-page.service";

export const authGuard: CanActivateFn = (route, state) => {
  const authService = inject(ChatPageService);
  const apiStores = inject(apiRepo);
  const router = inject(Router);
  const signService = inject(SignInPageService);
  return true;
  if (apiStores.usersStore.query(getAllEntities()).find(item => item.name === signService.signInForm.getRawValue().login)) {
    return true;
  } else {
    console.log(apiStores.usersStore.query(getAllEntities()).find(item => item.name === signService.signInForm.getRawValue().login))
    console.log(apiStores.usersStore.query(getAllEntities()))
    router.navigate(['/auth']);
    return false;
  }
};
