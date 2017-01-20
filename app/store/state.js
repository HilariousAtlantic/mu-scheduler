export default {

    terms: [],
    courses: [],
    schedules: [],

    selectedTerm: {},
    selectedCourses: [],
    selectedSchedules: [],

    courseFilter: '',
    scheduleFilters: [

      {id: 0, type: 'time', options: {operator: 'Start After', time: 600, days: ['M', 'W', 'F']}, active: false, test: () => true},

      {id: 1, type: 'time', options: {operator: 'Finish Before', time: 1080, days: ['T', 'R']}, active: false, test: () => true},

      {id: 2, type: 'class', options: {operator: 'At Most', amount: 2, days: ['M']}, active: false, test: () => true},

      {id: 3, type: 'class', options: {operator: 'Exactly', amount: 3, days: ['T', 'R']}, active: false, test: () => true}

    ],

    coursesCache: {},
    detailedCourseCache: {},
    schedulesCache: {},

    requestingTerms: false,
    requestingCourses: false,
    requestingDetailedCourse: false,
    requestingSchedules: false

}
