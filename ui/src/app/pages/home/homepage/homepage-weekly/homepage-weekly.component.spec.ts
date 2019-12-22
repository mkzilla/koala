import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HomepageWeeklyComponent } from './homepage-weekly.component';

describe('HomepageWeeklyComponent', () => {
  let component: HomepageWeeklyComponent;
  let fixture: ComponentFixture<HomepageWeeklyComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HomepageWeeklyComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HomepageWeeklyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
