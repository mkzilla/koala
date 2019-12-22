import { Component, OnInit } from '@angular/core';
import {CdkDragDrop, moveItemInArray, transferArrayItem} from '@angular/cdk/drag-drop';
import { NzNotificationService } from 'ng-zorro-antd';

import { Kanban } from '../../../models/data/kanban';
import { Task } from '../../../models/task';
import { TaskService } from '../../../services/task.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  kanban: Kanban = new Kanban();

  constructor(private taskService: TaskService, private notification: NzNotificationService) { }

  ngOnInit() {
    this.taskService.doKanban().then((response) => {
      if (response.will === null) {
        response.will = [];
      }
      if (response.doing === null) {
        response.doing = [];
      }
      if (response.done === null) {
        response.done = [];
      }
      this.kanban = response;
    });
  }

  drop(event: CdkDragDrop<Task[]>) {
    moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
    if (event.previousContainer !== event.container) {
      transferArrayItem(event.previousContainer.data,
        event.container.data,
        event.previousIndex,
        event.currentIndex);
    }
    for (const w of this.kanban.will) {
      if (w.state !== 1 ) {
        w.state = 1;
        this.taskService.doUpdate(w).then(() => {
          return;
        });
      }
    }
    for (const w of this.kanban.done) {
      if (w.state !== 3 ) {
        w.state = 3;
        this.taskService.doUpdate(w).then(() => {
          return;
        });
      }
    }
    for (const w of this.kanban.doing) {
      if (w.state !== 2 ) {
        w.state = 2;
        this.taskService.doUpdate(w).then(() => {
          return;
        });
      }
    }
  }
}
