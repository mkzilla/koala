import { Injectable } from '@angular/core';
import {HttpClient, HttpErrorResponse, HttpHeaders, HttpParams} from '@angular/common/http';
import { NzNotificationService } from 'ng-zorro-antd/notification';
import { Task } from '../models/task';
import { Router } from '@angular/router';
import { Comment } from '../models/comment';
import { Params } from '../models/data/params';

@Injectable({
  providedIn: 'root'
})
export class TaskService {
  headers;

  constructor(private http: HttpClient,
              private notification: NzNotificationService,
              private router: Router) {
    this.headers = new HttpHeaders({'Content-type': 'application/json'});
  }

  doCreate(data: Task): Promise<any> {
    return this.http.post('/api/v1/task', data, { headers: this.headers }).toPromise().catch( (error) => {
      this.handleError(error, '创建任务失败');
    });
  }

  doUpdate(data: Task): Promise<any> {
    return this.http.post('/api/v1/task/' + data.id, data, { headers: this.headers }).toPromise().then(
      (response: any) => {
        this.notification.success('更新任务成功', response.title);
      }
    ).catch( (error) => {
      this.handleError(error, '更新任务失败');
    });
  }

  doView(id: string): Promise<any> {
    return this.http.get('/api/v1/task/' + id, { headers: this.headers }).toPromise().catch( (error) => {
      this.handleError(error, '读取任务详情失败');
    });
  }

  doList(type?: string, state= -1, pageNumber= 0, username= ''): Promise<any> {
    const params = new HttpParams().set('type', type).set('state', '' + state).set('pageNumber', '' + pageNumber).set('username', username);
    return this.http.get('/api/v1/tasks', { headers: this.headers , params }).toPromise().catch( (error) => {
      this.handleError(error, '获取任务列表失败');
    });
  }

  doRecents(): Promise<any> {
    return this.http.get('/api/v1/tasks/recent', { headers: this.headers }).toPromise().catch( (error) => {
      this.handleError(error, '获取最近个人参与的任务失败');
    });
  }

  doKanban(): Promise<any> {
    return this.http.get('/api/v1/tasks/kanban', { headers: this.headers }).toPromise().catch( (error) => {
      this.handleError(error, '获取看板失败');
    });
  }

  doVote(data: Task): Promise<any> {
    return this.http.post('/api/v1/task/' + data.id + '/vote', {}, { headers: this.headers }).toPromise().then(
      (response: any) => {
        this.notification.success('投票成功', response.title);
      }
    ).catch( (error) => {
      this.handleError(error, '投票失败');
    });
  }
  doWatch(data: Task): Promise<any> {
    return this.http.post('/api/v1/task/' + data.id + '/watch', {}, { headers: this.headers }).toPromise().then(
      (response: any) => {
        this.notification.success('关注成功', response.title);
      }
    ).catch( (error) => {
      this.handleError(error, '关注失败');
    });
  }

  doComemnt(data: Comment): Promise<any>  {
    return this.http.post('/api/v1/task/' + data.task_id + '/comment', data, { headers: this.headers }).toPromise().then(
      (response: any) => {
        this.notification.success('添加进展/评论成功', response.title);
      }
    ).catch( (error) => {
      this.handleError(error, '添加进展/评论失败');
    });
  }

  doListComemnt(id): Promise<any>  {
    return this.http.get('/api/v1/task/' + id + '/comment', { headers: this.headers }).toPromise().
    catch( (error) => {
      this.handleError(error, '获取评论失败');
    });
  }

  doListWatch(id): Promise<any> {
    return this.http.get('/api/v1/task/' + id + '/watch', {headers: this.headers}).toPromise().catch((error) => {
      this.handleError(error, '获取关注列表失败');
    });
  }

  doListVote(id): Promise<any> {
    return this.http.get('/api/v1/task/' + id + '/vote', {headers: this.headers}).toPromise().catch((error) => {
      this.handleError(error, '获取投票列表失败');
    });
  }

  doListOKR(pageNumber= 0, username= '', type= ''): Promise<any> {
    const params = new HttpParams().set('pageNumber', '' + pageNumber).set('username', username).set('type', type);
    return this.http.get('/api/v1/okrs', { headers: this.headers , params }).toPromise().catch( (error) => {
      this.handleError(error, '获取OKR列表失败');
    });
  }

  statistics(): Promise<any> {
    return this.http.get('/api/v1/manage/statistics', {headers: this.headers}).toPromise().catch((error) => {
      this.handleError(error, '获取统计失败');
    });
  }

  handleError(error: HttpErrorResponse, message) {
    this.notification.error(message, error.error);
    if (error.status === 401) {
      window.location.href = '/login';
    }
  }
}
