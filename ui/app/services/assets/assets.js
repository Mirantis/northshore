System.register(['@angular/core'], function(exports_1, context_1) {
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
    var core_1;
    var AssetsService;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            }],
        execute: function() {
            AssetsService = (function () {
                function AssetsService() {
                    this.assets = {
                        api: {
                            blueprintsUrl: 'api/v1/blueprints',
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
                }
                AssetsService.prototype.asset = function (key) {
                    return this.assets[key];
                };
                AssetsService = __decorate([
                    core_1.Injectable(), 
                    __metadata('design:paramtypes', [])
                ], AssetsService);
                return AssetsService;
            }());
            exports_1("AssetsService", AssetsService);
        }
    }
});
//# sourceMappingURL=assets.js.map