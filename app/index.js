import Vue from 'vue';
import Vuex from 'vuex';
import VueRouter from 'vue-router';

import state from './store/state';
import getters from './store/getters';
import mutations from './store/mutations';
import actions from './store/actions';

import App from './components/app.vue';
import Courses from './views/courses.vue';
import Filters from './views/filters.vue';
import Schedules from './views/schedules.vue';

Vue.use(Vuex);
Vue.use(VueRouter);

let store = new Vuex.Store({state, getters, mutations, actions});

let router = new VueRouter({

  mode: 'history',

  routes: [

    {path: '/', redirect: '/courses', component: App, children: [

      {path: '/courses', component: Courses},

      {path: '/filters', component: Filters},

      {path: '/schedules', component: Schedules}

    ]}

  ]

})

new Vue({store, router}).$mount('#app')
