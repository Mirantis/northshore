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
    var IterateMapPipe, KeysPipe;
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
            // http://stackoverflow.com/a/35536052
            KeysPipe = (function () {
                function KeysPipe() {
                }
                KeysPipe.prototype.transform = function (map, args) {
                    if (args === void 0) { args = null; }
                    if (!map)
                        return null;
                    var keys = [];
                    for (var key in map) {
                        keys.push(key);
                    }
                    return keys;
                };
                KeysPipe = __decorate([
                    core_1.Pipe({ name: 'keys' }), 
                    __metadata('design:paramtypes', [])
                ], KeysPipe);
                return KeysPipe;
            }());
            exports_1("KeysPipe", KeysPipe);
        }
    }
});
//# sourceMappingURL=iterate.js.map