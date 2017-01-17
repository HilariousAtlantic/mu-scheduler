export default {

  selectedCourses(state) {

    return state.selectedCourses.map(id => state.detailedCourseCache[id]);

  }

}
