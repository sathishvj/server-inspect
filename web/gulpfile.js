var gulp = require('gulp')
	, sass = require('gulp-sass')
	, autoprefixer = require('gulp-autoprefixer')
	, minifycss = require('gulp-minify-css')
	, rename = require('gulp-rename')
	, util = require('gulp-util')
	, plumber = require('gulp-plumber')
	;

gulp.task('sass', function() {
	return gulp.src('app/sass/*.scss')
		.pipe(plumber({
	        errorHandler: function (err) {
	            console.log(err.plugin, err.message);
	            this.emit('end');
	            util.beep();
	        }
	    }))
		.pipe(sass({
			style: 'expanded',
			keepBreaks: 'true',
			keepSpecialComments: '*'
		}))
		.pipe(autoprefixer())
		.pipe(gulp.dest('app/css'))
		.pipe(rename({suffix: '.min'}))
		.pipe(minifycss())
		.pipe(gulp.dest('app/css'))
		;
})
;

gulp.task('watch', function() {
	gulp.watch('app/sass/*.scss', ['sass']);
})
;

gulp.task('default', ['watch'], function() {

})
;