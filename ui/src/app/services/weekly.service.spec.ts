import { TestBed } from '@angular/core/testing';

import { WeeklyService } from './weekly.service';

describe('WeeklyService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: WeeklyService = TestBed.get(WeeklyService);
    expect(service).toBeTruthy();
  });
});
