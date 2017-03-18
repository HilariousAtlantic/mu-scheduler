const readline = require('readline');
const fs = require('fs');

const terms = require('./terms.json');

let writer = fs.createWriteStream('courses.csv');

terms.forEach(term => {

  console.log('Formatting ' + term.name + '...');

  let lines = [];

  let reader = readline.createInterface({
    input: fs.createReadStream(`./courses/${term.file}.csv`, {encoding: 'UTF-8'})
  });

  reader.on('line', function(line) {

    if (line != ',,,,,,,,,,' && !line.startsWith('CRN')) {

      lines.push(line);

    }

  })

  reader.on('close', function() {

    let currentSection = {};

    lines.map(function(line) {

      let [crn, subject, number, name, credits, title, days, time, instructor, date, location, attribute] = line.split(',');

      let type = 'section';
      let [start_time, end_time] = formatTime(time);
      let [start_date, end_date] = date.split('-');

      instructor = instructor.split(';')[0].replace(' (P)', '');
      let names = instructor.split(' ');
      instructor = names[0] + ' ' + names[names.length-1];
      if (instructor == 'TBA TBA') instructor = 'TBA';

      if (!crn) {

        type = start_date == end_date ? 'test' : 'meet';

        ({crn, subject, number, title, name, credits, attribute} = currentSection);

      } else {

        currentSection = {crn, subject, number, title, name, credits, attribute};

      }

      return [type, term.name, crn, subject, number, title, name, credits, days, start_time, end_time, location, instructor, start_date, end_date, attribute].join(',').replace(/,,/g, ',TBA,');s

    }).forEach(line => writer.write(line + '\n'));

  })

})

function formatTime(time) {

  if (!time || time == 'TBA') return ['-1', '-1'];

  let [start, end] = time.split('-');

  if (!end) console.log(time);

  let [startTime, startPeriod] = start.split(' ');
  let [endTime, endPeriod] = end.split(' ');

  let [startHours, startMinutes] = startTime.split(':');
  let [endHours, endMinutes] = endTime.split(':');

  startHours = parseInt(startHours);
  startMinutes = parseInt(startMinutes);
  endHours = parseInt(endHours);
  endMinutes = parseInt(endMinutes);

  if (startPeriod == 'pm' && startHours < 12) startHours += 12;
  if (endPeriod == 'pm' && endHours < 12) endHours += 12;

  return [startHours*60+startMinutes, endHours*60+endMinutes];

}
