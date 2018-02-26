let mix = require('laravel-mix');

mix.js('resources/assets/js/forum.js', 'public/js')
    .sass('resources/assets/sass/forum.scss', 'public/css');