import {getDefaultOptions, getFilter} from '../lib/filter';

export default {

  REQUEST_TERMS(state) {

    state.requestingTerms = true;

  },

  RECEIVE_TERMS(state, terms) {

    state.requestingTerms = false;

    state.terms = terms.sort((termA, termB) => {

      let seasons = ['Winter', 'Spring', 'Summer', 'Fall'];

      let [seasonA, typeA, yearA] = termA.name.split(' ');
      let [seasonB, typeB, yearB] = termB.name.split(' ');

      if (yearA === yearB) {

        return seasons.indexOf(seasonB) - seasons.indexOf(seasonA);

      } else {

        return parseInt(yearB) - parseInt(yearA);

      }

    });

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

      let timesByDay = {M: [], T: [], W: [], R: [], F: []};
      let startTimes = {M: 1440, T: 1440, W: 1440, R: 1440, F: 1440};
      let endTimes = {M: 0, T: 0, W: 0, R: 0, F: 0};
      let classLoads = {M: 0, T: 0, W: 0, R: 0, F: 0};

      schedule.sections.forEach(({course_id, section_id}) => {

        let course = state.detailedCourseCache[course_id];
        let section = course.sections.find(({id}) => id === section_id);

        section.meets.forEach(({days, start_time, end_time}) => {

          days.split('').forEach(day => {

            timesByDay[day].push({start: start_time, end: end_time});
            startTimes[day] = Math.min(startTimes[day], start_time);
            endTimes[day] = Math.max(endTimes[day], end_time);
            classLoads[day]++;

          });

        });

        courses.push({
          name: course.subject + ' ' + course.number + ' ' + section.name,
          meets: section.meets
        });

      });

      let start = Math.min(...Object.values(startTimes));
      let end = Math.max(...Object.values(endTimes));
      let length = end-start;
      let gpa = schedule.average_gpa;

      return {courses, gpa, start, end, length, timesByDay, startTimes, endTimes, classLoads};

    }).sort((a,b) => b.gpa - a.gpa);

    state.schedulesCache[courses] = state.schedules;

  },

  RECEIVE_CACHED_SCHEDULES(state, schedules) {

    state.schedules = schedules;

  },

  CREATE_FILTER(state, type) {

    let id = Math.max(0, ...state.filters.map(filter => filter.id))+1;
    let options = getDefaultOptions(type);
    let test = getFilter(type, options);

    state.filters.push({id, type, options, active: false});

  },

  TOGGLE_FILTER(state, id) {

    let filter = state.filters.find(filter => filter.id === id);

    filter.active = !filter.active;

  },

  CHANGE_FILTER(state, {id, changes}) {

    let filter = state.filters.find(filter => filter.id === id);

    filter.options = changes;

  },

  UPDATE_FILTER(state, id) {

    let filter = state.filters.find(filter => filter.id === id);

    filter.test = getFilter(filter.type, filter.options);

  }

}
