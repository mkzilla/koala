import {Task} from '../task';

export class Kanban {
  will: Task[] = [];
  doing: Task[] = [];
  done: Task[] = [];
}
