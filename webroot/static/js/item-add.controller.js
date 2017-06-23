'use strict';

angular.module('microStock')
  .controller('ItemAddCtrl',['$scope','$http','$stateParams',function($scope,$http,$stateParams){
    $scope.product = {};
    $scope.obj = {};
    $scope.msg = {};
    $scope.count = 1;

    $http.get(config.microservice_dependencies.productById.replace("%d", $stateParams.productid)).then(function(res) {
      $scope.product = res.data
    });

    $scope.submit = function(){
      $http.post(config.store.goods.productById.replace("%d",$stateParams.productid)+'?count='+$scope.count,$scope.obj).then(function(){
        $scope.obj = {};
        $scope.msg = {type:'success',text:'Saved '+$scope.count+' good(s) from product '+$scope.product.name+'.'};
      }, function(e){
        if(e.status == 403){
          $scope.msg = {type:'error',text:'You are not allowed to add goods, maybe you should login!'};
        }else{
          $scope.msg = {type:'error',text:'An error occurred while saving good(s): '+e.data};
        }
      });
    };
  }]);
