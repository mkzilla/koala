import { Component, OnInit } from '@angular/core';
import { TaskService } from '../../../services/task.service';
import { TaskTimeLine } from '../../../models/data/taskTimeLine';
import Objective from '../../../models/objective';
import {CacheService} from '../../../services/cache.service';
import {User} from '../../../models/user';

@Component({
  selector: 'app-list-recent-task',
  templateUrl: './list-recent-task.component.html',
  styleUrls: ['./list-recent-task.component.scss']
})
export class ListRecentTaskComponent implements OnInit {
  tasks: TaskTimeLine[];
  okrs: Objective[];
  user = new User();

  constructor(private taskService: TaskService, private cache: CacheService) {
    this.cache.currentUser.subscribe(u => {
      this.user = u;
    });
  }

  ngOnInit() {
    this.taskService.doRecents().then((response) => {
      this.tasks = response;
    });
    this.taskService.doListOKR(0, '', 'is_doing').then((response) => {
      this.okrs = response;
    });
  }
}
