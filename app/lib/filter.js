const defaultOptions = {

  start: {operator: 'After', time: 600, days: ['M', 'W', 'F']},

  finish: {operator: 'Before', time: 960, days: ['M', 'W', 'F']},

  break: {start: 660, finish: 780, days: ['T', 'W', 'R']},

  class: {operator: 'At Most', amount: 3, days: ['T', 'R']}

}

export function getDefaultOptions(type) {

  return defaultOptions[type];

}

export function getFilter(type, options) {

  switch (type) {

    case 'start':

      return getStartFilter(options);

    case 'finish':

      return getFinishFilter(options);

    case 'break':

      return getBreakFilter(options);

    case 'class':

      return getClassFilter(options);

  }

}

export function getStartFilter({operator, time, days}) {

  return function({startTimes}) {

    for (let day of days) {

      switch (operator.toLowerCase()) {

        case 'before':

          if (startTimes[day] >= time) return false;

          break;

        case 'after':

          if (startTimes[day] <= time) return false;

          break;

        case 'exactly':

          if (startTimes[day] != time) return false;

          break;

      }

    }

    return true;

  }

}

export function getFinishFilter({operator, time, days}) {

  return function({endTimes}) {

    for (let day of days) {

      switch (operator.toLowerCase()) {

        case 'before':

          if (endTimes[day] >= time) return false;

          break;

        case 'after':

          if (endTimes[day] <= time) return false;

          break;

        case 'exactly':

          if (endTimes[day] != time) return false;

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

        if (time.start < finish && time.end > start) {

          return false;

        }

      }

    }

    return true;

  }

}
