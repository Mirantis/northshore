import { bootstrap }    from '@angular/platform-browser-dynamic';
import { provide, PLATFORM_DIRECTIVES, PLATFORM_PIPES } from '@angular/core';
import { HTTP_PROVIDERS } from '@angular/http';
import { provideRouter, ROUTER_DIRECTIVES } from '@angular/router';
//import {enableProdMode} from '@angular/core';

import { AppRoutes } from './routes';
import { AppComponent } from './components/app/app';
import { KeysPipe } from './pipes/iterate';

//if (process.env.ENV === 'production') {
//  enableProdMode();
//}

bootstrap(AppComponent, [
  HTTP_PROVIDERS,
  provide(PLATFORM_DIRECTIVES, { useValue: [ROUTER_DIRECTIVES], multi: true }),
  provide(PLATFORM_PIPES, { useValue: [KeysPipe], multi: true }),
  provideRouter(AppRoutes),
]);
