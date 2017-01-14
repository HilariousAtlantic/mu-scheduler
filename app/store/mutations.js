export default {

  REQUEST_TERMS(state) {

    state.requestingTerms = true;

  },

  RECEIVE_TERMS(state, terms) {

    state.terms = terms;
    state.requestingTerms = false;

  },

  SELECT_TERM(state, term) {

    if (term != state.selectedTerm) {

      state.selectedCourses = [];
      state.selectedTerm = term;

    }

  },

  REQUEST_COURSES(state) {

    state.requestingCourses = true;

  },

  RECEIVE_COURSES(state, {term, courses}) {

    state.coursesCache[term.id] = courses;
    state.courses = courses;
    state.requestingCourses = false;

  },

  RECEIVE_CACHED_COURSES(state, courses) {

    state.courses = courses;

  },

  REQUEST_DETAILED_COURSE(state) {

    state.requestingDetailedCourse = true;

  },

  RECEIVE_DETAILED_COURSE(state, detailedCourse) {

    state.detailedCourseCache[detailedCourse.id] = detailedCourse;
    state.requestingDetailedCourse = false;

  },

  SELECT_COURSE(state, course) {

    console.log(course);

    let index = state.selectedCourses.indexOf(course);

    if (index == -1) {

      state.selectedCourses.push(course);

    }

  },

  DESELECT_COURSE(state, course) {

    let index = state.selectedCourses.indexOf(course);

    if (index != -1) {

      state.selectedCourses.splice(index, 1);

    }

  },

  SET_FILTER(state, filter) {

    state.coursesFilter = filter;

  }

}
