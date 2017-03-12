import {getDefaultOptions, getFilter} from '../lib/filter';

export default {

  REQUEST_TERMS(state) {

    state.requestingTerms = true;

  },

  RECEIVE_TERMS(state, terms) {

    state.requestingTerms = false;

    state.terms = terms;

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

    state.schedules = schedules;

    state.schedulesCache[courses] = schedules;

  },

  RECEIVE_CACHED_SCHEDULES(state, schedules) {

    state.schedules = schedules;

  },

  CREATE_FILTER(state, type) {

    let id = Math.max(0, ...state.filters.map(filter => filter.id))+1;
    let options = getDefaultOptions(type);

    state.filters.push({id, type, options, active: false});

  },

  TOGGLE_FILTER(state, id) {

    state.filters = state.filters.map(filter => {

      if (filter.id === id) {

        return {...filter, active: !filter.active}

      } else return filter;

    });

  },

  CHANGE_FILTER(state, {id, changes}) {

    state.filters = state.filters.map(filter => {

      if (filter.id === id) {

        return {...filter, changes};

      } else return filter;

    });

  },

  DELETE_FILTER(state, id) {

    state.filters = state.filters.filter(filter => filter.id !== id);

  },

  APPLY_FILTERS(state) {

    state.filters = state.filters.map(filter => {

      if (filter.changes) {

        let {id, type, active, changes} = filter;

        return {id, type, active, options: changes, test: getFilter(type, changes)};

      }

      return filter;

    });

  }

}
