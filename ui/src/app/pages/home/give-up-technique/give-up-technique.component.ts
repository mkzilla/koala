import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {AuthService} from '../../../services/auth.service';
import { NzNotificationService} from 'ng-zorro-antd';

@Component({
  selector: 'app-give-up-technique',
  templateUrl: './give-up-technique.component.html',
  styleUrls: ['./give-up-technique.component.scss']
})
export class GiveUpTechniqueComponent implements OnInit {
  @Output() refresh = new EventEmitter<boolean>();
  isVisible = false;

  constructor(private authService: AuthService, private notification: NzNotificationService) {}

  showModal(): void {
    this.isVisible = true;
  }

  handleOk(): void {
    this.authService.doDoneTechnique().then((response) => {
      if (response !== undefined) {
        this.notification.success('成功取消番茄时间', '');
        this.refresh.emit(this.isVisible);
      }
    });
    this.isVisible = false;
  }

  handleCancel(): void {
    this.isVisible = false;
  }

  ngOnInit() {
  }

}

