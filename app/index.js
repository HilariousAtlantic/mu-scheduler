import Vue from 'vue';
import Vuex from 'vuex';
import VueRouter from 'vue-router';

import state from './store/state';
import mutations from './store/mutations';
import actions from './store/actions';

import App from './components/app.vue';
import CourseSelection from './pages/course-selection.vue';
import ScheduleSelection from './pages/schedule-selection.vue';

Vue.use(Vuex);
Vue.use(VueRouter);

let store = new Vuex.Store({state, mutations, actions});

let router = new VueRouter({

  mode: 'history',

  routes: [

    {path: '/', redirect: '/courses', component: App, children: [

      {path: '/courses', component: CourseSelection},

      {path: '/schedules', component: ScheduleSelection}

    ]}

  ]

})

new Vue({store, router}).$mount('#app')
