import Vuex from 'vuex';

import mutations from './mutations';
import actions from './actions';

export default function createStore() {

  return new Vuex.Store({mutations, actions,

    state: {

      semesters: [],
      courses: [],

      selectedSemester: null,
      selectedCourses: [],
      coursesFilter: '',

      coursesCache: {},
      requestingSemesters: false,
      requestingCourses: false

    }

  })

}
