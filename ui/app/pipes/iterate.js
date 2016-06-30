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
    var IterateMapPipe, KeysPipe, SumIfValuePipeOptions, SumIfValuePipe;
    return {
        setters:[
            function (core_1_1) {
                core_1 = core_1_1;
            }],
        execute: function() {
            // http://stackoverflow.com/a/37479557
            IterateMapPipe = (function () {
                function IterateMapPipe() {
                }
                IterateMapPipe.prototype.transform = function (map, args) {
                    if (args === void 0) { args = null; }
                    if (!map)
                        return null;
                    return Object.keys(map)
                        .map(function (key) { return ({ 'key': key, 'value': map[key] }); });
                };
                IterateMapPipe = __decorate([
                    core_1.Pipe({ name: 'iterateMap' }), 
                    __metadata('design:paramtypes', [])
                ], IterateMapPipe);
                return IterateMapPipe;
            }());
            exports_1("IterateMapPipe", IterateMapPipe);
            KeysPipe = (function () {
                function KeysPipe() {
                }
                KeysPipe.prototype.transform = function (map) {
                    return Object.keys(map);
                };
                KeysPipe = __decorate([
                    core_1.Pipe({ name: 'keys' }), 
                    __metadata('design:paramtypes', [])
                ], KeysPipe);
                return KeysPipe;
            }());
            exports_1("KeysPipe", KeysPipe);
            SumIfValuePipeOptions = (function () {
                function SumIfValuePipeOptions() {
                    this.filter = [];
                    this.hideZero = false;
                }
                return SumIfValuePipeOptions;
            }());
            SumIfValuePipe = (function () {
                function SumIfValuePipe() {
                }
                SumIfValuePipe.prototype.transform = function (map, options) {
                    var sum = 0;
                    // ES7 Array.prototype.includes() missing in TypeScript #2340
                    for (var key in map) {
                        if (options.filter.indexOf(map[key]) > -1) {
                            sum++;
                        }
                    }
                    return (sum === 0 && options.hideZero) ? null : sum;
                };
                SumIfValuePipe = __decorate([
                    core_1.Pipe({ name: 'sumIfValue' }), 
                    __metadata('design:paramtypes', [])
                ], SumIfValuePipe);
                return SumIfValuePipe;
            }());
            exports_1("SumIfValuePipe", SumIfValuePipe);
        }
    }
});
//# sourceMappingURL=iterate.js.map