import { Component, OnInit } from '@angular/core';
import { TaskService } from '../../../services/task.service';
import { TaskTimeLine } from '../../../models/data/taskTimeLine';

@Component({
  selector: 'app-list-task',
  templateUrl: './list-task.component.html',
  styleUrls: ['./list-task.component.scss']
})
export class ListTaskComponent implements OnInit {
  tasks: TaskTimeLine[] = [];
  pageNumber = 0;
  hasMore = true;

  constructor(private taskService: TaskService) { }

  ngOnInit() {
    this.search();
  }

  search() {
    this.taskService.doList('', 0, this.pageNumber, '').then((response) => {
      if (response.length > 0) {
        this.tasks.push(... response);
      } else {
        this.hasMore = false;
      }
    });
  }
}
