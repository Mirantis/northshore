import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { SumIfValuePipe } from '../../pipes/iterate';

import { Blueprint, APIService } from '../../services/api/api';
import { BlueprintDetailsComponent } from '../blueprint-details/blueprint-details';

declare var __moduleName: string;

@Component({
  directives: [
    BlueprintDetailsComponent,
  ],
  moduleId: __moduleName,
  pipes: [
    SumIfValuePipe,
  ],
  providers: [
    APIService,
  ],
  selector: 'blueprints',
  templateUrl: 'blueprints.html',
})

export class BlueprintsComponent implements OnDestroy, OnInit {

  blueprints: Blueprint[] = [];
  bpSelected: Blueprint;
  private bpSelectedName: String;
  private subscriptions: any[] = [];

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
      .map(params => params['name'])
      .subscribe(name => {
        this.bpSelectedName = name;
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
    if (this.bpSelectedName) {
      for (let bp of this.blueprints) {
        if (bp.name == this.bpSelectedName) {
          this.bpSelected = bp;
        }
      }
    }
  }

}
