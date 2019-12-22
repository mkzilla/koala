import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { User } from '../models/user';

@Injectable({
  providedIn: 'root'
})
export class CacheService {
  private user = new BehaviorSubject<User>(new User());
  currentUser = this.user.asObservable();

  constructor() {}

  changeUser(u: User) {
    this.user.next(u);
  }
}
