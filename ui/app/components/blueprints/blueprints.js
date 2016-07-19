"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
var core_1 = require('@angular/core');
var router_1 = require('@angular/router');
var api_1 = require('../../services/api/api');
var blueprint_details_1 = require('../blueprint-details/blueprint-details');
var BlueprintsComponent = (function () {
    function BlueprintsComponent(apiService, route) {
        this.apiService = apiService;
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
            .map(function (params) { return params['id']; })
            .subscribe(function (id) {
            _this.idSelected = id;
            _this.getSelected();
        });
        this.subscriptions.push(sub);
    };
    BlueprintsComponent.prototype.getBlueprints = function () {
        var _this = this;
        var sub = this.apiService.getBlueprints()
            .subscribe(function (blueprints) {
            _this.blueprints = blueprints;
            _this.getSelected();
        });
        this.subscriptions.push(sub);
    };
    BlueprintsComponent.prototype.getSelected = function () {
        if (this.idSelected) {
            for (var _i = 0, _a = this.blueprints; _i < _a.length; _i++) {
                var bp = _a[_i];
                if (bp.id == this.idSelected) {
                    this.bpSelected = bp;
                }
            }
        }
    };
    BlueprintsComponent = __decorate([
        core_1.Component({
            directives: [
                blueprint_details_1.BlueprintDetailsComponent,
            ],
            providers: [
                api_1.APIService,
            ],
            selector: 'blueprints',
            template: require('./blueprints.html'),
        }), 
        __metadata('design:paramtypes', [api_1.APIService, router_1.ActivatedRoute])
    ], BlueprintsComponent);
    return BlueprintsComponent;
}());
exports.BlueprintsComponent = BlueprintsComponent;
//# sourceMappingURL=blueprints.js.map