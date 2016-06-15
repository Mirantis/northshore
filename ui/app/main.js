System.register(['@angular/platform-browser-dynamic', '@angular/http', './components/app/app'], function(exports_1, context_1) {
    "use strict";
    var __moduleName = context_1 && context_1.id;
    var platform_browser_dynamic_1, http_1, app_1;
    return {
        setters:[
            function (platform_browser_dynamic_1_1) {
                platform_browser_dynamic_1 = platform_browser_dynamic_1_1;
            },
            function (http_1_1) {
                http_1 = http_1_1;
            },
            function (app_1_1) {
                app_1 = app_1_1;
            }],
        execute: function() {
            platform_browser_dynamic_1.bootstrap(app_1.AppComponent, [http_1.HTTP_PROVIDERS]);
        }
    }
});
//# sourceMappingURL=main.js.map