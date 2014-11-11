angular.module('app',[])
  .controller('BlogIndexController', ['$scope', '$http', function($scope, $http) {
    $http.get('/api/v1/blogs').success(function(data) {
      $scope.blogs = data.Pages;
    });
  }]);
