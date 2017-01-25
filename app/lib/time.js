export function toTime(minutes, hidePeriod) {

  let h = Math.floor(minutes/60);
  let m = minutes - h*60;
  let p = 'AM'

  if (h >= 12) {

    p = 'PM';

  }

  if (h > 12) {

    h -= 12;

  }

  if (h === 0) {

    h = 12;

  }

  if (hidePeriod) {

    return h + ':' + ('00'+m).slice(-2);

  } else {

    return h + ':' + ('00'+m).slice(-2) + ' ' + p;

  }

}

export function formatTime(minutes, format) {

  let H = Math.floor(minutes/60);
  let h = H;
  let m = minutes - H*60;
  let p = 'AM'

  if (H >= 12) {

    p = 'PM';

  }

  if (H > 12) {

    h -= 12;

  }

  if (H === 0) {

    h = 12;

  }

  return format
    .replace(/HH/g, ('00'+H).slice(-2))
    .replace(/H/g, H)
    .replace(/hh/g, ('00'+h).slice(-2))
    .replace(/h/g, h)
    .replace(/mm/g, ('00'+m).slice(-2))
    .replace(/m/g, m)
    .replace(/P/g, p.toUpperCase())
    .replace(/p/g, p.toLowerCase())

}

export function toMinutes(time) {

  let [t, p] = time.toLowerCase().split(' ');
  let [h, m] = t.split(':');

  h = parseInt(h);
  m = parseInt(m);

  if (p === 'pm') {

    h += 12;

  }

  if (p === 'am' && h == 12) {

    h = 0;

  }

  return h*60 + m;

}

export function getTimes(step = 10) {

  let times = [];

  for (let i = 0; i < 1440; i += step) {

    times.push(toTime(i));

  }

  return times.sort();

}

export function getSuggestedTimes(step, term) {

  let trim = (term) => term.toLowerCase().replace(/\s+/, '');

  return getTimes(step).filter(time => trim(time).startsWith(trim(term)));


}
