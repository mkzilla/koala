import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HomepageOverviewComponent } from './homepage-overview.component';

describe('HomepageOverviewComponent', () => {
  let component: HomepageOverviewComponent;
  let fixture: ComponentFixture<HomepageOverviewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HomepageOverviewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HomepageOverviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
