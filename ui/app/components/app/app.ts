import { Component } from '@angular/core';
import { CORE_DIRECTIVES } from '@angular/common';
import { Router, RouteConfig, ROUTER_DIRECTIVES, ROUTER_PROVIDERS } from '@angular/router-deprecated';
import { AlertComponent } from 'ng2-bootstrap/components/alert';
import { CollapseDirective } from 'ng2-bootstrap/components/collapse';

import { AlertsService } from '../../services/alerts/alerts';
import { AssetsService } from '../../services/assets/assets';
import { DashboardComponent } from '../dashboard/dashboard';
import { HomeComponent } from '../home/home';

@Component({
  directives: [
    AlertComponent,
    CollapseDirective,
    CORE_DIRECTIVES,
    ROUTER_DIRECTIVES,
  ],
  providers: [
    AlertsService,
    AssetsService,
    ROUTER_PROVIDERS,
  ],
  selector: 'my-app',
  templateUrl: 'app/components/app/app.html',
})

@RouteConfig([
  {
    path: '/',
    name: 'Home',
    component: HomeComponent,
    useAsDefault: true
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: DashboardComponent,
  },
])

export class AppComponent {
  title = 'NorthShore: A Pipeline Generator';
  alertDismiss = this.assetsService.asset('timers').alertDismiss;

  constructor(
    private alertsService: AlertsService,
    private assetsService: AssetsService,
    private router: Router
  ) { }

  deleteAlert(idx: number): void {
    this.alertsService.deleteAlert(idx);
  }

  getAlerts(): Array<Object> {
    return this.alertsService.getAlerts();
  }

  checkAlerts(): void {
    this.alertsService.alert('Some alert here');
    this.alertsService.alertError('Some error alert here');
    this.alertsService.alertSuccess('Some success alert here');
  }

}
