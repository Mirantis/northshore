import { Injectable } from '@angular/core';

@Injectable()
export class AssetsService {

  private assets = {
    api: {
      blueprintsUrl: 'ui/api/v1/blueprints',
      parseBlueprintUrl: 'ui/api/v1/blueprints/parse',
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
