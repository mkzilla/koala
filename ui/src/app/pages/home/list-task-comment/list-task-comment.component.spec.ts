import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ListTaskCommentComponent } from './list-task-comment.component';

describe('ListTaskCommentComponent', () => {
  let component: ListTaskCommentComponent;
  let fixture: ComponentFixture<ListTaskCommentComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ListTaskCommentComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ListTaskCommentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
