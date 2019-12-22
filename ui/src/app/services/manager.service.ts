import { Injectable } from '@angular/core';
import {NzNotificationService} from 'ng-zorro-antd/notification';
import {HttpClient, HttpErrorResponse, HttpHeaders, HttpParams} from '@angular/common/http';
import {Router} from '@angular/router';
import {User} from '../models/user';
import {Group} from '../models/group';
import Objective from '../models/objective';

@Injectable({
  providedIn: 'root'
})
export class ManagerService {
  headers;

  constructor(private http: HttpClient,
              private notification: NzNotificationService,
              private router: Router) {
    this.headers = new HttpHeaders({'Content-type': 'application/json'});
  }

  doListUsers(pageNumber= 1, username= ''): Promise<any> {
    const params = new HttpParams().set('pageNumber', '' + pageNumber).set('like', username);
    return this.http.get('/api/v1/manage/users', { headers: this.headers , params }).toPromise().catch( (error) => {
      this.handleError(error, '获取最新用户列表失败');
    });
  }

  doCreateUser(data: User): Promise<any> {
    return this.http.post('/api/v1/manage/users/', data, { headers: this.headers }).toPromise().
    catch( (error) => {
      this.handleError(error, '创建用户失败');
    });
  }

  doUpdateUser(data: User): Promise<any> {
    return this.http.put('/api/v1/manage/users/' + data.id, data, { headers: this.headers }).toPromise().
    catch( (error) => {
      this.handleError(error, '更新用户失败');
    });
  }

  doListInGroups(pageNumber= 1, name= '', enable= true): Promise<any> {
    const params = new HttpParams().set('pageNumber', '' + pageNumber).set('name', name).set('enable', '' + enable);
    return this.http.get('/api/v1/manage/groups', { headers: this.headers , params }).toPromise().catch( (error) => {
      this.handleError(error, '获取组失败');
    });
  }

  doListGroupUsers(group: Group, pageNumber= 1, name= ''): Promise<any> {
    const params = new HttpParams().set('pageNumber', '' + pageNumber).set('name', name);
    return this.http.get('/api/v1/manage/groups/' + group.id, { headers: this.headers , params }).toPromise().catch( (error) => {
      this.handleError(error, '获取组用户失败');
    });
  }

  doAddGroupUsers(group: Group, username: string): Promise<any> {
    return this.http.post('/api/v1/manage/groups/' + group.id + '/user/' + username, {}, { headers: this.headers}).
    toPromise().catch( (error) => {
      this.handleError(error, '添加组用户失败');
    });
  }

  doRemoveGroupUsers(group: Group, username: string): Promise<any> {
    return this.http.delete('/api/v1/manage/groups/' + group.id + '/user/' + username, {headers: this.headers}).
    toPromise().catch((error) => {
      this.handleError(error, '删除组用户失败');
    });
  }

  doCheckIn(group: Group): Promise<any> {
    return this.http.post('/api/v1/manage/groups/' + group.id + '/checkin', {}, {headers: this.headers}).toPromise().catch((error) => {
      this.handleError(error, '加入用户组失败');
    });
  }

  doCheckOut(group: Group): Promise<any> {
    return this.http.post('/api/v1/manage/groups/' + group.id + '/checkout', {}, {headers: this.headers}).toPromise().catch((error) => {
      this.handleError(error, '退出用户组失败');
    });
  }

  doCreateGroup(data: Group): Promise<any> {
    return this.http.post('/api/v1/manage/groups/', data, { headers: this.headers }).toPromise().
    catch( (error) => {
      this.handleError(error, '创建失败');
    });
  }

  doGetGroup(id: string): Promise<any> {
    return this.http.get('/api/v1/manage/group/' + id, { headers: this.headers }).toPromise().
    catch( (error) => {
      this.handleError(error, '获取组信息失败');
    });
  }

  doGetGroupOKRs(id: string, pageNumber: number): Promise<any> {
    const params = new HttpParams().set('pageNumber', '' + pageNumber);
    return this.http.get('/api/v1/manage/group/' + id + '/okr', { headers: this.headers, params }).toPromise().
    catch( (error) => {
      this.handleError(error, '获取组信息失败');
    });
  }

  doCreateObjective(objective: Objective): Promise<any> {
    if (objective.groupID !== 0) {
      return this.http.post('/api/v1/manage/group/' + objective.groupID + '/okr', objective,
        { headers: this.headers }).toPromise().catch((error) => {
        this.handleError(error, '创建目标失败');
        });
    }
    if (objective.userID !== 0) {
      return this.http.post('/api/v1/okr/', objective,
        { headers: this.headers }).toPromise().catch((error) => {
        this.handleError(error, '创建目标失败');
      });
    }
  }

  doEditObjective(objective: Objective): Promise<any> {
    if (objective.groupID !== 0) {
      return this.http.put('/api/v1/manage/group/' + objective.groupID + '/okr', objective,
        { headers: this.headers }).toPromise().catch((error) => {
        this.handleError(error, '编辑目标失败');
      });
    }
    if (objective.userID !== 0) {
      return this.http.put('/api/v1/okr/', objective,
        { headers: this.headers }).toPromise().catch((error) => {
        this.handleError(error, '编辑目标失败');
      });

    }
  }

  handleError(error: HttpErrorResponse, message) {
    this.notification.error(message, error.error);
    if (error.status === 401) {
      window.location.href = '/login';
    }
  }
}
