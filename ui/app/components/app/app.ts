import { Component } from '@angular/core';

import { AlertsService } from '../../services/alerts/alerts';
import { AssetsService } from '../../services/assets/assets';

import '../../../assets/custom.css';

@Component({
  selector: 'my-app',
  template: require('./app.html'),
})

export class AppComponent {
  title = 'NorthShore: A Pipeline Generator';
  alertDismiss = this.assetsService.asset('timers').alertDismiss;

  constructor(
    private alertsService: AlertsService,
    private assetsService: AssetsService
  ) { }

  deleteAlert(idx: number): void {
    this.alertsService.deleteAlert(idx);
  }

  getAlerts(): Array<Object> {
    return this.alertsService.getAlerts();
  }

}
