import {Component, Input, OnInit} from '@angular/core';
import {TaskService} from '../../../../services/task.service';
import {Task} from '../../../../models/task';
import {AuthService} from '../../../../services/auth.service';
import {User} from '../../../../models/user';

@Component({
  selector: 'app-homepage-overview',
  templateUrl: './homepage-overview.component.html',
  styleUrls: ['./homepage-overview.component.scss']
})
export class HomepageOverviewComponent implements OnInit {
  tasks: Task[] = [];
  @Input() username = '';
  user: User = new User();

  constructor(private taskService: TaskService, private authService: AuthService) { }

  ngOnInit() {
    this.authService.doGetUserInfo(this.username).then((response: any) => {
      if (response !== undefined) {
        this.user = response;
      }
    });
    this.taskService.doList('is_assign', 1, 0 , this.username).then((response) => {
      this.tasks.push(...response);
    });
    this.taskService.doList('is_assign', 2, 0, this.username).then((response) => {
      this.tasks.push(...response);
    });
  }

}
