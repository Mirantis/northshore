import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { Blueprint, BlueprintsService } from '../../services/blueprints/blueprints';
import { BlueprintDetailsComponent } from '../blueprint-details/blueprint-details';

@Component({
  selector: 'my-dashboard',
  directives: [
    BlueprintDetailsComponent,
  ],
  providers: [
    BlueprintsService,
  ],
  templateUrl: 'app/components/blueprints/blueprints.html',
})

export class BlueprintsComponent implements OnInit {

  blueprints: Blueprint[] = [];
  bpSelected: Blueprint;
  private bpSelectedName: String;

  constructor(
    private blueprintsService: BlueprintsService,
    private route: ActivatedRoute
  ) { }

  ngOnInit() {
    this.getBlueprints();
    this.route.params
      .map(params => params['name'])
      .subscribe(name => {
        this.bpSelectedName = name;
        this.getSelected();
      });
  }

  private getBlueprints() {
    this.blueprintsService.getBlueprints()
      .subscribe(blueprints => {
        this.blueprints = blueprints;
        this.getSelected();
      });
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
