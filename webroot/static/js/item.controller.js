'use strict';

angular.module('microStock')
  .controller('ItemCtrl',['$scope','$http','$stateParams',function($scope,$http,$stateParams){
    $scope.obj = {};
    $scope.list = [];

    function load(){
      $http.get(config.microservice_dependencies.productById.replace("%d", $stateParams.productid)).then(function(res) {
        $scope.obj = res.data
      });
      $http.get(config.store.goods.productById.replace("%d",$stateParams.productid)).then(function(res) {
        $scope.list = res.data;
      },function(){
        $scope.list = [];
      });
    }
    load();
    $scope.delete = function(id){
      $http.delete(config.store.goods.productById.replace("%d",id)).then(load);
    }
  }]);
