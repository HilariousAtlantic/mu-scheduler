export default {

  selectedCourses(state) {

    return state.selectedCourses.map(id => state.detailedCourseCache[id]);

  },

  filteredSchedules(state) {

    let schedules = state.schedules;

    state.filters.forEach(({active, test}) => {

      if (active) {

        schedules = schedules.filter(test);

      }

    });

    return schedules;

  }

}
