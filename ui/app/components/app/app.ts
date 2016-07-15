import { Component } from '@angular/core';
import { CORE_DIRECTIVES } from '@angular/common';
import { AlertComponent } from 'ng2-bootstrap/components/alert';
import { CollapseDirective } from 'ng2-bootstrap/components/collapse';

import { AlertsService } from '../../services/alerts/alerts';
import { AssetsService } from '../../services/assets/assets';
import { BlueprintsComponent } from '../blueprints/blueprints';
import { HomeComponent } from '../home/home';

import '../../../assets/custom.css';

@Component({
  directives: [
    AlertComponent,
    CollapseDirective,
    CORE_DIRECTIVES,
  ],
  precompile: [
    BlueprintsComponent,
    HomeComponent,
  ],
  providers: [
    AlertsService,
    AssetsService,
  ],
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
