import axios from 'axios';

export default {

  fetchTerms({commit}) {

    commit('REQUEST_TERMS');

    axios.get('/api/terms').then(response => {

      let terms = response.data;

      commit('RECEIVE_TERMS', terms);

    });

  },

  selectTerm({commit, state}, term) {

    commit('SELECT_TERM', term);

    let courses = state.coursesCache[term.id];

    if (courses) {

      commit('RECEIVE_CACHED_COURSES', courses);

    } else {

      commit('REQUEST_COURSES');

      axios.get('/api/courses?term='+term.id).then(response => {

        let courses = response.data;

        commit('RECEIVE_COURSES', {term, courses});

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
