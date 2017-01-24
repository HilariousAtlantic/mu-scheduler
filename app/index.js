import Vue from 'vue';
import Vuex from 'vuex';
import VueRouter from 'vue-router';

import state from './store/state';
import getters from './store/getters';
import mutations from './store/mutations';
import actions from './store/actions';

import App from './components/app.vue';
import Generator from './components/generator.vue';

import Courses from './views/courses.vue';
import Filters from './views/filters.vue';
import Schedules from './views/schedules.vue';

Vue.use(Vuex);
Vue.use(VueRouter);

let store = new Vuex.Store({state, getters, mutations, actions});

let router = new VueRouter({

  mode: 'history',

  routes: [

    {path: '/', redirect: '/generator/courses', component: App, children: [

      {path: '/generator', component: Generator, children: [

        {name: 'courses', path: '/generator/courses', component: Courses},

        {name: 'filters', path: '/generator/filters', component: Filters},

        {name: 'schedules', path: '/generator/schedules', component: Schedules}

      ]}

    ]}

  ]

})

new Vue({store, router}).$mount('#app')
