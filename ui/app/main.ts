import { bootstrap }    from '@angular/platform-browser-dynamic';
import { HTTP_PROVIDERS } from '@angular/http';

import { AppComponent } from './components/app/app';

bootstrap(AppComponent, [ HTTP_PROVIDERS ]);
