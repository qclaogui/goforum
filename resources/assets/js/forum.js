/**
 * First we will load all of this project's JavaScript dependencies which
 * includes Vue and other libraries.
 */

require('./bootstrap');

window.Vue = require('vue');



window.events = new Vue();
window.flash = function (message) {
    window.events.$emit('flash', message)
};

/**
 * Next, we will create a fresh Vue application instance and attach it to
 * the page. Then, you may begin adding components to this application
 * or customize the JavaScript scaffolding to fit your unique needs.
 */

Vue.component('flash-msg', require('./components/flash-msg.vue'));
Vue.component('home-component', require('./components/home.vue'));
Vue.component('reply-component', require('./components/reply.vue'));

const app = new Vue({
    el: '#app'
});