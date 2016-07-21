import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { Blueprint, APIService } from '../../services/api/api';

@Component({
  directives: [],
  providers: [
    APIService,
  ],
  selector: 'add-blueprint',
  template: require('./add-blueprint.html'),
})

export class AddBlueprintComponent implements OnDestroy, OnInit {

  private subscriptions: any[] = [];

  constructor(
  ) { }

  ngOnDestroy() {
    for (let sub of this.subscriptions) {
      sub.unsubscribe();
    }
  }

  ngOnInit() {
  }

}
