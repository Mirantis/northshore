import { Routes, RouterModule } from '@angular/router';

import { AddBlueprintComponent } from './components/add-blueprint/add-blueprint';
import { BlueprintsComponent } from './components/blueprints/blueprints';
import { HomeComponent } from './components/home/home';

const appRoutes: Routes = [
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

export const routing = RouterModule.forRoot(appRoutes);
