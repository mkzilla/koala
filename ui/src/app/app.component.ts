import {Component, ElementRef, ViewChild} from '@angular/core';
import {CreateTaskComponent} from './pages/home/create-task/create-task.component';
import {User} from './models/user';
import {AuthService} from './services/auth.service';
import {CacheService} from './services/cache.service';
import Technique from './models/technique';
import {GiveUpTechniqueComponent} from './pages/home/give-up-technique/give-up-technique.component';
import {TechniqueComponent} from './pages/home/technique/technique.component';
import transformTime from './utils/tomatoMethod';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  @ViewChild(CreateTaskComponent, {static: false }) createTask: CreateTaskComponent;
  @ViewChild(GiveUpTechniqueComponent, {static: false }) giveUpTechniqueComponent: GiveUpTechniqueComponent;
  @ViewChild(TechniqueComponent, {static: false }) techniqueComponent: TechniqueComponent;
  @ViewChild('rain', {static: false }) rain: ElementRef;

  isCollapsed = false;
  user: User = new User();
  technique = new Technique();
  counter = '加载番茄时间';

  constructor(private authService: AuthService, private cache: CacheService) {
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
