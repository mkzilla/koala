import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ListRecentTaskComponent } from './list-recent-task.component';

describe('ListRecentTaskComponent', () => {
  let component: ListRecentTaskComponent;
  let fixture: ComponentFixture<ListRecentTaskComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ListRecentTaskComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ListRecentTaskComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
