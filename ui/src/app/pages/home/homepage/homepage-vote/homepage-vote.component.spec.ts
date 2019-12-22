import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HomepageVoteComponent } from './homepage-vote.component';

describe('HomepageVoteComponent', () => {
  let component: HomepageVoteComponent;
  let fixture: ComponentFixture<HomepageVoteComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HomepageVoteComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HomepageVoteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
