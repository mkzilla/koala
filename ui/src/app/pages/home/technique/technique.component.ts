import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {AuthService} from '../../../services/auth.service';
import { NzNotificationService} from 'ng-zorro-antd';
import {Router} from '@angular/router';

@Component({
  selector: 'app-technique',
  templateUrl: './technique.component.html',
  styleUrls: ['./technique.component.scss']
})
export class TechniqueComponent implements OnInit {
  @Output() refresh = new EventEmitter<boolean>();
  isVisible = false;

  constructor(private authService: AuthService, private notification: NzNotificationService, private route: Router) {}

  showModal(): void {
    this.isVisible = true;
  }

  handleOk(): void {
    this.authService.doDoneTechnique().then((response) => {
      if (response !== undefined) {
        this.refresh.emit(this.isVisible);
        this.route.navigate([`/dashboard`]);
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


