'use strict';

angular.module('microStock')
  .controller('ListCtrl',['$scope','$http',function($scope,$http){
    $scope.list = [];
    $http.get(config.microservice_dependencies.products).then(function(res) {
        $scope.list = res.data
      });
  }]);
