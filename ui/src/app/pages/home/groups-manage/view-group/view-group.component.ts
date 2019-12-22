import {Component, OnInit, ViewChild} from '@angular/core';
import {ManagerService} from '../../../../services/manager.service';
import {NzNotificationService} from 'ng-zorro-antd/notification';
import {ActivatedRoute} from '@angular/router';
import {Group} from '../../../../models/group';
import Objective from '../../../../models/objective';
import {CreateEditOkrComponent} from '../../create-edit-okr/create-edit-okr.component';
import {CreateGroupComponent} from '../create-group/create-group.component';

@Component({
  selector: 'app-view-group',
  templateUrl: './view-group.component.html',
  styleUrls: ['./view-group.component.scss']
})
export class ViewGroupComponent implements OnInit {
  @ViewChild(CreateEditOkrComponent, {static: false }) createEditOKR: CreateEditOkrComponent;
  group = new Group();
  okrs: Objective[] = [];
  pageNumber = 0;

  constructor(private manager: ManagerService,
              private route: ActivatedRoute,
              private notification: NzNotificationService) { }

  ngOnInit() {
    this.manager.doGetGroup(this.route.snapshot.paramMap.get('id')).then((data) => {
      if (data !== undefined) {
        this.group = data;
      }
    });
    this.search();

  }

  search() {
    this.manager.doGetGroupOKRs(this.route.snapshot.paramMap.get('id'), this.pageNumber).then((data) => {
      if (data !== undefined) {
        this.okrs = data;
      } else if (this.pageNumber > 0) {
        this.pageNumber = this.pageNumber - 1;
        this.notification.warning('没有更多的数据', '');
        return;
      }
    });
  }

  changePage(num: number) {
    if (this.pageNumber + num < 0) {
      this.notification.warning('当前已经是第一页', '');
      return;
    }
    this.pageNumber = this.pageNumber + num;
    this.search();
  }

  createObjective() {
    this.createEditOKR.showCreateObjModal(this.group.id);
  }

  editObjective(okr: Objective) {
    this.createEditOKR.showEditObjModal(okr);
  }
}
