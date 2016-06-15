System.register(['@angular/core', '@angular/http', 'rxjs/Observable', 'rxjs/add/observable/throw', 'rxjs/add/operator/catch', 'rxjs/add/operator/map', '../alerts/alerts', '../assets/assets'], function(exports_1, context_1) {
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
    var core_1, http_1, Observable_1, alerts_1, assets_1;
    var Blueprint, BlueprintsService;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            },
            function (http_1_1) {
                http_1 = http_1_1;
            },
            function (Observable_1_1) {
                Observable_1 = Observable_1_1;
            },
            function (_1) {},
            function (_2) {},
            function (_3) {},
            function (alerts_1_1) {
                alerts_1 = alerts_1_1;
            },
            function (assets_1_1) {
                assets_1 = assets_1_1;
            }],
        execute: function() {
            Blueprint = (function () {
                function Blueprint() {
                }
                return Blueprint;
            }());
            exports_1("Blueprint", Blueprint);
            BlueprintsService = (function () {
                function BlueprintsService(alertsService, assetsService, http) {
                    this.alertsService = alertsService;
                    this.assetsService = assetsService;
                    this.http = http;
                    this.blueprintsUrl = this.assetsService.asset('api').blueprints;
                }
                BlueprintsService.prototype.extractData = function (res) {
                    var body = res.json();
                    return body.data || {};
                };
                BlueprintsService.prototype.handleError = function (error) {
                    console.error('#BlueprintsService,#Error', error);
                    // TODO: Solve an issue
                    // platform-browser.umd.js:962 EXCEPTION: TypeError: Cannot read property 'alertError' of undefined
                    // this.alertsService.alertError('Some error alert here');
                    return Observable_1.Observable.throw(error.message || error);
                };
                BlueprintsService.prototype.getBlueprints = function () {
                    var _this = this;
                    return this.http.get(this.blueprintsUrl)
                        .map(this.extractData)
                        .catch(function (error) {
                        console.error('#BlueprintsService,#Error', error);
                        try {
                            // handle JSONAPI Errors
                            var o = error.json();
                            if (o && o.errors) {
                                for (var i in o.errors) {
                                    _this.alertsService.alertError(o.errors[i].details);
                                }
                            }
                        }
                        catch (e) {
                            _this.alertsService.alertError();
                        }
                        return Observable_1.Observable.throw(error);
                    });
                };
                BlueprintsService = __decorate([
                    core_1.Injectable(), 
                    __metadata('design:paramtypes', [alerts_1.AlertsService, assets_1.AssetsService, http_1.Http])
                ], BlueprintsService);
                return BlueprintsService;
            }());
            exports_1("BlueprintsService", BlueprintsService);
        }
    }
});
//# sourceMappingURL=blueprints.js.map