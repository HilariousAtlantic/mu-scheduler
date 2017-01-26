export default {

  selectedCourses(state) {

    return state.selectedCourses.map(id => state.detailedCourseCache[id]);

  },

  filteredSchedules(state) {

    let schedules = state.schedules;

    let activeFilters = state.filters.filter(({active}) => active);

    activeFilters.forEach(({test}) => {

      schedules = schedules.filter(test);

    });

    return schedules;

  }

}
