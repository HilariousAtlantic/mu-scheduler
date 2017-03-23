export function formatTime(minutes, format) {

  let H = Math.floor(minutes/60);
  let h = H == 0 ? 12 : (H > 12 ? H-12 : H);
  let m = minutes - H*60;
  let p = H >= 12 ? 'PM' : 'AM';

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

export function getMinutes(time) {

  let [t, h, m, p] = time.match(/(\d+):(\d+) (AM|PM)/);

  let hours = parseInt(h);
  let minutes = parseInt(m);

  switch (p) {

    case 'AM':

      return hours == 12 ? minutes : hours*60 + minutes;

    case 'PM':

      return hours*60 + minutes + (hours == 12 ? 0 : 720);

  }

}
