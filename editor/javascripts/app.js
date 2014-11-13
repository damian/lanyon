var app = angular.module('app', [
  'ngRoute'
]);

app.controller('BlogIndexController', ['$scope', '$http', function($scope, $http) {
  $http.get('/api/v1/blogs').success(function(data) {
    $scope.blogs = data.Pages;
  });
}]);

app.controller('BlogShowController', ['$scope', '$routeParams', '$http', '$sce', function($scope, $routeParams, $http, $sce) {
  $http.get('/api/v1/blogs/' + $routeParams.blogId).success(function(data) {
    $scope.blog = data;
    $scope.blog.SanitizedFormattedBody = $sce.trustAsHtml(data.FormattedBody);
  });
}]);

app.controller('BlogEditController', ['$scope', '$routeParams', '$http', '$sce', function($scope, $routeParams, $http, $sce) {
  $http.get('/api/v1/blogs/' + $routeParams.blogId).success(function(data) {
    $scope.blog = data;
  });
}]);

app.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/blog', {
    templateUrl: 'partials/blog.html',
    controller: 'BlogIndexController',
  }).
  when('/:blogId', {
    templateUrl: 'partials/show.html',
    controller: 'BlogShowController'
  }).
  when('/:blogId/edit', {
    templateUrl: 'partials/edit.html',
    controller: 'BlogEditController'
  }).
  when('/', {
    templateUrl: 'partials/index.html'
  }).otherwise({
    redirectTo: '/'
  });
}]);
