import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { Blueprint, APIService } from '../../services/api/api';
import { BlueprintDetailsComponent } from '../blueprint-details/blueprint-details';

@Component({
  selector: 'my-dashboard',
  directives: [
    BlueprintDetailsComponent,
  ],
  providers: [
    APIService,
  ],
  templateUrl: 'app/components/blueprints/blueprints.html',
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
