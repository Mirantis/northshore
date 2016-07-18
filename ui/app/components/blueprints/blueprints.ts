import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { Blueprint, APIService } from '../../services/api/api';
import { BlueprintDetailsComponent } from '../blueprint-details/blueprint-details';

@Component({
  directives: [
    BlueprintDetailsComponent,
  ],
  providers: [
    APIService,
  ],
  selector: 'blueprints',
  template: require('./blueprints.html'),
})

export class BlueprintsComponent implements OnDestroy, OnInit {

  blueprints: Blueprint[] = [];
  bpSelected: Blueprint;
  private subscriptions: any[] = [];
  private uuidSelected: String;

  constructor(
    private apiService: APIService,
    private route: ActivatedRoute
  ) { }

  ngOnDestroy() {
    for (let sub of this.subscriptions) {
      sub.unsubscribe();
    }
  }

  ngOnInit() {
    this.getBlueprints();
    let sub = this.route.params
      .map(params => params['uuid'])
      .subscribe(uuid => {
        this.uuidSelected = uuid;
        this.getSelected();
      });
    this.subscriptions.push(sub);
  }

  private getBlueprints() {
    let sub = this.apiService.getBlueprints()
      .subscribe(blueprints => {
        this.blueprints = blueprints;
        this.getSelected();
      });
    this.subscriptions.push(sub);
  }

  private getSelected() {
    if (this.uuidSelected) {
      for (let bp of this.blueprints) {
        if (bp.uuid == this.uuidSelected) {
          this.bpSelected = bp;
        }
      }
    }
  }

}
