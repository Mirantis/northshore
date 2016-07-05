import { Component, Input } from '@angular/core';

import { Blueprint } from '../../services/api/api';

declare var __moduleName: string;

@Component({
  moduleId: __moduleName,
  selector: 'blueprint-details',
  templateUrl: 'blueprint-details.html',
})

export class BlueprintDetailsComponent {
  @Input()
  blueprint: Blueprint;
}
