'use strict';

angular.module('microStock')
  .controller('GlobalCtrl',['$scope','$http', function($scope, $http){
    $scope.loggedIn = false;

    $scope.login = function(){
      if($scope.loggedIn){
        $http.defaults.headers.common["session"] = "logoff";
        $scope.loggedIn = false;
      }else {
        $http.defaults.headers.common["session"] = "testsessionkey";
        $scope.loggedIn = true;
      }
    };
  }]);
