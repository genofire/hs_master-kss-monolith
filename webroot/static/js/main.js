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
      .state('item-add', {
        url: '/product/:productid/add',
        templateUrl: '/static/html/item-add.html',
        controller:  'ItemAddCtrl'
      })
      .state('statistics', {
        url: '/statistics',
        templateUrl: '/static/html/statistics.html',
        controller:  'StatisticsCtrl'
      });
  }])
  .filter('reloadSrc', function () {
    return function (input) {
      if (input)
          return input + '?v=' + new Date().getTime();
      }
  });


var config = {
  'store': {
    'status': '/api/status',
    'goods': {
      'productById': '/api/good/%d'
    },
  },
  'microservice_dependencies': {
    'products': '/api-test/product/',
    'productById': '/api-test/product/%d/'
  }
};
