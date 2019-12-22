import { Component, OnInit } from '@angular/core';
import {Reports} from '../../../models/data/reports';
import {DateRange} from '../../../models/data/dateRange';
import {WeeklyService} from '../../../services/weekly.service';
import {AuthService} from '../../../services/auth.service';
import {Group} from '../../../models/group';
import {ManagerService} from '../../../services/manager.service';

@Component({
  selector: 'app-weekly-report',
  templateUrl: './weekly-report.component.html',
  styleUrls: ['./weekly-report.component.scss']
})
export class WeeklyReportComponent implements OnInit {
  reportGroup: Reports[] = [];
  groups: Group[] = [];
  groupId: number;
  dateRange: DateRange = new DateRange();

  constructor(private week: WeeklyService, private auth: ManagerService) {
  }

  changDate(num: number) {
    this.dateRange.from.setDate(this.dateRange.from.getDate() + num);
    this.dateRange.to.setDate(this.dateRange.to.getDate() + num);
    this.week.doListGroup(this.groupId, this.dateRange).then((response) => {
      this.reportGroup = response;
    });
  }

  ngOnInit() {
    this.dateRange.from.setDate(this.dateRange.from.getDate() - this.dateRange.from.getDay());
    this.dateRange.from.setHours(0, 0, 0);
    this.dateRange.to.setDate(this.dateRange.from.getDate() + 7);
    this.dateRange.to.setHours(0, 0, 0);
    this.auth.doListInGroups(0, '', true).then((grps) => {
      this.groups = grps;
      this.groupId = this.groups[0].id;
      this.week.doListGroup(this.groupId, this.dateRange).then((response) => {
        this.reportGroup = response;
      });
    });
  }

}
