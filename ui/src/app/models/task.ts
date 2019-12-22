import {User} from './user';

export class Task {
  id = 0;
  create_time: Date = new Date();
  update_time: Date = new Date();
  title = '';
  desc = '';
  state = 0;
  importance = 1;
  priority = 1;
  type = '';
  assign_to = 0;
  is_vote = false;
  is_watch = false;
  user: User;
}
