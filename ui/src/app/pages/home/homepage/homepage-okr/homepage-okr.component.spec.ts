import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HomepageOkrComponent } from './homepage-okr.component';

describe('HomepageOkrComponent', () => {
  let component: HomepageOkrComponent;
  let fixture: ComponentFixture<HomepageOkrComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HomepageOkrComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HomepageOkrComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
