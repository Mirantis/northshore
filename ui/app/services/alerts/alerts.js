System.register(['@angular/core', '../assets/assets'], function(exports_1, context_1) {
    "use strict";
    var __moduleName = context_1 && context_1.id;
    var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
        var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
        if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
        else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
        return c > 3 && r && Object.defineProperty(target, key, r), r;
    };
    var __metadata = (this && this.__metadata) || function (k, v) {
        if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
    };
    var core_1, assets_1;
    var AlertsService;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            },
            function (assets_1_1) {
                assets_1 = assets_1_1;
            }],
        execute: function() {
            AlertsService = (function () {
                function AlertsService(assetsService) {
                    this.assetsService = assetsService;
                    this.alerts = [];
                    this.assetAlerts = this.assetsService.asset('alerts');
                }
                AlertsService.prototype.alert = function (message, type) {
                    this.alerts.push({
                        message: message
                            ? message
                            : this.assetAlerts.AlertsService.notImplemented,
                        type: (type && this.assetAlerts.uibAlertTypes[type])
                            ? this.assetAlerts.uibAlertTypes[type]
                            : this.assetAlerts.uibAlertTypes.default,
                    });
                };
                AlertsService.prototype.alertError = function (message) {
                    this.alert(message ? message : this.assetAlerts.AlertsService.error, 'error');
                };
                AlertsService.prototype.alertSuccess = function (message) {
                    this.alert(message, 'success');
                };
                AlertsService.prototype.clearAlerts = function () {
                    this.alerts.length = 0;
                };
                AlertsService.prototype.deleteAlert = function (idx) {
                    this.alerts.splice(idx, 1);
                };
                AlertsService.prototype.getAlerts = function () {
                    return this.alerts;
                };
                AlertsService = __decorate([
                    core_1.Injectable(), 
                    __metadata('design:paramtypes', [assets_1.AssetsService])
                ], AlertsService);
                return AlertsService;
            }());
            exports_1("AlertsService", AlertsService);
        }
    }
});
//# sourceMappingURL=alerts.js.map