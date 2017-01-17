export default {

  selectedCourses(state) {

    return state.selectedCourses.map(id => state.detailedCourseCache[id]);

  },

  generatedSchedules(state, getters) {

    function getSection(id) {

      for (let course of getters.selectedCourses) {

        let found = course.sections.find(section => section.id === id);

        if (found) {

          return {meets: found.meets, name: course.subject + ' ' + course.number + ' ' + found.name}

        }

      }

    }

    return state.schedules.map(schedule => ({courses: schedule.sections.map(getSection)}));

  }

}
