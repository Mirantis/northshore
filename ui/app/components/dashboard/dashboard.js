System.register(['@angular/core', '../../services/blueprints/blueprints'], function(exports_1, context_1) {
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
    var core_1, blueprints_1;
    var DashboardComponent;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            },
            function (blueprints_1_1) {
                blueprints_1 = blueprints_1_1;
            }],
        execute: function() {
            DashboardComponent = (function () {
                function DashboardComponent(blueprintsService) {
                    this.blueprintsService = blueprintsService;
                }
                DashboardComponent.prototype.ngOnInit = function () {
                    this.getBlueprints();
                };
                DashboardComponent.prototype.getBlueprints = function () {
                    var _this = this;
                    this.blueprintsService
                        .getBlueprints()
                        .then(function (blueprints) { return _this.blueprints = blueprints; })
                        .catch(function (error) { return _this.error = error; }); // TODO: Display error message
                };
                DashboardComponent = __decorate([
                    core_1.Component({
                        selector: 'my-dashboard',
                        providers: [blueprints_1.BlueprintsService],
                        styleUrls: ['app/components/dashboard/dashboard.css'],
                        templateUrl: 'app/components/dashboard/dashboard.html',
                    }), 
                    __metadata('design:paramtypes', [blueprints_1.BlueprintsService])
                ], DashboardComponent);
                return DashboardComponent;
            }());
            exports_1("DashboardComponent", DashboardComponent);
        }
    }
});
//# sourceMappingURL=dashboard.js.map