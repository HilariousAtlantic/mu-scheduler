export default {

  selectedCourses(state) {

    return state.selectedCourses.map(id => state.detailedCourseCache[id]);

  },

  filteredSchedules(state) {

    let schedules = state.schedules;

    let filters = state.scheduleFilters;

  }

}
