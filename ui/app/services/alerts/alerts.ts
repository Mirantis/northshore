import { Injectable } from '@angular/core';
import { AssetsService } from '../assets/assets';

@Injectable()
export class AlertsService {

  _alerts: Array<Object>;
  _assetAlerts = this.assetsService.asset('alerts');

  constructor(private assetsService: AssetsService) {
    this._alerts = [];
  }

  alert(message: string, type?: string): void {
    this._alerts.push({
      message: message,
      type: (type && this._assetAlerts.uibAlertTypes[type])
        ? this._assetAlerts.uibAlertTypes[type]
        : this._assetAlerts.uibAlertTypes.default,
    });
  }

  alertError(message: string): void {
    this.alert(
      message ? message : this._assetAlerts.AlertsService.error,
      'error'
    );
  }

  alertSuccess(message: string): void {
    this.alert(message, 'success');
  }

  clearAlerts(): void {
    this._alerts.length = 0;
  }

  deleteAlert(idx: number): void {
    this._alerts.splice(idx, 1);
  }

  getAlerts(): Array<Object> {
    return this._alerts;
  }

}
