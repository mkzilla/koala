import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import Objective from '../../../models/objective';
import {ManagerService} from '../../../services/manager.service';
import {NzNotificationService} from 'ng-zorro-antd/notification';
import {Group} from '../../../models/group';

@Component({
  selector: 'app-create-edit-okr',
  templateUrl: './create-edit-okr.component.html',
  styleUrls: ['./create-edit-okr.component.scss']
})
export class CreateEditOkrComponent implements OnInit {
  @Output() refresh = new EventEmitter<Objective>();
  objective = new Objective();
  createObj = false;
  editObj = false;

  constructor(private manager: ManagerService, private notification: NzNotificationService, ) {}

  showCreateObjModal(groupID?: number): void {
    this.objective = new Objective();
    if (groupID !== 0 && groupID !== undefined && groupID !== null) {
      this.objective.groupID = groupID;
    }
    this.createObj = true;
  }

  showCreateObjModalV2(userID?: number): void {
    this.objective = new Objective();
    if (userID !== 0 && userID !== undefined && userID !== null) {
      this.objective.userID = userID;
    }
    this.createObj = true;
  }


  showEditObjModal(okr: Objective): void {
    this.objective = okr;
    this.editObj = true;
  }

  handleConfirmCreateObj(): void {
    if (this.createObj) {
      this.manager.doCreateObjective(this.objective).then((data) => {
        this.createObj = false;
        if (data !== undefined) {
          this.refresh.emit(this.objective);
          this.notification.success('创建 OKR 成功!', '');
        }
      });
    } else {
      this.manager.doEditObjective(this.objective).then((data) => {
        this.editObj = false;
        if (data !== undefined) {
          this.refresh.emit(this.objective);
          this.notification.success('编辑 OKR 成功!', '');
        }
      });
    }
  }

  handleCancel(): void {
    this.createObj = false;
    this.editObj = false;
  }

  ngOnInit() {
  }

}
