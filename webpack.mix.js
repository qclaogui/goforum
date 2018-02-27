let mix = require('laravel-mix');

mix.setPublicPath('public')
    .js('resources/assets/js/forum.js', 'public/js')
    .sass('resources/assets/sass/forum.scss', 'public/css')
    .version().disableNotifications();