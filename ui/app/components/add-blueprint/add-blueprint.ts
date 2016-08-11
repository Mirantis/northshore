import { Component, OnDestroy, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TAB_DIRECTIVES } from 'ng2-bootstrap/components/tabs';

import { AlertsService } from '../../services/alerts/alerts';
import { AssetsService } from '../../services/assets/assets';
import { Blueprint, BlueprintYAML, APIService } from '../../services/api/api';

@Component({
  directives: [
    TAB_DIRECTIVES,
  ],
  providers: [
    APIService,
  ],
  selector: 'add-blueprint',
  template: require('./add-blueprint.html'),
})

export class AddBlueprintComponent implements OnDestroy, OnInit {

  formData: string = "";
  private subscriptions: any[] = [];
  private alertOnParseSuccess = this.assetsService.asset('alerts').Blueprints.onParseSuccess;

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

  ngOnInit() {
  }

  submitParseBlueprint(v: BlueprintYAML) {
    let sub = this.apiService.parseBlueprint(v)
      .subscribe((res: Blueprint) => {
        // onSuccess
        this.alertsService.alertSuccess(this.alertOnParseSuccess);
        this.router.navigate(['/blueprints', res.id]);
      });
    this.subscriptions.push(sub);
  }

}
