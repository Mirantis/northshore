import { BlueprintsComponent } from './components/blueprints/blueprints';
import { HomeComponent } from './components/home/home';

export const AppRoutes = [
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
