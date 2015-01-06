'use strict';

// Declare app level module which depends on views, and components
angular.module('myApp', [
  'myApp.sysinfo',
  'myApp.env',
  'myApp.files',
  'myApp.version',
  'myApp.filters',
  'ngMaterial',
  'ui.grid',
  'ui.router',
  'angular-loading-bar',
  'nvd3'
])

.config(["$stateProvider", "$urlRouterProvider", function($stateProvider, $urlRouterProvider) {
	$urlRouterProvider.otherwise("/");

	$stateProvider
    .state('/', {
      url: ''
    })
}])

.controller('headerController', function($scope) {
	$scope.sideMenuItems = [
	        {
		        "Name": "Files",
				"Link": "files"
			},
	        {
		        "Name": "Env Vars",
				"Link": "env"
			},
	        {
		        "Name": "System Info",
				"Link": "sysinfo"
			}
		];
	})
;
