import {getFilter} from '../lib/filter';

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

      let start = 1440;
      let end = 0;
      let startTimes = {};
      let endTimes = {};
      let classLoads = {};

      schedule.sections.forEach(({course_id, section_id}) => {

        let course = state.detailedCourseCache[course_id];
        let section = course.sections.find(({id}) => id === section_id);

        section.meets.forEach(({days, start_time, end_time}) => {

          start = Math.min(start, start_time);
          end = Math.max(end, end_time);

          days.split('').forEach(day => {

            startTimes[day] = Math.min(startTimes[day] || 1440, start_time);
            endTimes[day] = Math.max(endTimes[day] || 0, end_time);



            if (!classLoads[day]) classLoads[day] = 1;
            else classLoads[day]++;

          });

        });

        courses.push({name: course.subject + ' ' + course.number + ' ' + section.name, meets: section.meets});

      });

      return {courses, gpa: schedule.average_gpa, start, end, length: end-start, startTimes, endTimes, classLoads};

    });

    state.schedulesCache[courses] = state.schedules;

  },

  RECEIVE_CACHED_SCHEDULES(state, schedules) {

    state.schedules = schedules;

  },

  CREATE_FILTER(state, filter) {

    let id = Math.max(0, ...state.filters.map(filter => filter.id))+1;

    state.filters.push(Object.assign({}, filter, {id}))

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
