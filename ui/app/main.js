System.register(['@angular/platform-browser-dynamic', '@angular/core', '@angular/http', '@angular/router', './routes', './components/app/app', './pipes/iterate'], function(exports_1, context_1) {
    "use strict";
    var __moduleName = context_1 && context_1.id;
    var platform_browser_dynamic_1, core_1, http_1, router_1, routes_1, app_1, iterate_1;
    return {
        setters:[
            function (platform_browser_dynamic_1_1) {
                platform_browser_dynamic_1 = platform_browser_dynamic_1_1;
            },
            function (core_1_1) {
                core_1 = core_1_1;
            },
            function (http_1_1) {
                http_1 = http_1_1;
            },
            function (router_1_1) {
                router_1 = router_1_1;
            },
            function (routes_1_1) {
                routes_1 = routes_1_1;
            },
            function (app_1_1) {
                app_1 = app_1_1;
            },
            function (iterate_1_1) {
                iterate_1 = iterate_1_1;
            }],
        execute: function() {
            platform_browser_dynamic_1.bootstrap(app_1.AppComponent, [
                http_1.HTTP_PROVIDERS,
                core_1.provide(core_1.PLATFORM_DIRECTIVES, { useValue: [router_1.ROUTER_DIRECTIVES], multi: true }),
                core_1.provide(core_1.PLATFORM_PIPES, { useValue: [iterate_1.KeysPipe], multi: true }),
                router_1.provideRouter(routes_1.AppRoutes),
            ]);
        }
    }
});
//# sourceMappingURL=main.js.map