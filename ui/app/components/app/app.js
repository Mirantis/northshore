System.register(['@angular/core', '@angular/common', 'ng2-bootstrap/components/alert', 'ng2-bootstrap/components/collapse', '../../services/alerts/alerts', '../../services/assets/assets', '../blueprints/blueprints', '../home/home'], function(exports_1, context_1) {
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
    var core_1, common_1, alert_1, collapse_1, alerts_1, assets_1, blueprints_1, home_1;
    var AppComponent;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            },
            function (common_1_1) {
                common_1 = common_1_1;
            },
            function (alert_1_1) {
                alert_1 = alert_1_1;
            },
            function (collapse_1_1) {
                collapse_1 = collapse_1_1;
            },
            function (alerts_1_1) {
                alerts_1 = alerts_1_1;
            },
            function (assets_1_1) {
                assets_1 = assets_1_1;
            },
            function (blueprints_1_1) {
                blueprints_1 = blueprints_1_1;
            },
            function (home_1_1) {
                home_1 = home_1_1;
            }],
        execute: function() {
            AppComponent = (function () {
                function AppComponent(alertsService, assetsService) {
                    this.alertsService = alertsService;
                    this.assetsService = assetsService;
                    this.title = 'NorthShore: A Pipeline Generator';
                    this.alertDismiss = this.assetsService.asset('timers').alertDismiss;
                }
                AppComponent.prototype.deleteAlert = function (idx) {
                    this.alertsService.deleteAlert(idx);
                };
                AppComponent.prototype.getAlerts = function () {
                    return this.alertsService.getAlerts();
                };
                AppComponent = __decorate([
                    core_1.Component({
                        directives: [
                            alert_1.AlertComponent,
                            collapse_1.CollapseDirective,
                            common_1.CORE_DIRECTIVES,
                        ],
                        precompile: [
                            blueprints_1.BlueprintsComponent,
                            home_1.HomeComponent,
                        ],
                        providers: [
                            alerts_1.AlertsService,
                            assets_1.AssetsService,
                        ],
                        selector: 'my-app',
                        templateUrl: 'app/components/app/app.html',
                    }), 
                    __metadata('design:paramtypes', [alerts_1.AlertsService, assets_1.AssetsService])
                ], AppComponent);
                return AppComponent;
            }());
            exports_1("AppComponent", AppComponent);
        }
    }
});
//# sourceMappingURL=app.js.map