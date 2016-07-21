import { RouterConfig } from '@angular/router';

import { AddBlueprintComponent } from './components/add-blueprint/add-blueprint';
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
    path: 'blueprints/add',
    component: AddBlueprintComponent,
  },
  {
    path: 'blueprints/:id',
    component: BlueprintsComponent,
  },
  {
    path: '**',
    redirectTo: '',
  },
];
