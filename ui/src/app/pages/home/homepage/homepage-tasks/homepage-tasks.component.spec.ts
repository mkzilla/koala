import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HomepageTasksComponent } from './homepage-tasks.component';

describe('HomepageTasksComponent', () => {
  let component: HomepageTasksComponent;
  let fixture: ComponentFixture<HomepageTasksComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HomepageTasksComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HomepageTasksComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
