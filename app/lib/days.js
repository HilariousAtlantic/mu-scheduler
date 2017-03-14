export function formatDay(day) {

  switch (day.toUpperCase()) {

    case 'M':

      return 'Monday';

    case 'T':

      return 'Tuesday';

    case 'W':

      return 'Wednesday';

    case 'R':

      return 'Thursday';

    case 'F':

      return 'Friday';

  }

}

export function formatDayList(days) {

  let order = 'MTWRF';

  days = days.sort((a, b) => order.indexOf(a)-order.indexOf(b));

  if (days.length === 0) {

    return 'None';

  } else if (days.length === 1) {

    return formatDay(days[0]);

  } else if (days.length === 2) {

    return formatDay(days[0]) + ' and ' + formatDay(days[1]);

  } else {

    if ('MTWRFS'.indexOf(days.join('')) != -1) {

      return formatDay(days[0]) + ' through ' + formatDay(days[days.length-1]);

    } else {

      return days.slice(0, days.length-1).map(formatDay).join(', ') + ', and ' + formatDay(days[days.length-1]);

    }

  }

}
