import {Component, Output, EventEmitter, OnInit} from '@angular/core';
import {Comment} from '../../../models/comment';
import {TaskService} from '../../../services/task.service';

@Component({
  selector: 'app-schedule-task',
  templateUrl: './schedule-task.component.html',
  styleUrls: ['./schedule-task.component.scss']
})
export class ScheduleTaskComponent implements OnInit {
  @Output() done = new EventEmitter<Comment>();
  comment: Comment = new Comment();
  isVisible = false;

  public showModal(id, type): void {
    this.comment = new Comment();
    this.comment.task_id = id;
    this.comment.type = type;
    this.isVisible = true;
  }

  handleOk(): void {
    this.taskService.doComemnt(this.comment);
    this.isVisible = false;
    this.done.emit(this.comment);
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  constructor(private taskService: TaskService) {}

  ngOnInit() {
  }

}
