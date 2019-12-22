import {Component, OnInit, ViewChild} from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import {TaskService} from '../../../services/task.service';
import {Task} from '../../../models/task';
import {EditTaskComponent} from '../edit-task/edit-task.component';
import {ScheduleTaskComponent} from '../schedule-task/schedule-task.component';
import {CacheService} from '../../../services/cache.service';
import {User} from '../../../models/user';
import {Comment} from '../../../models/comment';
import {ListTaskCommentComponent} from '../list-task-comment/list-task-comment.component';

@Component({
  selector: 'app-view-task',
  templateUrl: './view-task.component.html',
  styleUrls: ['./view-task.component.scss']
})
export class ViewTaskComponent implements OnInit {
  @ViewChild(EditTaskComponent, {static: false }) editTask: EditTaskComponent;
  @ViewChild(ScheduleTaskComponent, {static: false }) scheduleTask: ScheduleTaskComponent;
  @ViewChild(ListTaskCommentComponent, {static: false}) listTaskComment: ListTaskCommentComponent;
  task: Task = new Task();
  user: User = new User();

  constructor(private taskService: TaskService, private route: ActivatedRoute, private cache: CacheService) {
    this.cache.currentUser.subscribe(u => {
      this.user = u;
    });
  }

  ngOnInit() {
    this.taskService.doView(this.route.snapshot.paramMap.get('id')).then((response) => {
      this.task = response;
    });
  }

  vote() {
    this.task.is_vote = !this.task.is_vote;
    this.taskService.doVote(this.task);
  }

  watch() {
    this.task.is_watch = !this.task.is_watch;
    this.taskService.doWatch(this.task);
  }

  changeState(state: number) {
    this.task.state = state;
    this.taskService.doUpdate(this.task);
  }

  refreshComment(m: Comment) {
    this.listTaskComment.add(m);
  }

}
