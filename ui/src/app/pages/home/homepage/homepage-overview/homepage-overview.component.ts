import {Component, Input, OnInit, ViewChild} from '@angular/core';
import {TaskService} from '../../../../services/task.service';
import {Task} from '../../../../models/task';
import {AuthService} from '../../../../services/auth.service';
import {User} from '../../../../models/user';
import Objective from '../../../../models/objective';
import {CreateEditOkrComponent} from '../../create-edit-okr/create-edit-okr.component';
import {CeditPlanStatusComponent} from './cedit-plan-status/cedit-plan-status.component';
import Plan from '../../../../models/plan';
import Status from '../../../../models/status';
import {CacheService} from '../../../../services/cache.service';

@Component({
  selector: 'app-homepage-overview',
  templateUrl: './homepage-overview.component.html',
  styleUrls: ['./homepage-overview.component.scss']
})
export class HomepageOverviewComponent implements OnInit {
  @ViewChild(CeditPlanStatusComponent, {static: false }) ceditPlanStatusComponent: CeditPlanStatusComponent;
  tasks: Task[] = [];
  okrs: Objective[];
  @Input() username = '';
  user: User = new User();
  plan = new Plan();
  status = new Status();

  constructor(private taskService: TaskService, private authService: AuthService, private cache: CacheService) { }

  ngOnInit() {
    this.cache.currentUser.subscribe(u => {
      this.user = u;
    });
    this.taskService.doListOKR(0, '', 'is_doing').then((response) => {
      this.okrs = response;
    });
    this.taskService.doList('is_assign', 2, 0, this.username).then((response) => {
      this.tasks = response;
    });
    this.taskService.doGetPlan(this.username).then((response) => {
      this.plan = response;
    });
    this.taskService.doGetStatus(this.username).then((response) => {
      this.status = response;
    });
  }

  createStatus() {
    this.ceditPlanStatusComponent.title = '创建状态指标';
    this.ceditPlanStatusComponent.desc = this.status.desc;
    this.ceditPlanStatusComponent.placeholder = '状态指标是你当前工作需要核心聚焦的点，可以随着每周的工作方向变化而调整';
    this.ceditPlanStatusComponent.isStatus = true;
  }
  createPlan() {
    this.ceditPlanStatusComponent.title = '创建未来四周工作计划';
    this.ceditPlanStatusComponent.desc = this.plan.desc;
    this.ceditPlanStatusComponent.placeholder = '未来四周的工作计划既不同于任务，也和目标不同，它是一种介于目标和任务之间的状态，用于指导你的计划';
    this.ceditPlanStatusComponent.isPlan = true;
  }

}
