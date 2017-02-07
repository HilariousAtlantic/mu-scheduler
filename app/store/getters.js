export default {

  selectedCourses(state) {

    return state.selectedCourses.map(id => state.detailedCourseCache[id]);

  },

  filteredSchedules(state) {

    let schedules = state.schedules;

    console.log('dsdf');

    state.filters.forEach(({active, test}) => {

      if (active) {

        schedules = schedules.filter(test);

      }

    });

    return schedules;

  }

}
