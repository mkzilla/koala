import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CommentTaskComponent } from './comment-task.component';

describe('CommentTaskComponent', () => {
  let component: CommentTaskComponent;
  let fixture: ComponentFixture<CommentTaskComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CommentTaskComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CommentTaskComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
