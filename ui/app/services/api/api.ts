import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';

import 'object-assign';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/observable/fromPromise';
import 'rxjs/add/observable/interval';
import 'rxjs/add/observable/throw';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/share';
import 'rxjs/add/operator/startWith';
import 'rxjs/add/operator/switchMap';

import { AlertsService } from '../alerts/alerts';
import { AssetsService } from '../assets/assets';

export class Blueprint {
  name: string;
  provisioner: string; //Provisioner type (docker/...)
  stages: Object[];
  state: string;
  type: string; //Type of blueprint (pipeline/application)
  version: string;
  ui: {
    stagesStatesBages: {};
  };
  id: string;
}

export class BlueprintYAML {
  data: string;
}

@Injectable()
export class APIService {

  private blueprints: Observable<Blueprint[]>;
  private blueprintsInterval = this.assetsService.asset('timers').blueprintsInterval;
  private blueprintsUrl = this.assetsService.asset('api').blueprintsUrl;
  private parseBlueprintUrl = this.assetsService.asset('api').parseBlueprintUrl;

  constructor(
    private alertsService: AlertsService,
    private assetsService: AssetsService,
    private http: Http
  ) {

    let JSONAPIDeserializer = require('jsonapi-serializer').Deserializer;
    let p = new JSONAPIDeserializer({ keyForAttribute: 'camelCase' });

    this.blueprints = Observable
      .interval(this.blueprintsInterval)
      .startWith(0)
      .switchMap(() => this.http.get(this.blueprintsUrl))
      .switchMap(res => Observable.fromPromise(p.deserialize(res.json())))
      .map(this.extendBlueprintsData)
      .share()
      .catch(error => this.handleError(error, '#APIService.getBlueprints,#Error'));

  }

  private extendBlueprintsData(bps: {}) {
    let stagesStatesBages = {};
    let filters = {
      green: ['running'],
      orange: ['new', 'created'],
      grey: ['deleted', 'paused', 'stopped'],
    };
    for (let f in filters) {
      stagesStatesBages[f] = 0;
    }

    for (let i in bps) {
      let bp = bps[i]
      bp.ui = {
        stagesStatesBages: Object.assign({}, stagesStatesBages)
      }

      for (let s in bp.stages) {
        for (let f in filters) {
          if (filters[f].indexOf(bp.stages[s].state) > -1) {
            bp.ui.stagesStatesBages[f]++;
            break;
          }
        }
      }
    }
    return bps;
  }

  private handleError(error: any, logTags?: string) {
    console.error(logTags ? logTags : '#APIService,#Error', error);
    // handle JSONAPI Errors
    try {
      let o = error.json()
      if (o && o.errors) {
        for (let i in o.errors) {
          this.alertsService.alertError(
            o.errors[i].title + ' ' + o.errors[i].detail
          );
        }
      }
    } catch (e) {
      this.alertsService.alertError();
    }

    return Observable.throw(error);
  }

  /**
    @description Returns the Observable that repeats the XHR while subscribed.
   */
  getBlueprints(): Observable<Blueprint[]> {
    return this.blueprints;
  }

  parseBlueprint(v: BlueprintYAML) {
    let headers = new Headers({
      'Content-Type': 'application/vnd.api+json'
    });

    let JSONAPISerializer = require('jsonapi-serializer').Serializer;
    let s = new JSONAPISerializer('blueprintYAML', { attributes: ['data'], pluralizeType: false });
    let payload = s.serialize(v)
    // https://github.com/SeyZ/jsonapi-serializer/issues/70#issuecomment-197310834
    delete (payload.data.id);

    return this.http
      .post(this.parseBlueprintUrl, payload, { headers: headers })
      .toPromise()
      .catch(error => this.handleError(error, '#APIService.parseBlueprint,#Error'));
  }

}
