import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TAB_DIRECTIVES } from 'ng2-bootstrap/components/tabs';

import { APIFormParseBlueprint, APIService } from '../../services/api/api';

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

  formData: string;
  private subscriptions: any[] = [];

  constructor(
    private apiService: APIService
  ) { }

  ngOnDestroy() {
    for (let sub of this.subscriptions) {
      sub.unsubscribe();
    }
  }

  ngOnInit() {
  }

  submitParseBlueprint(v: APIFormParseBlueprint) {
    this.apiService.parseBlueprint(v);
  }

}
