import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ScheduleTaskComponent } from './schedule-task.component';

describe('ScheduleTaskComponent', () => {
  let component: ScheduleTaskComponent;
  let fixture: ComponentFixture<ScheduleTaskComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ScheduleTaskComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ScheduleTaskComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
