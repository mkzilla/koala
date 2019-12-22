import {Component, OnInit, ViewChild} from '@angular/core';
import {ManagerService} from '../../../services/manager.service';
import {User} from '../../../models/user';
import {NzNotificationService} from 'ng-zorro-antd/notification';
import {EditTaskComponent} from '../edit-task/edit-task.component';
import {CreateUserComponent} from './create-user/create-user.component';

@Component({
  selector: 'app-users-manage',
  templateUrl: './users-manage.component.html',
  styleUrls: ['./users-manage.component.scss']
})
export class UsersManageComponent implements OnInit {
  @ViewChild(CreateUserComponent, {static: false }) createUser: CreateUserComponent;
  username: string;
  pageNumber = 0;

  users: User[];

  constructor(private manager: ManagerService, private notification: NzNotificationService) { }

  ngOnInit() {
    this.search();
  }

  search() {
    this.manager.doListUsers(this.pageNumber, this.username).then((response) => {
      if (response.length === 0) {
        this.notification.warning('没有更多的用户数据', '');
        this.pageNumber = this.pageNumber > 0 ? this.pageNumber - 1 : this.pageNumber;
        return;
      }
      this.users = response;
    });
  }
  pagePre() {
    if (this.pageNumber === 0) {
      this.notification.warning('当前在第一页', '');
      return;
    }
    this.pageNumber = this.pageNumber - 1;
    this.search();
  }
  pagePost() {
    this.pageNumber = this.pageNumber + 1;
    this.search();
  }
  changeEnable(user: User) {
    this.manager.doUpdateUser(user).then((response) => {
      if (!response.enable) {
        this.notification.success('用户已失效，将失去系统访问能力', response.username);
      } else {
        this.notification.success('用户已生效，可正常登陆使用', response.username);
      }
    });

  }

}
