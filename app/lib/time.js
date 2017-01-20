export function toTime(minutes) {

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

  return h + ':' + ('00'+m).slice(-2) + ' ' + p;

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
