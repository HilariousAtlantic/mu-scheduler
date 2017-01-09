import axios from 'axios';

export default {

  fetchSemesters({commit}) {

    commit('REQUEST_SEMESTERS');

    axios.get('/api/semesters').then(response => {

      let semesters = response.data;

      commit('RECEIVE_SEMESTERS', semesters);

    });

  },

  selectSemester({commit, state}, semester) {

    commit('SELECT_SEMESTER', semester);

    let courses = state.coursesCache[semester.id];

    if (courses) {

      commit('RECEIVE_CACHED_COURSES', courses);

    } else {

      commit('REQUEST_COURSES');

      axios.get('/api/courses').then(response => {

        let courses = response.data;

        commit('RECEIVE_COURSES', {semester, courses});

      });

    }

  },

  setFilter({commit}, filter) {

    commit('SET_FILTER', filter);

  },

  selectCourse({commit}, course) {

    commit('SELECT_COURSE', course);

  },

  deselectCourse({commit}, course) {

    commit('DESELECT_COURSE', course);

  }

}
