import {Component, Input, OnInit} from '@angular/core';
import {User} from '../../../models/user';
import {CacheService} from '../../../services/cache.service';
import {TaskService} from '../../../services/task.service';
import {AuthService} from '../../../services/auth.service';
import {Task} from '../../../models/task';

@Component({
  selector: 'app-edit-task',
  templateUrl: './edit-task.component.html',
  styleUrls: ['./edit-task.component.scss']
})
export class EditTaskComponent implements OnInit {
  @Input() task: Task = new Task();
  isVisible = false;
  user: User = new User();
  assignList: Array<{ value: number; text: string }> = [];

  public showModal(): void {
    this.isVisible = true;
  }

  handleOk(): void {
    this.taskService.doUpdate(this.task);
    this.isVisible = false;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  constructor(private taskService: TaskService, private cache: CacheService, private authService: AuthService) {
    this.cache.currentUser.subscribe(u => {
      this.user = u;
    });
  }

  ngOnInit() {
    this.search(this.task.user.username);
  }

  search(value: string): void {
    this.authService.doSearchUser(value).then((data) => {
      const listOfOption: Array<{ value: number; text: string }> = [];
      data.forEach(item => {
        if (item.id !== this.user.id) {
          listOfOption.push({
            value: item.id,
            text: item.nickname,
          });
        }
      });
      this.assignList = listOfOption;
    });
  }

}
