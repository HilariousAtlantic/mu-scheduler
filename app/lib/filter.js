const defaultOptions = {

  time: {operator: 'Start After', time: 600, days: ['M', 'W', 'F']},

  class: {operator: 'Exactly', amount: 3, days: ['T', 'R']},

  break: {start: 720, finish: 780, days: ['T', 'W', 'R']}

}

export function getDefaultOptions(type) {

  return defaultOptions[type];

}

export function getFilter(type, options) {

  switch (type) {

    case 'time':

      return getTimeFilter(options);

    case 'class':

      return getClassFilter(options);

    case 'break':

      return getBreakFilter(options);

  }

}

export function getTimeFilter({operator, time, days}) {

  return function({startTimes, endTimes}) {

    for (let day of days) {

      switch (operator.toLowerCase()) {

        case 'start before':

          if (startTimes[day] > time) return false;

          break;

        case 'start after':

          if (startTimes[day] < time) return false;

          break;

        case 'finish before':

          if (endTimes[day] > time) return false;

          break;

        case 'finish after':

          if (endTimes[day] < time) return false;

          break;

      }

    }

    return true;

  }

}

export function getClassFilter({operator, amount, days}) {

  return function({classLoads}) {

    for (let day of days) {

      switch (operator.toLowerCase()) {

        case 'at least':

          if (classLoads[day] < amount) return false;

          break;

        case 'exactly':

          if (classLoads[day] != amount) return false;

          break;

        case 'at most':

          if (classLoads[day] > amount) return false;

          break;

      }

    }

    return true;

  }

}

export function getBreakFilter({start, finish, days}) {

  return function({timesByDay}) {

    for (let day of days) {

      let times = timesByDay[day];

      for (let time of times) {

        console.log(time, start, finish);

        if (time.start < finish && time.end > start) {

          return false;

        }

      }

    }

    return true;

  }

}
