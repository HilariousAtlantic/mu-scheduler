export default {

    terms: [],
    courses: [],
    schedules: [],

    selectedTerm: {},
    selectedCourses: [],
    selectedSchedules: [],

    coursesFilter: '',
    schedulesIndex: 0,
    scheduleFilters: [

      {id: 0, type: 'time', options: {operator: 'Start After', time: 600, days: ['M', 'W', 'F']}, active: true},

      {id: 1, type: 'time', options: {operator: 'End Before', time: 1080, days: ['T', 'R']}, active: false},

      {id: 2, type: 'class', options: {operator: 'At Most', amount: 2, days: ['M']}, active: false},

      {id: 3, type: 'class', options: {operator: 'Exactly', amount: 3, days: ['T', 'R']}, active: true}

    ],

    coursesCache: {},
    detailedCourseCache: {},
    schedulesCache: {},

    requestingTerms: false,
    requestingCourses: false,
    requestingDetailedCourse: false,
    requestingSchedules: false

}
