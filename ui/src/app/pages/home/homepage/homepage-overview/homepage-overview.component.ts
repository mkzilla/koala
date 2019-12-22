import {Component, Input, OnInit} from '@angular/core';
import {TaskService} from '../../../../services/task.service';
import {Task} from '../../../../models/task';

@Component({
  selector: 'app-homepage-overview',
  templateUrl: './homepage-overview.component.html',
  styleUrls: ['./homepage-overview.component.scss']
})
export class HomepageOverviewComponent implements OnInit {
  tasks: Task[] = [];
  @Input() username = '';

  constructor(private taskService: TaskService) { }

  ngOnInit() {
    this.taskService.doList('is_assign', 1, 0 , this.username).then((response) => {
      this.tasks.push(...response);
    });
    this.taskService.doList('is_assign', 2, 0, this.username).then((response) => {
      this.tasks.push(...response);
    });
  }

}
