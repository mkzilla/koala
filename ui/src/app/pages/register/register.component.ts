import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup} from '@angular/forms';
import {AuthService} from '../../services/auth.service';
import {User} from '../../models/user';
import {ActivatedRoute, Router} from '@angular/router';
import {NzNotificationService} from 'ng-zorro-antd/notification';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {
  user = new User();
  username = '';
  token = '';
  password = '';
  email = '';

  passwordVisible = false;


  constructor(private fb: FormBuilder, private authService: AuthService,
              private route: ActivatedRoute, private notification: NzNotificationService, private router: Router) {
  }

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.token = params.token;
    });
    this.route.params.subscribe(params => {
      this.username = params.username;
    });
  }

  onRegister() {
    this.authService.doRegister(this.user);
  }

  onPassword() {
    if (this.password !== this.user.password) {
      this.notification.warning('两次输入的密码不一致', '请重新输入');
      return;
    }
    if (this.user.password.length < 8) {
      this.notification.warning('密码长度至少为 8 位', '请确保密码安全性');
      return;
    }
    this.authService.doPassword(this.token, this.user).then(() => {
      this.notification.success('密码设置成功', '请使用用户名 ' + this.username + ' 完成登录！');
      this.router.navigate([`/login`]);
      }
    );
  }

}
