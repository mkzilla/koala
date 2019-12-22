import {Component, ViewChild} from '@angular/core';
import {CreateTaskComponent} from './pages/home/create-task/create-task.component';
import {User} from './models/user';
import {AuthService} from './services/auth.service';
import {CacheService} from './services/cache.service';
import {Route} from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  @ViewChild(CreateTaskComponent, {static: false }) createTask: CreateTaskComponent;

  isCollapsed = false;
  user: User = new User();

  constructor(private authService: AuthService, private cache: CacheService) {
    this.authService.doGetUserInfo().then((response: any) => {
      if (response !== undefined) {
        this.user = response;
        this.cache.changeUser(this.user);
      }
    });
  }
}
