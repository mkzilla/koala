import { Task } from '../task';
import { Comment } from '../comment';

export class Report {
  task: Task = new Task();
  schedules: Comment[] = [];
}
