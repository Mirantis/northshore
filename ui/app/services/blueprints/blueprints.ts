import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import 'rxjs/add/operator/toPromise';

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

  private handleError(error: any) {
    console.error('#BlueprintsService,#Error', error);
    this.alertsService.alertError('Some error alert here');
    return Promise.reject(error.message || error);
  }

  getBlueprints(): Promise<Blueprint[]> {
    return this.http.get(this.blueprintsUrl)
      .toPromise()
      .then(response => response.json().data)
      .catch(this.handleError);
  }

}
