import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HomepageWatchComponent } from './homepage-watch.component';

describe('HomepageWatchComponent', () => {
  let component: HomepageWatchComponent;
  let fixture: ComponentFixture<HomepageWatchComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HomepageWatchComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HomepageWatchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
