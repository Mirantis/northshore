import { Injectable } from '@angular/core';
import { Headers, Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/observable/throw';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';

import { AlertsService } from '../alerts/alerts';
import { AssetsService } from '../assets/assets';

export class Blueprint {
  name: string;
  provisioner: string; //Provisioner type (docker/...)
  stages: Object[];
  state: string;
  stateStages: Object[];
  type: string; //Type of blueprint (pipeline/application)
  version: string;
}

@Injectable()
export class BlueprintsService {

  private blueprintsUrl = this.assetsService.asset('api').blueprints;

  constructor(
    private alertsService: AlertsService,
    private assetsService: AssetsService,
    private http: Http
  ) { }

  private extractData(res: Response) {
    let body = res.json();
    return body.data || {};
  }

  private handleError(error: any) {
    console.error('#BlueprintsService,#Error', error);
    // TODO: Solve an issue
    // platform-browser.umd.js:962 EXCEPTION: TypeError: Cannot read property 'alertError' of undefined
    // this.alertsService.alertError('Some error alert here');

    return Observable.throw(error.message || error);
  }

  getBlueprints(): Observable<Blueprint[]> {
    return this.http.get(this.blueprintsUrl)
      .map(this.extractData)
      // TODO: move to some generic handler
      // .catch(this.handleError);
      .catch((error: any) => {
        console.error('#BlueprintsService,#Error', error);
        try {
          // handle JSONAPI Errors
          let o = error.json()
          if (o && o.errors) {
            for (let i in o.errors) {
              this.alertsService.alertError(o.errors[i].details);
            }
          }
        } catch (e) {
          this.alertsService.alertError();
        }
        return Observable.throw(error);
      });
  }

}
