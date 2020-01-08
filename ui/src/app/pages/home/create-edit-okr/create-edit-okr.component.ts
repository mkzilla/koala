import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import Objective from '../../../models/objective';
import {ManagerService} from '../../../services/manager.service';
import {NzNotificationService} from 'ng-zorro-antd/notification';
import {TranslateService} from '@ngx-translate/core';

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

  constructor(private manager: ManagerService,
              private notification: NzNotificationService,
              public translate: TranslateService) {}

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
          this.translate.get('Objective and key results successfully created!').subscribe((s: string) => {
            this.refresh.emit(this.objective);
            this.notification.success(s, '');
          });
        }
      });
    } else {
      this.manager.doEditObjective(this.objective).then((data) => {
        this.editObj = false;
        if (data !== undefined) {
          this.translate.get('Objective and key results successfully edited!').subscribe((s: string) => {
            this.refresh.emit(this.objective);
            this.notification.success(s, '');
          });
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
