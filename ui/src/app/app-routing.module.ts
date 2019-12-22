import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { HomeComponent } from './pages/home/home.component';
import { ViewTaskComponent } from './pages/home/view-task/view-task.component';
import {HomepageComponent} from './pages/home/homepage/homepage.component';
import {DashboardComponent} from './pages/home/dashboard/dashboard.component';
import {WeeklyReportComponent} from './pages/home/weekly-report/weekly-report.component';
import {UsersManageComponent} from './pages/home/users-manage/users-manage.component';
import {GroupsManageComponent} from './pages/home/groups-manage/groups-manage.component';
import {StatisticsComponent} from './pages/home/statistics/statistics.component';
import {ViewGroupComponent} from './pages/home/groups-manage/view-group/view-group.component';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'login', component: LoginComponent },
  { path: 'task/:id', component: ViewTaskComponent },
  { path: 'homepage/:username', component: HomepageComponent },
  { path: 'dashboard', component: DashboardComponent },
  { path: 'weekly', component: WeeklyReportComponent },
  { path: 'work-group', component: GroupsManageComponent },
  { path: 'work-group/:id', component: ViewGroupComponent },
  { path: 'users-manage', component: UsersManageComponent },
  { path: 'statistics', component: StatisticsComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
