import {Component, Input, OnInit} from '@angular/core';
import {TaskService} from '../../../../services/task.service';
import {TaskTimeLine} from '../../../../models/data/taskTimeLine';
import {NzNotificationService} from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-homepage-watch',
  templateUrl: './homepage-watch.component.html',
  styleUrls: ['./homepage-watch.component.scss']
})
export class HomepageWatchComponent implements OnInit {
  tasks: TaskTimeLine[];
  pageNumber = 0;
  stateNumber = -1;
  @Input() username = '';

  constructor(private taskService: TaskService, private notification: NzNotificationService,
  ) { }

  ngOnInit() {
    this.search();
  }

  search() {
    this.taskService.doList('is_watch', this.stateNumber, this.pageNumber, this.username).then((response) => {
      if (this.pageNumber > 0 && response.length === 0) {
        this.pageNumber = this.pageNumber - 1;
        this.notification.warning('没有更多的数据', '');
        return;
      }
      this.tasks = response;
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

}
