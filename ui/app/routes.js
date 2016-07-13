System.register(['./components/blueprints/blueprints', './components/home/home'], function(exports_1, context_1) {
    "use strict";
    var __moduleName = context_1 && context_1.id;
    var blueprints_1, home_1;
    var AppRoutes;
    return {
        setters:[
            function (blueprints_1_1) {
                blueprints_1 = blueprints_1_1;
            },
            function (home_1_1) {
                home_1 = home_1_1;
            }],
        execute: function() {
            exports_1("AppRoutes", AppRoutes = [
                {
                    path: '',
                    component: home_1.HomeComponent,
                },
                {
                    path: 'blueprints',
                    component: blueprints_1.BlueprintsComponent,
                },
                {
                    path: 'blueprints/:uuid',
                    component: blueprints_1.BlueprintsComponent,
                },
                {
                    path: '**',
                    redirectTo: '',
                },
            ]);
        }
    }
});
//# sourceMappingURL=routes.js.map