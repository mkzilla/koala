import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GiveUpTechniqueComponent } from './give-up-technique.component';

describe('GiveUpTechniqueComponent', () => {
  let component: GiveUpTechniqueComponent;
  let fixture: ComponentFixture<GiveUpTechniqueComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GiveUpTechniqueComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GiveUpTechniqueComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
