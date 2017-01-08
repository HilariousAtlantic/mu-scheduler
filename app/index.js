import Vue from 'vue';
import Vuex from 'vuex';
import VueRouter from 'vue-router';

import App from './components/app.vue';
import CourseSelection from './pages/course-selection.vue';
import ScheduleSelection from './pages/schedule-selection.vue';

import createStore from './state/store';

Vue.use(Vuex);
Vue.use(VueRouter);

let router = new VueRouter({

  mode: 'history',

  routes: [

    {path: '/', redirect: '/courses', component: App, children: [

      {path: '/courses', component: CourseSelection},

      {path: '/schedules', component: ScheduleSelection}

    ]}

  ]

})

new Vue({store: createStore(), router}).$mount('#app')
