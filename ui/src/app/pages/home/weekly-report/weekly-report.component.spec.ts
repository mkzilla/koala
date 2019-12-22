import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WeeklyReportComponent } from './weekly-report.component';

describe('WeeklyReportComponent', () => {
  let component: WeeklyReportComponent;
  let fixture: ComponentFixture<WeeklyReportComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ WeeklyReportComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WeeklyReportComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
