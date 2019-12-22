import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TaskService } from '../../../services/task.service';
import { TaskVote } from '../../../models/taskVote';
import { TaskWatch } from '../../../models/taskWatch';

@Component({
  selector: 'app-task-card',
  templateUrl: './task-card.component.html',
  styleUrls: ['./task-card.component.scss']
})
export class TaskCardComponent implements OnInit {
  votes: TaskVote[];
  watches: TaskWatch[];

  constructor(private taskService: TaskService, private route: ActivatedRoute) {
  }

  ngOnInit() {
    this.taskService.doListVote(this.route.snapshot.paramMap.get('id')).then((response) => {
        this.votes = response;
      }
    );
    this.taskService.doListWatch(this.route.snapshot.paramMap.get('id')).then((response) => {
        this.watches = response;
      }
    );

  }

}
