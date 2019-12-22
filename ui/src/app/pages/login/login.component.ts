import { Component, OnInit } from '@angular/core';
import * as particlesJS from 'particlesjs/dist/particles';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuthService } from '../../services/auth.service';
import Config from '../../models/config';


declare var require: any;

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  version = require('../../../../package.json').version;

  basicLogin: FormGroup;
  config: Config = new Config();


  constructor(private fb: FormBuilder, private authService: AuthService) {}

  ngOnInit(): void {
    this.basicLogin = this.fb.group({
      username: [null, [Validators.required]],
      password: [null, [Validators.required]],
    });
    particlesJS.init({
      selector: '.background',
      color: ['#DA0463', '#404B69', '#DBEDF3'],
      connectParticles: true
    });
    this.authService.doGetConfig().then((r) => {
        this.config = r;
      }
    );
  }

  onLogin() {
    for (const i of Object.keys(this.basicLogin.controls)) {
      this.basicLogin.controls[i].markAsDirty();
      this.basicLogin.controls[i].updateValueAndValidity();
    }
    this.authService.doBasicLogin(this.basicLogin.value);
  }

  onOAuth() {
    window.location.href = '/api/v1beta/login';
  }
}
