import Vue from 'vue';
import Vuex from 'vuex';
import VueRouter from 'vue-router';

import state from './store/state';
import getters from './store/getters';
import mutations from './store/mutations';
import actions from './store/actions';

import Generator from './views/generator.vue';
import Courses from './views/courses.vue';
import Schedules from './views/schedules.vue';

Vue.use(Vuex);
Vue.use(VueRouter);

let store = new Vuex.Store({state, getters, mutations, actions});

let router = new VueRouter({

  mode: 'history',

  routes: [

    {path: '/', redirect: 'courses', component: Generator, children: [

      {name: 'courses', path: '/courses', component: Courses},

      {name: 'schedules', path: '/schedules', component: Schedules}

    ]},

    {path: '*', redirect: 'courses'}

  ]

})

new Vue({store, router}).$mount('#app')
