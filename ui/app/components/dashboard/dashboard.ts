import { Component, OnInit } from '@angular/core';

import { Blueprint, BlueprintsService } from '../../services/blueprints/blueprints';

@Component({
  selector: 'my-dashboard',
  providers: [BlueprintsService],
  styleUrls: ['app/components/dashboard/dashboard.css'],
  templateUrl: 'app/components/dashboard/dashboard.html',
})

export class DashboardComponent implements OnInit {

  blueprints: Blueprint[];
  error: any;

  constructor(
    private blueprintsService: BlueprintsService
  ) { }

  ngOnInit() {
    this.getBlueprints();
  }

  getBlueprints() {
    this.blueprintsService
      .getBlueprints()
      .then(blueprints => this.blueprints = blueprints)
      .catch(error => this.error = error); // TODO: Display error message
  }
 
}
