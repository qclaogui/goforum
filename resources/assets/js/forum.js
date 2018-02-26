/**
 * First we will load all of this project's JavaScript dependencies which
 * includes Vue and other libraries.
 */

require('./bootstrap');

window.Vue = require('vue');

/**
 * Next, we will create a fresh Vue application instance and attach it to
 * the page. Then, you may begin adding components to this application
 * or customize the JavaScript scaffolding to fit your unique needs.
 */

Vue.component('example-component', require('./components/example-component.vue'));
Vue.component('home-component', require('./components/home-component.vue'));

const app = new Vue({
    el: '#app'
});