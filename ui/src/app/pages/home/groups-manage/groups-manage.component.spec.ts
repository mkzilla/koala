import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupsManageComponent } from './groups-manage.component';

describe('GroupsManageComponent', () => {
  let component: GroupsManageComponent;
  let fixture: ComponentFixture<GroupsManageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GroupsManageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupsManageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
