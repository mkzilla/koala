import {Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {CreateTaskComponent} from './pages/home/create-task/create-task.component';
import {User} from './models/user';
import {AuthService} from './services/auth.service';
import {CacheService} from './services/cache.service';
import Technique from './models/technique';
import {GiveUpTechniqueComponent} from './pages/home/give-up-technique/give-up-technique.component';
import {TechniqueComponent} from './pages/home/technique/technique.component';


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
    setInterval(() => {
      this.counter = this.transform(this.technique.createTime);
    }, 1000);
  }

  ngOnInit(): void {
    this.technique = new Technique();
    this.authService.doGetTechnique().then((data: any) => {
      if (data !== undefined) {
        this.technique = data;
        this.rain.nativeElement.play();
      }
    });
  }

  transform(value: Date): string {
    const un = (new Date(value)).getTime();
    const minute = 1000 * 60;
    const second = 1000;
    const now = new Date().getTime();
    const diffValue = now - un;
    if (diffValue < 0) {
      return;
    }
    const minC = Math.floor(diffValue / minute);
    if (minC >= 25) {
      this.rain.nativeElement.pause();
      return '已完成';
    }
    const secC = Math.floor((diffValue - minC * minute) / second);
    return ('00' + minC).slice(-2) + ':' + ('00' + secC).slice(-2);
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
