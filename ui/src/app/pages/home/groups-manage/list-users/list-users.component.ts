import { Component, OnInit } from '@angular/core';
import {NzNotificationService} from 'ng-zorro-antd/notification';
import {Group} from '../../../../models/group';
import {ManagerService} from '../../../../services/manager.service';
import {User} from '../../../../models/user';
import {CacheService} from '../../../../services/cache.service';

@Component({
  selector: 'app-list-users',
  templateUrl: './list-users.component.html',
  styleUrls: ['./list-users.component.scss']
})
export class ListUsersComponent implements OnInit {
  group: Group = new Group();
  isVisible = false;
  users: User[] = [];
  username: string;
  me = new User();

  removeusername = '';
  alertRemove = false;
  alertCheckout = false;

  constructor(private manager: ManagerService, private notification: NzNotificationService, private cache: CacheService) {}

  showModal(group: Group): void {
    this.group = group;
    this.search('');
    this.isVisible = true;
  }

  handleOk(): void {
    this.isVisible = false;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  ngOnInit() {
    this.cache.currentUser.subscribe(u => {
      this.me = u;
    });
  }

  search(value: string): void {
    this.manager.doListGroupUsers(this.group, 0, value).then((data) => {
      this.users = data;
    });
  }

  addUser() {
    this.manager.doAddGroupUsers(this.group, this.username).then((data) => {
      this.users.push(data);
      this.notification.success('添加组用户成功', data.username);
    });
  }

  checkout() {
    this.manager.doCheckOut(this.group).then((response) => {
      this.search('');
      this.notification.success(this.group.name, '退出组成功');
    });
  }

  removeUser(username: string) {
    this.manager.doRemoveGroupUsers(this.group, username).then((data) => {
      this.search('');
      this.notification.success('删除组用户成功', data.username);
    });
  }

}
