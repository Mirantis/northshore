"use strict";

var browserSync = require('browser-sync');
var proxyMiddleware = require('http-proxy-middleware');
var proxyApi = proxyMiddleware('/ui/api', {target: 'http://localhost:8998/'});


module.exports = {
  server: {
    middleware: [proxyApi],
    routes: {
      "/ui": "./",
    },
  }
};
