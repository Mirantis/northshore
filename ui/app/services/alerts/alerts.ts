import { Injectable } from '@angular/core';
import { AssetsService } from '../assets/assets';

@Injectable()
export class AlertsService {

  private alerts: Array<Object> = [];
  private assetAlerts = this.assetsService.asset('alerts');

  constructor(private assetsService: AssetsService) {
  }

  alert(message?: string, type?: string): void {
    this.alerts.push({
      message: message
        ? message
        : this.assetAlerts.AlertsService.notImplemented,
      type: (type && this.assetAlerts.uibAlertTypes[type])
        ? this.assetAlerts.uibAlertTypes[type]
        : this.assetAlerts.uibAlertTypes.default,
    });
  }

  alertError(message?: string): void {
    this.alert(
      message ? message : this.assetAlerts.AlertsService.error,
      'error'
    );
  }

  alertSuccess(message: string): void {
    this.alert(message, 'success');
  }

  clearAlerts(): void {
    this.alerts.length = 0;
  }

  deleteAlert(idx: number): void {
    this.alerts.splice(idx, 1);
  }

  getAlerts(): Array<Object> {
    return this.alerts;
  }

}
