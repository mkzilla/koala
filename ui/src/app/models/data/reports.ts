import {Report} from './report';
import {User} from '../user';

export class Reports {
  from: Date = new Date();
  to: Date = new Date();
  reports: Report[] = [];
  user: User = new User();
}
