import { Injectable } from '@angular/core';
import {HttpClient, HttpErrorResponse, HttpHeaders, HttpParams} from '@angular/common/http';
import {Router} from '@angular/router';
import {NzNotificationService} from 'ng-zorro-antd/notification';
import {DateRange} from '../models/data/dateRange';
import { DatePipe } from '@angular/common';

@Injectable({
  providedIn: 'root',
})
export class WeeklyService {
  headers;

  constructor(private http: HttpClient,
              private notification: NzNotificationService,
              private router: Router, private datepipe: DatePipe) {
    this.headers = new HttpHeaders({'Content-type': 'application/json' });
  }

  doList(dateRange: DateRange, username= ''): Promise<any> {
    const params = new HttpParams().set('from',
      this.datepipe.transform(dateRange.from, 'yyyy-MM-ddTHH:mm:ss')
    ).set('to',
      this.datepipe.transform(dateRange.to, 'yyyy-MM-ddTHH:mm:ss'))
      .set('username', username);
    return this.http.get('/api/v1/weekly/self', { headers: this.headers, params}).toPromise().catch( (error) => {
      this.handleError(error, '获取周报失败');
    });
  }

  doListGroup(id: number, dateRange: DateRange): Promise<any> {
    const params = new HttpParams().set('from',
      this.datepipe.transform(dateRange.from, 'yyyy-MM-ddTHH:mm:ss')
    ).set('to',
      this.datepipe.transform(dateRange.to, 'yyyy-MM-ddTHH:mm:ss'));
    return this.http.get('/api/v1/weekly/groups/' + id, { headers: this.headers, params}).toPromise().catch( (error) => {
      this.handleError(error, '获取所在的用户组失败');
    });
  }

  handleError(error: HttpErrorResponse, message) {
    this.notification.error(message, error.error);
    if (error.status === 401) {
      window.location.href = '/login';
    }
  }
}
