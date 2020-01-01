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

  doRegister(data: any) {
    this.http.post('/api/v1beta/register', data).toPromise().then(
      (response: any) => {
        this.notification.success('请登录你的邮箱从通知邮件的链接输入密码完成注册', '如果收信箱没有邮件，请检查垃圾箱');
      }
    ).catch( (error) => {
      this.notification.error('申请注册失败', error.message);
    });
  }

  doPassword(token: string, data: any): Promise<any> {
    return this.http.post('/api/v1beta/password?token=' + token, data).toPromise().then(
      (response: any) => {
        this.notification.success('请登录你的邮箱从通知邮件的链接输入密码完成注册', '如果收信箱没有邮件，请检查垃圾箱');
      }
    ).catch( (error) => {
      this.notification.error('申请注册失败', error.message);
    });
  }

  doBasicLogin(data: any, next = '/'): Promise<any> {
    return this.http.post('/api/v1beta/login', data).toPromise().then(
      (response: any) => {
        window.location.href = next;
      }
    ).catch( (error) => {
      this.notification.error('登录失败', error.message);
    });
  }

  doGetUserInfo(username: string): Promise<any> {
    return this.http.get('/api/v1/userinfo?username=' + username, {headers: this.headers}).toPromise();
  }

  doCreateTechnique(): Promise<any> {
    return this.http.post('/api/v1/technique', {}, {headers: this.headers}).toPromise();
  }

  doGetTechnique(): Promise<any> {
    return this.http.get('/api/v1/technique', {headers: this.headers}).toPromise();
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
