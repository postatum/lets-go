'use strict';

var peopleApp = angular.module('peopleApp', [
  'ngRoute',

  'peopleControllers',
],
function($interpolateProvider) {
  $interpolateProvider.startSymbol('[[');
  $interpolateProvider.endSymbol(']]');
});

// Routes
peopleApp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/people', {
        templateUrl: 'partials/peopleList.html',
        controller: 'PeopleListCtrl'
      }).
      otherwise({
        redirectTo: '/people'
      });
  }]);
