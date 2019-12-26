import { Component, OnInit } from '@angular/core';
import {User} from '../../../../models/user';
import {ManagerService} from '../../../../services/manager.service';
import {NzNotificationService} from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-create-user',
  templateUrl: './create-user.component.html',
  styleUrls: ['./create-user.component.scss']
})
export class CreateUserComponent implements OnInit {
  user: User = new User();
  isVisible = false;
  passwordVisible = false;

  constructor(private manager: ManagerService, private notification: NzNotificationService) {}

  showModal(): void {
    this.user = new User();
    this.isVisible = true;
  }

  handleOk(): void {
    this.manager.doCreateUser(this.user).then((response) => {
      this.notification.success('创建用户成功！', this.user.username);
    });
    this.isVisible = false;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  ngOnInit() {
  }

}
