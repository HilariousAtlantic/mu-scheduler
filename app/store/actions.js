import axios from 'axios';

export default {

  fetchTerms({commit, dispatch}) {

    commit('REQUEST_TERMS');

    return axios.get('/api/terms').then(response => {

      let terms = response.data.sort((termA, termB) => {

        let seasons = ['Winter', 'Spring', 'Summer', 'Fall'];

        let [seasonA, , yearA] = termA.name.split(' ');
        let [seasonB, , yearB] = termB.name.split(' ');

        if (yearA === yearB) {

          return seasons.indexOf(seasonB) - seasons.indexOf(seasonA);

        } else {

          return parseInt(yearB) - parseInt(yearA);

        }

      });

      commit('RECEIVE_TERMS', terms);

      dispatch('selectTerm', terms[0]);

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

        let courses = response.data.map(course =>
          Object.assign({}, course, {searchableName: course.subject + ' ' + course.number + ' - ' + course.title})
        ).sort((a, b) =>
          a.searchableName.localeCompare(b.searchableName)
        );

        commit('RECEIVE_COURSES', {term, courses});

      });

    }

  },

  updateCourseFilter({commit}, filter) {

    commit('UPDATE_COURSE_FILTER', filter);

  },

  selectCourse({commit, state}, course) {

    let detailedCourse = state.detailedCourseCache[course.id];

    if (!detailedCourse) {

      commit('REQUEST_DETAILED_COURSE');

      axios.get('/api/courses/'+course.id).then(response => {

        let detailedCourse = {...response.data,

          sections: response.data.sections.sort((a,b) =>
            a.name.localeCompare(b.name)
          )

        };

        commit('RECEIVE_DETAILED_COURSE', detailedCourse);

        commit('SELECT_COURSE', detailedCourse);

      });

    } else {

      commit('SELECT_COURSE', course);

    }

  },

  deselectCourse({commit}, course) {

    commit('DESELECT_COURSE', course);

  },

  generateSchedules({commit, state}) {

    if (state.selectedCourses.length === 0) return Promise.resolve(false);

    let courses = state.selectedCourses.join(',');

    let schedules = state.schedulesCache[courses];

    if (!schedules) {

      commit('REQUEST_SCHEDULES');

      return axios.get('/api/schedules?courses='+courses).then(response => {

          let schedules = response.data.map(schedule => {

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

          commit('RECEIVE_SCHEDULES', {courses, schedules});

          return true;

      });

    } else {

      commit('RECEIVE_CACHED_SCHEDULES', schedules);

      return Promise.resolve(true);

    }

  },

  createFilter({commit}, type) {

    commit('CREATE_FILTER', type);

  },

  toggleFilter({commit}, id) {

    commit('TOGGLE_FILTER', id);

  },

  changeFilter({commit}, {id, changes}) {

    commit('CHANGE_FILTER', {id, changes});

  },

  deleteFilter({commit}, id) {

    commit('DELETE_FILTER', id);

  },

  applyFilters({commit}) {

    commit('APPLY_FILTERS');

  }

}
