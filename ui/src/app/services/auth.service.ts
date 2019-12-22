import { Injectable } from '@angular/core';
import {HttpClient, HttpErrorResponse, HttpHeaders} from '@angular/common/http';
import { NzNotificationService } from 'ng-zorro-antd/notification';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  token: string;
  headers;

  constructor(private http: HttpClient,
              private notification: NzNotificationService,
              private router: Router) {
      this.headers = new HttpHeaders({'Content-type': 'application/json'});
  }

  doGetConfig(): Promise<any> {
    return this.http.get('/api/v1beta/config', {headers: this.headers}).toPromise();
  }

  doBasicLogin(data: any, next = '/'): Promise<any> {
    return this.http.post('/api/v1beta/login', data).toPromise().then(
      (response: any) => {
        localStorage.setItem('jsonwebtoken', response);
        window.location.href = next;
      }
    ).catch( (error) => {
      this.notification.error('登录失败', error.message);
    });
  }

  doGetUserInfo(): Promise<any> {
    return this.http.get('/api/v1/userinfo', {headers: this.headers}).toPromise().catch( (error) => {
      this.notification.error('登录失败', error.message);
    });
  }

  doSearchUser(username: string): Promise<any> {
    return this.http.get(`/api/v1/users/assign?like=${username}`, {headers: this.headers}).toPromise().catch((error) => {
      this.handleError(error, '获取查询用户列表失败');
    });
  }

  handleError(error: HttpErrorResponse, message) {
    this.notification.error(message, error.message);
    if (error.status === 401) {
      window.location.href = '/login';
    }
  }
}
