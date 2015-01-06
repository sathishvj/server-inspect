'use strict';

angular.module('myApp.env', [
	'ui.router'
	])

.config(["$stateProvider", function($stateProvider) {

	$stateProvider
    .state('env', {
	    url: '/env',
	    templateUrl: 'env/env.html',
	    controller: 'EnvCtrl'
    })
}])


.controller('EnvCtrl', ['$scope', '$http', function($scope, $http) {
	$scope.envList = [];

	$scope.refreshEnv = function() {
		$http.get('s/env')
			.success(function(data, status) {
				$scope.envList = data.sort();
				console.log('envList: ', $scope.envList);
			});
	};

	if ($scope.envList.length == 0) {
		$scope.refreshEnv();
	}
}]);