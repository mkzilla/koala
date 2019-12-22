import {Component, Input, OnInit} from '@angular/core';
import { Comment } from '../../../models/comment';
import { TaskService } from '../../../services/task.service';
import {Task} from '../../../models/task';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-list-task-comment',
  templateUrl: './list-task-comment.component.html',
  styleUrls: ['./list-task-comment.component.scss']
})
export class ListTaskCommentComponent implements OnInit {
  comments: Comment[];

  constructor(private taskService: TaskService, private route: ActivatedRoute) {
  }

  ngOnInit() {
    this.pull();
  }

  pull() {
    this.taskService.doListComemnt(this.route.snapshot.paramMap.get('id')).then((response) => {
        this.comments = response;
      }
    );
  }

  add(m: Comment) {
    const ms = [m];
    ms.push(... this.comments);
    this.comments = ms;
  }

}
