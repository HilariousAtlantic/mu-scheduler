export default {

  REQUEST_SEMESTERS(state) {

  },

  RECEIVE_SEMESTERS(state, semesters) {

    state.semesters = semesters;

  },

  SELECT_SEMESTER(state, semester) {

    state.selectedSemester = semester;

  },

  REQUEST_COURSES(state, semester) {

    state.requestingCourses = true;

  },

  RECEIVE_COURSES(state, {semester, courses}) {

    state.courses[semester.id] = courses;
    state.requestingCourses = false;

    console.log(state.courses);

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
