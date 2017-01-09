export default {

  REQUEST_SEMESTERS(state) {

    state.requestingSemesters = true;

  },

  RECEIVE_SEMESTERS(state, semesters) {

    state.semesters = semesters;
    state.requestingSemesters = false;

  },

  SELECT_SEMESTER(state, semester) {

    if (semester != state.selectedSemester) {

      state.selectedCourses = [];
      state.selectedSemester = semester;

    }

  },

  REQUEST_COURSES(state) {

    state.requestingCourses = true;

  },

  RECEIVE_COURSES(state, {semester, courses}) {

    state.coursesCache[semester.id] = courses;
    state.courses = courses;
    state.requestingCourses = false;

  },

  SELECT_COURSE(state, course) {

    state.selectedCourses.push(course);

  },

  DESELECT_COURSE(state, course) {

    let index = state.selectedCourses.indexOf(course);
    state.selectedCourses.splice(index, 1);

  },

  SET_FILTER(state, filter) {

    state.coursesFilter = filter;

  }

}
