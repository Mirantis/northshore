System.register(['@angular/core', '../../pipes/iterate', '../../services/blueprints/blueprints'], function(exports_1, context_1) {
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
    var core_1, iterate_1, blueprints_1;
    var BlueprintDetailsComponent;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            },
            function (iterate_1_1) {
                iterate_1 = iterate_1_1;
            },
            function (blueprints_1_1) {
                blueprints_1 = blueprints_1_1;
            }],
        execute: function() {
            BlueprintDetailsComponent = (function () {
                function BlueprintDetailsComponent() {
                }
                __decorate([
                    core_1.Input(), 
                    __metadata('design:type', blueprints_1.Blueprint)
                ], BlueprintDetailsComponent.prototype, "blueprint", void 0);
                BlueprintDetailsComponent = __decorate([
                    core_1.Component({
                        selector: 'blueprint-details',
                        pipes: [iterate_1.KeysPipe],
                        templateUrl: 'app/components/blueprint-details/blueprint-details.html',
                    }), 
                    __metadata('design:paramtypes', [])
                ], BlueprintDetailsComponent);
                return BlueprintDetailsComponent;
            }());
            exports_1("BlueprintDetailsComponent", BlueprintDetailsComponent);
        }
    }
});
//# sourceMappingURL=blueprint-details.js.map