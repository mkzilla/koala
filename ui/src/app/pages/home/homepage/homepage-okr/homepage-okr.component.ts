import {Component, Input, OnInit, ViewChild} from '@angular/core';
import {TaskService} from '../../../../services/task.service';
import {NzNotificationService} from 'ng-zorro-antd/notification';
import {TaskTimeLine} from '../../../../models/data/taskTimeLine';
import Objective from '../../../../models/objective';
import {CreateEditOkrComponent} from '../../create-edit-okr/create-edit-okr.component';
import {CacheService} from '../../../../services/cache.service';
import {User} from '../../../../models/user';

@Component({
  selector: 'app-homepage-okr',
  templateUrl: './homepage-okr.component.html',
  styleUrls: ['./homepage-okr.component.scss']
})
export class HomepageOkrComponent implements OnInit {
  @ViewChild(CreateEditOkrComponent, {static: false }) createEditOKR: CreateEditOkrComponent;
  okrs: Objective[];
  pageNumber = 0;
  @Input() username = '';
  user = new User();

  constructor(private taskService: TaskService, private notification: NzNotificationService, private cache: CacheService
  ) { }

  ngOnInit() {
    this.cache.currentUser.subscribe(u => {
      this.user = u;
    });
    this.search();
  }

  search() {
    this.taskService.doListOKR(this.pageNumber, this.username).then((response) => {
      if (this.pageNumber > 0 && (response === null || response.length === 0)) {
        this.pageNumber = this.pageNumber - 1;
        this.notification.warning('没有更多的数据', '');
        return;
      }
      this.okrs = response;
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
    this.createEditOKR.showCreateObjModalV2(this.user.id);
  }

  editObjective(okr: Objective) {
    this.createEditOKR.showEditObjModal(okr);
  }

}
