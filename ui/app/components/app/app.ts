import { Component } from '@angular/core';
import { Router, RouteConfig, ROUTER_DIRECTIVES, ROUTER_PROVIDERS } from '@angular/router-deprecated';
import { CollapseDirective } from 'ng2-bootstrap/components/collapse';

import { DashboardComponent } from '../dashboard/dashboard';
import { HomeComponent } from '../home/home';

@Component({
  selector: 'my-app',
  templateUrl: 'app/components/app/app.html',
  styleUrls: ['app/components/app/app.css'],
  directives: [CollapseDirective, ROUTER_DIRECTIVES],
  providers: [
    ROUTER_PROVIDERS,
  ]
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

  constructor(
    private router: Router) {
  }
}
