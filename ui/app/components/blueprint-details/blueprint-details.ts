import { Component, Input, OnDestroy } from '@angular/core';
import { Router } from '@angular/router';

import { AlertsService } from '../../services/alerts/alerts';
import { AssetsService } from '../../services/assets/assets';
import { Blueprint, APIService } from '../../services/api/api';

@Component({
  selector: 'blueprint-details',
  template: require('./blueprint-details.html'),
})

export class BlueprintDetailsComponent implements OnDestroy {
  @Input()
  blueprint: Blueprint;

  private alertOnDeleteSuccess = this.assetsService.asset('alerts').Blueprints.onDeleteSuccess;
  private alertOnRunSuccess = this.assetsService.asset('alerts').Blueprints.onRunSuccess;
  private subscriptions: any[] = [];

  constructor(
    private alertsService: AlertsService,
    private assetsService: AssetsService,
    private apiService: APIService,
    private router: Router
  ) { }

  ngOnDestroy() {
    for (let sub of this.subscriptions) {
      sub.unsubscribe();
    }
  }

  deleteBlueprint() {
    // TODO: modal
    console.log('#deleteBlueprint');

    let sub = this.apiService.deleteBlueprint(this.blueprint)
      .subscribe(() => {
        // onSuccess
        this.alertsService.alertSuccess(this.alertOnDeleteSuccess);
        this.router.navigate(['/blueprints']);
      });
    this.subscriptions.push(sub);
  }

  runBlueprint() {
    if (this.blueprint.state != 'new') {
      return;
    }
    this.blueprint.state = 'active';

    let sub = this.apiService.updateBlueprint(this.blueprint)
      .subscribe(() => {
        // onSuccess
        this.alertsService.alertSuccess(this.alertOnRunSuccess);
      });
    this.subscriptions.push(sub);
  }
}
