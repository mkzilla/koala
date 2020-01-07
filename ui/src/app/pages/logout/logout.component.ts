import { Component } from '@angular/core';
import {CacheService} from '../../services/cache.service';
import {User} from '../../models/user';

@Component({
  selector: 'app-logout',
  templateUrl: './logout.component.html',
  styleUrls: ['./logout.component.scss']
})
export class LogoutComponent {
  user = new User();
  isVisible = false;

  constructor(private cache: CacheService) {
    this.cache.currentUser.subscribe(u => {
      this.user = u;
    });
  }

  showModal(): void {
    this.isVisible = true;
  }

  handleOk(): void {
    const keys = document.cookie.match(/[^ =;]+(?=\=)/g);
    if (keys) {
      for (let i = keys.length; i--;) {
        document.cookie = keys[i] + '=0;expires=' + new Date(0).toUTCString();
      }
    }
    this.isVisible = false;
    window.location.href = '/login';
  }

  handleCancel(): void {
    this.isVisible = false;
  }

}
