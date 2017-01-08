import Vue from 'vue';
import VueRouter from 'vue-router';

import App from './components/app.vue';
import CourseSelection from './pages/course-selection.vue';
import ScheduleSelection from './pages/schedule-selection.vue';

Vue.use(VueRouter);

let router = new VueRouter({

  mode: 'history',

  routes: [

    {path: '/', component: App, children: [

      {path: 'courses', component: CourseSelection, alias: ''},

      {path: 'schedules', component: ScheduleSelection}

    ]}

  ]

})

new Vue({router}).$mount('#app')
