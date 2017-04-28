angular.module('microStock', [
    'angular-loading-bar',
    'ngAnimate',
    'ui.router'
  ])
  .config(['$urlRouterProvider','$httpProvider',function($urlRouterProvider,$httpProvider){
    $urlRouterProvider.otherwise('/');
    $httpProvider.defaults.withCredentials = true;
  }])
  .config(['$stateProvider',function ($stateProvider) {
    $stateProvider
      .state('list', {
        url: '/',
        templateUrl: '/static/html/list.html',
        controller:  'ListCtrl'
      })
      .state('item', {
        url: '/product/:productid',
        templateUrl: '/static/html/item.html',
        controller:  'ItemCtrl'
      })
      .state('statistics', {
        url: '/statistics',
        templateUrl: '/static/html/statistics.html',
        controller:  'StatisticsCtrl'
      });
  }]);


var config = {
  'store': {
    'status': '/api/status',
    'goods': {
      'productById': '/api/good/%d'
    },
  },
  'microservice_dependencies': {
    'products': 'http://localhost:8080/api-test/product/',
    'productById': 'http://localhost:8080/api-test/product/%d/'
  }
};