import {Component, OnInit, ViewChild} from '@angular/core';
import {NzNotificationService} from 'ng-zorro-antd/notification';
import {ManagerService} from '../../../services/manager.service';
import {Group} from '../../../models/group';
import {CreateGroupComponent} from '../groups-manage/create-group/create-group.component';
import {ListUsersComponent} from '../groups-manage/list-users/list-users.component';

@Component({
  selector: 'app-groups-manage',
  templateUrl: './groups-manage.component.html',
  styleUrls: ['./groups-manage.component.scss']
})
export class GroupsManageComponent implements OnInit {
  @ViewChild(CreateGroupComponent, {static: false }) createGroup: CreateGroupComponent;
  @ViewChild(ListUsersComponent, {static: false }) listUsers: ListUsersComponent;
  name: string;
  pageNumber = 0;

  groups: Group[];

  forcheck: Group[];

  constructor(private manager: ManagerService, private notification: NzNotificationService) { }

  ngOnInit() {
    this.search();
    this.searchForCheck();
  }

  searchForCheck() {
    this.manager.doListInGroups(this.pageNumber, this.name, false).then((response) => {
      this.forcheck = response;
    });
  }

  search() {
    this.manager.doListInGroups(this.pageNumber, this.name, true).then((response) => {
      if (this.pageNumber > 0 && response.length === 0) {
        this.notification.warning('没有更多的组数据', '');
        this.pageNumber = this.pageNumber - 1;
        return;
      }
      this.groups = response;
    });
  }

  pageChange(num= 0) {
    if (this.pageNumber === 0 && num < 0) {
      this.notification.warning('当前在第一页', '');
      return;
    }
    this.pageNumber = this.pageNumber + num;
    this.search();
  }

  checkin(group: Group) {
    this.manager.doCheckIn(group).then((response) => {
      this.search();
      this.searchForCheck();
      this.notification.success(group.name, '加入用户组成功');
    });
  }

  ignore(group: Group) {
    this.manager.doCheckOut(group).then((response) => {
      this.notification.success(group.name, '忽略邀请成功');
    });
  }
}

