import { Injectable } from '@angular/core';

@Injectable()
export class AssetsService {

  _assets = {
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
      fadeOut: 1000,
    }
  };

  asset(key: string): any {
    return this._assets[key];
  }

}
