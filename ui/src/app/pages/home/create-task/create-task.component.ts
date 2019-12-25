import {Component, OnInit} from '@angular/core';
import {Task} from '../../../models/task';
import {TaskService} from '../../../services/task.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-create-task',
  templateUrl: './create-task.component.html',
  styleUrls: ['./create-task.component.scss']
})
export class CreateTaskComponent implements OnInit {
  task: Task = new Task();
  isVisible = false;

  public showModal(): void {
    this.isVisible = true;
  }

  handleOk(): void {
    this.taskService.doCreate(this.task).then((id) => {
      this.isVisible = false;
      this.task = new Task();
      this.route.navigate([`/task/${id}`]);
    });
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  constructor(private taskService: TaskService, private route: Router) {
  }

  ngOnInit() {
  }

}
