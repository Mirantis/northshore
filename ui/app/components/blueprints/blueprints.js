System.register(['@angular/core', '@angular/router', '../../services/blueprints/blueprints', '../blueprint-details/blueprint-details'], function(exports_1, context_1) {
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
    var core_1, router_1, blueprints_1, blueprint_details_1;
    var BlueprintsComponent;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            },
            function (router_1_1) {
                router_1 = router_1_1;
            },
            function (blueprints_1_1) {
                blueprints_1 = blueprints_1_1;
            },
            function (blueprint_details_1_1) {
                blueprint_details_1 = blueprint_details_1_1;
            }],
        execute: function() {
            BlueprintsComponent = (function () {
                function BlueprintsComponent(blueprintsService, route) {
                    this.blueprintsService = blueprintsService;
                    this.route = route;
                    this.blueprints = [];
                    this.subscriptions = [];
                }
                BlueprintsComponent.prototype.ngOnDestroy = function () {
                    for (var _i = 0, _a = this.subscriptions; _i < _a.length; _i++) {
                        var sub = _a[_i];
                        sub.unsubscribe();
                    }
                };
                BlueprintsComponent.prototype.ngOnInit = function () {
                    var _this = this;
                    this.getBlueprints();
                    var sub = this.route.params
                        .map(function (params) { return params['name']; })
                        .subscribe(function (name) {
                        _this.bpSelectedName = name;
                        _this.getSelected();
                    });
                    this.subscriptions.push(sub);
                };
                BlueprintsComponent.prototype.getBlueprints = function () {
                    var _this = this;
                    var sub = this.blueprintsService.getBlueprints()
                        .subscribe(function (blueprints) {
                        _this.blueprints = blueprints;
                        _this.getSelected();
                    });
                    this.subscriptions.push(sub);
                };
                BlueprintsComponent.prototype.getSelected = function () {
                    if (this.bpSelectedName) {
                        for (var _i = 0, _a = this.blueprints; _i < _a.length; _i++) {
                            var bp = _a[_i];
                            if (bp.name == this.bpSelectedName) {
                                this.bpSelected = bp;
                            }
                        }
                    }
                };
                BlueprintsComponent = __decorate([
                    core_1.Component({
                        selector: 'my-dashboard',
                        directives: [
                            blueprint_details_1.BlueprintDetailsComponent,
                        ],
                        providers: [
                            blueprints_1.BlueprintsService,
                        ],
                        templateUrl: 'app/components/blueprints/blueprints.html',
                    }), 
                    __metadata('design:paramtypes', [blueprints_1.BlueprintsService, router_1.ActivatedRoute])
                ], BlueprintsComponent);
                return BlueprintsComponent;
            }());
            exports_1("BlueprintsComponent", BlueprintsComponent);
        }
    }
});
//# sourceMappingURL=blueprints.js.map