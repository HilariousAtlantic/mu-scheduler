const readline = require('readline');
const fs = require('fs');

const terms = require('./terms.json');

let writer = fs.createWriteStream('courses.csv');

terms.forEach(term => {

  let lines = [];

  let reader = readline.createInterface({
    input: fs.createReadStream(`./csv/${term.file}.csv`, {encoding: 'UTF-8'})
  });

  reader.on('line', function(line) {

    if (line != ',,,,,,,,,,' && !line.startsWith('CRN')) {

      lines.push(line);

    }

  })

  reader.on('close', function() {

    let currentSection = {};

    lines.map(function(line) {

      let [crn, subject, number, name, credits, title, days, time, instructor, date, location] = line.split(',');

      let type = 'section';
      let [start_time, end_time] = formatTime(time);
      let [start_date, end_date] = date.split('-');

      if (!crn) {

        type = start_date == end_date ? 'test' : 'meet';

        ({crn, subject, number, title, name, credits} = currentSection);

      } else {

        currentSection = {crn, subject, number, title, name, credits};

      }

      return [type, term.name, crn, subject, number, title, name, credits, days, start_time, end_time, location, instructor, start_date, end_date].join(',')

    }).forEach(line => writer.write(line + '\n'));

  })

})

function formatTime(time) {

  if (!time || time == 'TBA') return ['', ''];

  let [start, end] = time.split('-');

  let [startTime, startPeriod] = start.split(' ');
  let [endTime, endPeriod] = start.split(' ');

  let [startHours, startMinutes] = startTime.split(':');
  let [endHours, endMinutes] = endTime.split(':');

  if (startPeriod == 'pm' && parseInt(startHours) < 12) startHours = parseInt(startHours) + 12;
  if (endPeriod == 'pm' && parseInt(endHours) < 12) endHours = parseInt(endHours) + 12;

  return [startHours + ':' + startMinutes, endHours + ':' + endMinutes];

}
