import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AlertModule, CollapseModule, TabsModule } from 'ng2-bootstrap/ng2-bootstrap';

import { routing } from './routing';

import { AppComponent } from './components/app/app';
import { AddBlueprintComponent } from './components/add-blueprint/add-blueprint';
import { BlueprintsComponent } from './components/blueprints/blueprints';
import { BlueprintDetailsComponent } from './components/blueprint-details/blueprint-details';
import { HomeComponent } from './components/home/home';

import { KeysPipe } from './pipes/iterate';

import { AlertsService } from './services/alerts/alerts';
import { APIService } from './services/api/api';
import { AssetsService } from './services/assets/assets';

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    routing,
    // ng2-bootstrap
    AlertModule,
    CollapseModule,
    TabsModule,
  ],

  declarations: [
    // components
    AppComponent,
    AddBlueprintComponent,
    BlueprintsComponent,
    BlueprintDetailsComponent,
    HomeComponent,
    // pipes
    KeysPipe,
  ],

  providers: [
    AlertsService,
    APIService,
    AssetsService,
  ],

  bootstrap: [AppComponent],
})

export class AppModule { }
