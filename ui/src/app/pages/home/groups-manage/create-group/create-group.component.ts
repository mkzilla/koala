import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {NzNotificationService} from 'ng-zorro-antd/notification';
import {ManagerService} from '../../../../services/manager.service';
import {Group} from '../../../../models/group';

@Component({
  selector: 'app-create-group',
  templateUrl: './create-group.component.html',
  styleUrls: ['./create-group.component.scss']
})
export class CreateGroupComponent implements OnInit {
  @Output() refresh = new EventEmitter<Group>();
  group: Group = new Group();
  isVisible = false;

  constructor(private manager: ManagerService, private notification: NzNotificationService) {}

  showModal(): void {
    this.group = new Group();
    this.isVisible = true;
  }

  handleOk(): void {
    this.manager.doCreateGroup(this.group).then((response) => {
      if (response !== undefined) {
        this.refresh.emit(this.group);
        this.notification.success('创建组成功！', this.group.name);
      }
    });
    this.isVisible = false;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  ngOnInit() {
  }

}
