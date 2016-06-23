import { RouterConfig } from '@angular/router';

import { BlueprintsComponent } from './components/blueprints/blueprints';
import { HomeComponent } from './components/home/home';

export const AppRoutes: RouterConfig = [
  {
    path: '',
    component: HomeComponent,
  },
  {
    path: 'blueprints',
    component: BlueprintsComponent,
  },
  {
    path: 'blueprints/:name',
    component: BlueprintsComponent,
  },
];
