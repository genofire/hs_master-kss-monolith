'use strict';

angular.module('microStock')
  .controller('StatisticsCtrl',['$scope','$http',function($scope,$http){
    $scope.obj = {};

    $http.get(config.store.status).then(function(res) {
      $scope.obj = res.data
    });
  }]);
