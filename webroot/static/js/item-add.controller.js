'use strict';

angular.module('microStock')
  .controller('ItemAddCtrl',['$scope','$http','$stateParams',function($scope,$http,$stateParams){
    $scope.product = {};
    $scope.obj = {};
    $scope.count = 1;

    $http.get(config.microservice_dependencies.productById.replace("%d", $stateParams.productid)).then(function(res) {
      $scope.product = res.data
    });

    $scope.submit = function(){
      function request(){
        return $http.post(config.store.goods.productById.replace("%d",$stateParams.productid),$scope.obj);
      }
      var last = request();
      for(var i=1;i < $scope.count;i++){
        last.then(request);
      }
      last.then(null,function(){
        console.log("did not work");
      })
    };
  }]);
