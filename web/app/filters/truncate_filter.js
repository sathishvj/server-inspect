'use strict';

angular.module('myApp.filters')
.filter('truncate', function (){
  return function (text, length, end){
    if (text !== undefined){
      if (isNaN(length)){
        length = 40;
      }

      var ellipsis = "...";

      if (text.length <= length || text.length - ellipsis.length <= length){
        return text;
      }else{
        return ellipsis + String(text).substring(text.length - length - ellipsis.length, text.length);
      }
    }
  };
});