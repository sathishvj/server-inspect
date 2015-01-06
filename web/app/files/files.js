'use strict';

angular.module('myApp.files', [
	'ui.router'
	])

.config(["$stateProvider", function($stateProvider) {

	$stateProvider
    .state('files', {
      url: '/files',
      templateUrl: 'files/files.html',
      controller: 'FilesCtrl'
    })
}])


.controller('FilesCtrl', ['$scope', '$http', function($scope, $http) {
	$scope.filesList = [];

	$scope.refreshEnv = function() {
		$http.get('s/files')
			.success(function(data, status) {
				$scope.filesList = data;
				console.log('filesList: ', $scope.filesList);
			});
	};

	if ($scope.filesList.length == 0) {
		$scope.refreshEnv();
	}

	$scope.tailFile = function() {
		var ws = new WebSocket("ws://localhost:8083/ws/files/tail");
	    ws.onopen = function()
	    {
	       // Web Socket is connected, send data using send()
	       ws.send("Message to send");
	       console.log("Message is sent...");
	    };
	    ws.onmessage = function (evt) 
	    { 
	       var received_msg = evt.data;
	       console.log("Message is received...", received_msg);
	    };
	    ws.onclose = function()
	    { 
	       // websocket is closed.
	       console.log("Connection is closed..."); 
	    };
	};

}]);