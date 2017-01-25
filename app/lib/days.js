export const Days = {

  M: 'Monday',
  T: 'Tuesday',
  W: 'Wednesday',
  R: 'Thursday',
  F: 'Friday'

}

export function formatDay(day) {

  return Days[day];

}

export function formatDayList(days) {

  if (days.length === 0) {

    return 'no day';

  } else if (days.length === 6) {

    return 'each day';

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
