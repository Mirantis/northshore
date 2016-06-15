import { bootstrap }    from '@angular/platform-browser-dynamic';
import { HTTP_PROVIDERS } from '@angular/http';
import { provideRouter } from '@angular/router';

import { AppRoutes } from './routes';
import { AppComponent } from './components/app/app';

bootstrap(AppComponent, [
  HTTP_PROVIDERS,
  provideRouter(AppRoutes),
]);
