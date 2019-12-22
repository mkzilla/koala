import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateEditOkrComponent } from './create-edit-okr.component';

describe('CreateEditOkrComponent', () => {
  let component: CreateEditOkrComponent;
  let fixture: ComponentFixture<CreateEditOkrComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CreateEditOkrComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateEditOkrComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
