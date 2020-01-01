import { BrowserModule } from '@angular/platform-browser';
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { IconsProviderModule } from './icons-provider.module';
import { NgZorroAntdModule, NZ_I18N, zh_CN } from 'ng-zorro-antd';
import { QuillModule } from 'ngx-quill';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { DatePipe } from '@angular/common';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { registerLocaleData } from '@angular/common';
import zh from '@angular/common/locales/zh';
import { LoginComponent } from './pages/login/login.component';
import { HomeComponent } from './pages/home/home.component';
import { CreateTaskComponent } from './pages/home/create-task/create-task.component';
import { ListTaskComponent } from './pages/home/list-task/list-task.component';
import { RelativeTimePipe } from './utils/pipes/relativeTime.pipe';
import { ViewTaskComponent } from './pages/home/view-task/view-task.component';
import {ListRecentTaskComponent} from './pages/home/list-recent-task/list-recent-task.component';
import {ScheduleTaskComponent} from './pages/home/schedule-task/schedule-task.component';
import {CommentTaskComponent} from './pages/home/comment-task/comment-task.component';
import {EditTaskComponent} from './pages/home/edit-task/edit-task.component';
import {ListTaskCommentComponent} from './pages/home/list-task-comment/list-task-comment.component';
import {TaskCardComponent} from './pages/home/task-card/task-card.component';
import {HomepageComponent} from './pages/home/homepage/homepage.component';
import {HomepageOverviewComponent} from './pages/home/homepage/homepage-overview/homepage-overview.component';
import {HomepageTasksComponent} from './pages/home/homepage/homepage-tasks/homepage-tasks.component';
import {HomepageWatchComponent} from './pages/home/homepage/homepage-watch/homepage-watch.component';
import {HomepageWeeklyComponent} from './pages/home/homepage/homepage-weekly/homepage-weekly.component';
import {HomepageVoteComponent} from './pages/home/homepage/homepage-vote/homepage-vote.component';
import {DashboardComponent} from './pages/home/dashboard/dashboard.component';
import {WeeklyReportComponent} from './pages/home/weekly-report/weekly-report.component';
import {UsersManageComponent} from './pages/home/users-manage/users-manage.component';
import {GroupsManageComponent} from './pages/home/groups-manage/groups-manage.component';
import {CreateUserComponent} from './pages/home/users-manage/create-user/create-user.component';
import {CreateGroupComponent} from './pages/home/groups-manage/create-group/create-group.component';
import {ListUsersComponent} from './pages/home/groups-manage/list-users/list-users.component';
import {StatisticsComponent} from './pages/home/statistics/statistics.component';
import {ViewGroupComponent} from './pages/home/groups-manage/view-group/view-group.component';
import {CreateEditOkrComponent} from './pages/home/create-edit-okr/create-edit-okr.component';
import {HomepageOkrComponent} from './pages/home/homepage/homepage-okr/homepage-okr.component';
import {RegisterComponent} from './pages/register/register.component';
import {TechniqueComponent} from './pages/home/technique/technique.component';
import {GiveUpTechniqueComponent} from './pages/home/give-up-technique/give-up-technique.component';

registerLocaleData(zh);

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    RegisterComponent,
    HomeComponent,
    HomepageComponent,
    HomepageOverviewComponent,
    HomepageTasksComponent,
    HomepageWatchComponent,
    HomepageVoteComponent,
    HomepageWeeklyComponent,
    HomepageOkrComponent,
    DashboardComponent,
    WeeklyReportComponent,
    UsersManageComponent,
    ListUsersComponent,
    CreateUserComponent,
    GroupsManageComponent,
    CreateGroupComponent,
    CreateTaskComponent,
    CreateEditOkrComponent,
    ListTaskComponent,
    ListTaskCommentComponent,
    ListRecentTaskComponent,
    TaskCardComponent,
    ScheduleTaskComponent,
    CommentTaskComponent,
    EditTaskComponent,
    ViewTaskComponent,
    StatisticsComponent,
    ViewGroupComponent,
    TechniqueComponent,
    GiveUpTechniqueComponent,
    RelativeTimePipe,
    TechniqueComponent,
  ],
  imports: [
    BrowserModule,
    DragDropModule,
    AppRoutingModule,
    IconsProviderModule,
    NgZorroAntdModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    ReactiveFormsModule,
    QuillModule.forRoot(),

  ],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  providers: [{ provide: NZ_I18N, useValue: zh_CN }, DatePipe],
  bootstrap: [AppComponent]
})
export class AppModule { }
