import { Component, Input } from '@angular/core';

import { KeysPipe } from '../../pipes/iterate';
import { Blueprint } from '../../services/blueprints/blueprints';

@Component({
  selector: 'blueprint-details',
  pipes: [KeysPipe],
  templateUrl: 'app/components/blueprint-details/blueprint-details.html',
})

export class BlueprintDetailsComponent {
  @Input()
  blueprint: Blueprint;
}
