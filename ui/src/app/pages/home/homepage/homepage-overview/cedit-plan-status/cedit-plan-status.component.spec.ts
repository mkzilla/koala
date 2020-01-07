import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CeditPlanStatusComponent } from './cedit-plan-status.component';

describe('CeditPlanStatusComponent', () => {
  let component: CeditPlanStatusComponent;
  let fixture: ComponentFixture<CeditPlanStatusComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CeditPlanStatusComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CeditPlanStatusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
