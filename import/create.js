const readline = require('readline');
const fs = require('fs');

const terms = require('./terms.json');

terms.forEach(term => {

  console.log('Creating json for ' + term.name);

  let lines = [];

  let reader = readline.createInterface({
    input: fs.createReadStream(`./csv/${term.file}.csv`, {encoding: 'UTF-8'})
  });

  reader.on('line', line => {
    if (line != ',,,,,,,,,,' && !line.startsWith('CRN')) lines.push(line)
  });

  reader.on('close', function() {

    let courses = [];

    for (var i = 0; i < lines.length; i++) {

      let [crn, subject, number, name, credits, title,
        days, time, instructor, date, location] = lines[i].split(',');

      title = title.replace(';', ',');
      instructor = instructor.replace(';', ',');
      let {start_time, end_time} = formatTime(time);
      let [start_date, end_date] = date.split('-');

      let meets = [{days, start_time, end_time, location, instructor, start_date, end_date}];

      let tests = [];

      while (lines[i+1] && lines[i+1].startsWith(',')) {

        let [crn, subject, number, name, credits, title,
          days, time, instructor, date, location] = lines[++i].split(',');

        instructor = instructor.replace(';', ',');
        let {start_time, end_time} = formatTime(time);
        let [start_date, end_date] = date.split('-');

        if (start_date == end_date) {

          tests.push({start_time, end_time, location,

            date: start_date

          });

        } else {

          meets.push({days, start_time, end_time, location, instructor, start_date, end_date});

        }

      }

      let course = courses.find(course => course.subject == subject && course.number == number);

      let section = {crn, name, tests, meets};

      if (course) {

        course.sections.push(section);

      } else {

        courses.push({term_id: term.id, subject, number, title, credits, sections: [section]});

      }

    }

    let writer = fs.createWriteStream(`./json/${term.file}.json`);

    writer.write(JSON.stringify(courses));

  })

});

function formatTime(time) {

  if (!time || time == 'TBA') return {start_time: '', end_time: ''};

  let [start, end] = time.split('-');

  let [startTime, startPeriod] = start.split(' ');
  let [endTime, endPeriod] = start.split(' ');

  let [startHours, startMinutes] = startTime.split(':');
  let [endHours, endMinutes] = endTime.split(':');

  if (startPeriod == 'pm' && parseInt(startHours) < 12) startHours = parseInt(startHours) + 12;
  if (endPeriod == 'pm' && parseInt(endHours) < 12) endHours = parseInt(endHours) + 12;

  return {start_time: startHours + ':' + startMinutes, end_time: endHours + ':' + endMinutes};

}
