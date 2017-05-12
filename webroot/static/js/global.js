'use strict';

angular.module('microStock')
  .controller('GlobalCtrl',['$scope',function($scope){
    $scope.loggedIn = false;

    function setCookie(cname, cvalue, exdays) {
      var d = new Date();
      d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
      var expires = "expires="+d.toUTCString();
      document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
    }

    $scope.login = function(){
      if($scope.loggedIn){
        setCookie("session","logoff",1);
        $scope.loggedIn = false;
      }else {
        setCookie("session","testsessionkey",1);
        $scope.loggedIn = true;
      }
    };
  }]);
