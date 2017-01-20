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

    let index = state.selectedCourses.indexOf(course.id);

    if (index == -1) {

      state.selectedCourses.push(course.id);

    }

  },

  DESELECT_COURSE(state, course) {

    let index = state.selectedCourses.indexOf(course.id);

    if (index != -1) {

      state.selectedCourses.splice(index, 1);

    }

  },

  UPDATE_COURSE_FILTER(state, updatedFilter) {

    state.courseFilter = updatedFilter;

  },

  REQUEST_SCHEDULES(state) {

    state.requestingSchedules = true;

  },

  RECEIVE_SCHEDULES(state, {courses, schedules}) {

    state.requestingSchedules = false;

    state.schedules = schedules.map(schedule => {

      let courses = [];

      schedule.sections.forEach(({course_id, section_id}) => {

        let course = state.detailedCourseCache[course_id];
        let section = course.sections.find(({id}) => id === section_id);

        courses.push({name: course.subject + ' ' + course.number + ' ' + section.name, meets: section.meets});

      });

      return {courses};

    });

    state.schedulesCache[courses] = state.schedules;

  },

  RECEIVE_CACHED_SCHEDULES(state, schedules) {

    state.schedules = schedules;

  },

  TOGGLE_SCHEDULE_FILTER(state, id) {

    let filter = state.scheduleFilters.find(filter => filter.id === id);

    Object.assign(filter, {active: !filter.active});

  },

  UPDATE_SCHEDULE_FILTER(state, {id, changes}) {

    let filter = state.scheduleFilters.find(filter => filter.id === id);

    Object.assign(filter.options, changes);

  }

}
