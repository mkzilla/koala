import {Component, ElementRef, ViewChild, OnInit} from '@angular/core';
import {CreateTaskComponent} from './pages/home/create-task/create-task.component';
import {User} from './models/user';
import {AuthService} from './services/auth.service';
import {CacheService} from './services/cache.service';
import Technique from './models/technique';
import {GiveUpTechniqueComponent} from './pages/home/give-up-technique/give-up-technique.component';
import {TechniqueComponent} from './pages/home/technique/technique.component';
import transformTime from './utils/tomatoMethod';
import {LogoutComponent} from './pages/logout/logout.component';
import {TranslateService} from '@ngx-translate/core';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  @ViewChild(CreateTaskComponent, {static: false }) createTask: CreateTaskComponent;
  @ViewChild(GiveUpTechniqueComponent, {static: false }) giveUpTechniqueComponent: GiveUpTechniqueComponent;
  @ViewChild(TechniqueComponent, {static: false }) techniqueComponent: TechniqueComponent;
  @ViewChild('rain', {static: false }) rain: ElementRef;
  @ViewChild(LogoutComponent, {static: false }) logoutComponent: LogoutComponent;

  isCollapsed = false;
  user: User = new User();
  technique = new Technique();
  counter = '加载番茄时间';
  language = 'en';

  constructor(private authService: AuthService, private cache: CacheService, public translate: TranslateService) {
    translate.addLangs(['en', 'zh-Hans']);
    this.authService.doGetUserInfo('').then((response: any) => {
      if (response !== undefined) {
        this.user = response;
        this.cache.changeUser(this.user);
      }
    });
    this.pullTechnique();
    setInterval(() => {
      this.counter = transformTime(this.technique.createTime);
      if (this.counter === '已完成') {
        this.rain.nativeElement.pause();
      }
    }, 1000);
  }

  ngOnInit() {
    const lan = localStorage.getItem('language');
    if (lan) {
      this.language = lan;
      this.translate.use(lan);
    }
    this.translate.setDefaultLang(this.language);
  }

  changeLanguage() {
    console.log(this.language);
    localStorage.setItem('language', this.language);
    this.translate.use(this.language);
  }

  pullTechnique(): void {
    this.technique = new Technique();
    this.authService.doGetTechnique().then((data: any) => {
      if (data !== undefined) {
        this.technique = data;
        this.rain.nativeElement.play();
      } else {
        this.rain.nativeElement.pause();
      }
    });
  }

  checkTechnique() {
    if (this.counter !== '已完成') {
      this.giveUpTechniqueComponent.showModal();
    } else {
      this.techniqueComponent.showModal();
    }
  }

  createTechnique() {
   if (this.technique.id !== 0) {
     return;
   }
   this.authService.doCreateTechnique().then((data) => {
     if (data !== undefined) {
       this.technique = data;
       this.rain.nativeElement.play();
     }
   });
  }
}
