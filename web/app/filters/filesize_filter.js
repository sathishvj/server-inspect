'use strict';

angular.module('myApp.filters', [])

.filter('filesize', function () {
  var units = [
    'b',
    'KB',
    'MB',
    'GB',
    'TB',
    'PB'
  ];

  return function( bytes, precision ) {
    if ( isNaN( parseFloat( bytes )) || ! isFinite( bytes ) ) {
      return '?';
    }

    var unit = 0;

    while ( bytes >= 1024 ) {
      bytes /= 1024;
      unit ++;
    }

    return bytes.toFixed( + precision ) + '' + units[ unit ];
  };
})
;