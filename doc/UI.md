NorthShore UI
=============

NSUI is Web UI.
The frontend part is Angular 2 application.

The Angular 2 framework allows coding on JavaScript, TypeScript and Dart.
There we choose the TypeScript for the current project.
Some useful info about editor support for TypeScript can be found at
[TypeScript Editor Support](https://github.com/Microsoft/TypeScript/wiki/TypeScript-Editor-Support)


Development Convention
----------------------

* Use style-guides and best-practices
* RTFM
  - [A2 Style Guide](https://angular.io/docs/ts/latest/guide/style-guide.html)
  - [johnpapa/angular-styleguide](https://github.com/johnpapa/angular-styleguide)
  - [JSON API](http://jsonapi.org/)


Client Side Routing
-------------------

There is the specific of single_page_apps with client-side routing.

To handle the current application state such as the current view the location params are used.
It allows store/share link with the application state.

Previously the anchors were used (/appname#view_1).
And since browsers implemented the HTML5 History API the separate links are used (/appname/view_1).
It requires support from the Backend: when user reloads the browser on some app-handled route the app index.html should be returned instead of 404 error.

So the backend have to know if 404 is allowed.


Backend Routing
---------------

In the NSUI [a2app example](https://github.com/johnpapa/angular2-tour-of-heroes/blob/master/index.html) we have few request types:

- GET for index.html from brsr, even on reload at some a2app route, the location depends on the _base_ tag:

        /
        /dashboard_view_1
        /dashboard_view_2/action

- GET for static that hardcoded in tags (img, link, script), the location depends on the _base_ tag:

        src="node_modules/es6-shim/es6-shim.min.js"
        src="node_modules/angular2/bundles/angular2-polyfills.js"

- XHR GET for static JS called by SystemJS loader, the location depends on the System.config({baseURL}):

        .import('app/main')

- XHR GET for static from a2app, the location depends on the _base_ tag:

        styleUrls: ['app/dashboard.component.css']
        templateUrl: 'app/dashboard.component.html'

- XHR API calls from a2app, the location in general hardcoded as

        prefix + callname => /api/v1/callname


So, we have to handle few cases:

- 404 allowed, static resources such as CSS, IMG, node_modules
- 404 disallowed, index.html instead on a2app routes

And there are 2 different a2app config options:

- the _base_ tag, in general '/' for clear url-line in one_app_domain
- System.config({baseURL}), that partially affects a2app

Instead of using some special prefix in a2app for static XHR GET, let's handle reqs as:

- \* - default - 404 disallowed, index.html instead
- ^/api/* - starts with _/api/_ - 404 allowed, API for a2app
- ^/app/* - starts with _/app/_ - 404 allowed, static resources for a2app
- ^/assets/* - starts with _/assets/_ - 404 allowed, static resources such as CSS, IMG, third party libs
- ^/node_modules/* - starts with _/node_modules/_ - 404 allowed, static resources for node_modules

That scheme allows integration with third party services - the default should be defined as a2app route list.

Add the "ui" prefix to routes, it's combines and marks them as targeted for UI.
The allowed routes start with BASE (ie `http://host:port/`) and so:

- BASE/ui/api/
- BASE/ui/app/
- BASE/ui/assets/
- BASE/ui/node_modules/


UI API Scheme
-------------

There we choose the [JSON API](http://jsonapi.org/) for scheme of calls from
the Frontend to the Backend for data and actions.


Local Environment
-----------------

Angular application developers rely on the _npm_ package manager to install the libraries and packages their apps require.

* Install NodeJS

        curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash -

    It adds [deb.nodesource.com](https://deb.nodesource.com/node_6.x) repository
    and updates apt-get lists.

        sudo apt-get install nodejs

    Check installed

        node -v
          v6.2.0
        npm -v
          3.8.9

* Adding the libraries and packages we need with _npm_
  and compile the TS sources into JS

        cd ${GOPATH}/src/github.com/Mirantis/northshore
        cd ./ui
        npm install
        npm run tsc

* Run local Backend server

        cd ${GOPATH}/src/github.com/Mirantis/northshore
        go run cmd/nshore/nshore.go run local

* Run browser http://localhost:8998/ui


Setup Development Environment with Atom on Ubuntu
-------------------------------------------------

* RTFM [Atom TypeScript](https://atom.io/packages/atom-typescript)

* Install Atom

        sudo add-apt-repository ppa:webupd8team/atom
        sudo apt-get update
        sudo apt-get install atom

        apm install atom-typescript
        apm install linter
        apm install editorconfig

* Run local _lite-server_ for UI Development

  There are few helpful scripts from  [angular/quickstart](https://angular.io/docs/ts/latest/quickstart.html#!#config-files).

        cd ${GOPATH}/src/github.com/Mirantis/northshore
        cd ./ui
        npm start

  It runs the compiler and a server at the same time, both in "watch mode" for  changes to TypeScript files and recompiling when it sees them.
  There is the proxy for the API calls to the Backend http://localhost:8998.

* Run browser http://localhost:3000
