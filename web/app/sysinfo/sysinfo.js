'use strict';

angular.module('myApp.sysinfo', [
	'ui.router'
	])

.config(["$stateProvider", function($stateProvider) {
	$stateProvider
    .state('sysinfo', {
	    url: '/sysinfo',
	    templateUrl: 'sysinfo/sysinfo.html',
	    controller: 'SysInfoCtrl'
    })
}])


.controller('SysInfoCtrl', ['$scope', '$http', function($scope, $http) {
	$scope.sysinfoList = [];

	$scope.refreshSysInfo = function() {
		$http.get('s/sysinfo')
			.success(function(data, status) {
				$scope.sysinfoList = data;
				console.log('sysinfoList: ', $scope.sysinfoList);
			});
	};

	if ($scope.sysinfoList.length == 0) {
		$scope.refreshSysInfo();
	}

	$scope.meminfoList = [];

	$scope.refreshMemInfo = function() {
		$http.get('s/mem')
			.success(function(data, status) {
				// $scope.meminfoList = data;
				var freeValues = [];
				_.forEach(data, function(v) {freeValues.push({"Val":v.Free/(1024*1024), "At":v.At});});

				var totValues = [];
				_.forEach(data, function(v) {totValues.push({"Val":v.Total/(1024*1024), "At":v.At});});

				var usedValues = [];
				_.forEach(data, function(v) {usedValues.push({"Val":v.Used/(1024*1024), "At":v.At});});
				
				$scope.meminfoList = [
					/*
					{
						values: freeValues,
						key: "Free",
						color: "#0000ff"	
					},
					*/
					{
						values: totValues,
						key: "Total",
						color: "#ff0000"	
					},
					{
						values: usedValues,
						key: "Used",
						color: "#00ff00"	
					}
				];
				console.log('meminfoList: ', $scope.meminfoList);


			});
	};

	if ($scope.meminfoList.length == 0) {
		$scope.refreshMemInfo();
	}

	$scope.meminfoOptions = {
        chart: {
            type: 'lineChart',
            height: 250,
            width: 450,
            margin : {
                top: 20,
                right: 20,
                bottom: 40,
                left: 55
            },
            x: function(d){ return d.At; },
            y: function(d){ return d.Val; },
            useInteractiveGuideline: true,
            dispatch: {
            },

            xAxis: {
                axisLabel: 'Time (ms)',
                tickFormat: function(d) {
					return d3.time.format('%c')(new Date(d));
				}
            },
            yAxis: {
                axisLabel: 'Memory (MB)',
                //tickFormat: function(d){
                //    return d3.format('.02f')(d);
                //},
                axisLabelDistance: 30
            }        },
        title: {
            enable: true,
            text: 'Memory'
        }
    };

}]);