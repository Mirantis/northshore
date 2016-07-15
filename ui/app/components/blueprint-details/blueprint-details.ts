import { Component, Input } from '@angular/core';

import { Blueprint } from '../../services/api/api';

@Component({
  selector: 'blueprint-details',
  template: require('./blueprint-details.html'),
})

export class BlueprintDetailsComponent {
  @Input()
  blueprint: Blueprint;
}
