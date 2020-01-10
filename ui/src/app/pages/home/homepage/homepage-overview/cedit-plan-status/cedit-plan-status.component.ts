import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {TaskService} from '../../../../../services/task.service';
import Plan from '../../../../../models/plan';
import Status from '../../../../../models/status';

@Component({
  selector: 'app-cedit-plan-status',
  templateUrl: './cedit-plan-status.component.html',
  styleUrls: ['./cedit-plan-status.component.scss']
})
export class CeditPlanStatusComponent implements OnInit {
  @Output() refresh = new EventEmitter<string>();
  desc = '';
  isPlan = false;
  isStatus = false;
  title = '';
  placeholder = '';


  constructor(private taskService: TaskService) { }

  ngOnInit() {
  }

  handleCancel() {
    this.isPlan = false;
    this.isStatus = false;
  }

  handleOk() {
    if (this.isPlan) {
      const plan = new Plan();
      plan.desc = this.desc;
      this.taskService.doCreatePlan(plan).then((data) => {
        if (data !== undefined) {
          console.log(data);
        }
        this.refresh.emit(this.desc);
        this.isPlan = false;
      });
    } else if (this.isStatus) {
      const s = new Status();
      s.desc = this.desc;
      this.taskService.doCreateStatus(s).then((data) => {
        if (data !== undefined) {
          console.log(data);
        }
        this.refresh.emit(this.desc);
        this.isStatus = false;
      });
    }
  }
}
