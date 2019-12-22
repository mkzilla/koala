import {Component, Input, OnInit} from '@angular/core';
import {WeeklyService} from '../../../../services/weekly.service';
import {Reports} from '../../../../models/data/reports';
import {DateRange} from '../../../../models/data/dateRange';

@Component({
  selector: 'app-homepage-weekly',
  templateUrl: './homepage-weekly.component.html',
  styleUrls: ['./homepage-weekly.component.scss'],
})
export class HomepageWeeklyComponent implements OnInit {
  reports: Reports = new Reports();
  dateRange: DateRange = new DateRange();
  @Input() username = '';

  constructor(private week: WeeklyService) {
  }

  changDate(num: number) {
    if (num === 0) {
      this.dateRange.from.setDate(this.dateRange.from.getDate() - this.dateRange.from.getDay());
      this.dateRange.to.setDate(this.dateRange.from.getDate() + 7);
    } else {
      this.dateRange.from.setDate(this.dateRange.from.getDate() + num);
      this.dateRange.to.setDate(this.dateRange.to.getDate() + num);
    }
    this.dateRange.from.setHours(0, 0, 0);
    this.dateRange.to.setHours(0, 0, 0);
    this.week.doList(this.dateRange, this.username).then((response) => {
      this.reports = response;
    });
  }

  ngOnInit() {
    this.changDate(0);
  }

}
