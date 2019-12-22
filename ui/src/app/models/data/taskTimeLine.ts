export class TaskLog {
  id = 0;
  task_id = 0;
  username = '';
  user_id = 0;
  create_time: Date = new Date();
  update_time: Date = new Date();
  to_user_id = 0;
  to_username = '';
}

export class TaskTimeLine {
  task_log: TaskLog = new TaskLog();
  nickname = '';
  title = '';
  state = '';
  importance = 1;
  priority = 1;
  type = '';
}
