import { Injectable } from '@angular/core';

@Injectable()
export class AssetsService {

  private assets = {
    api: {
      blueprintsUrl: 'api/v1/blueprints',
      parseBlueprintUrl: 'api/v1/parse/blueprint',
    },
    alerts: {
      AlertsService: {
        error: "There are some internal issues while trying to process your request.",
        notImplemented: "Not implemented yet.",
      },
      uibAlertTypes: {
        default: 'info',
        error: 'danger',
        success: 'success',
        warning: 'warning',
      },
      Blueprints: {
        onParseSuccess: 'The new Blueprint successfully stored',
      },
    },
    timers: {
      alertDismiss: 9000,
      blueprintsInterval: 5000,
      fadeOut: 1000,
    }
  };

  asset(key: string): any {
    return this.assets[key];
  }

}
