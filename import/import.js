const readline = require('readline');
const fs = require('fs');

const terms = require('./terms');

terms.forEach(term => {

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

      let meets = [{

        days: days,
        start_time: formatTime(time.split('-')[0]),
        end_time: formatTime(time.split('-')[1]),
        location: location,
        instructor: instructor.replace(';', ','),
        start_date: date.split('-')[0],
        end_date: date.split('-')[1]

      }];

      let tests = [];

      while (lines[i+1] && lines[i+1].startsWith(',')) {

        let [crn, subject, number, name, credits, title,
          days, time, instructor, date, location] = lines[++i].split(',');

        if (date.split('-')[0] == date.split('-')[1]) {

          tests.push({

            date: date.split('-')[0],
            start_time: formatTime(time.split('-')[0]),
            end_time: formatTime(time.split('-')[1]),
            location: location

          });

        } else {

          meets.push({

            days: days,
            start_time: formatTime(time.split('-')[0]),
            end_time: formatTime(time.split('-')[1]),
            location: location,
            instructor: instructor.replace(';', ','),
            start_date: date.split('-')[0],
            end_date: date.split('-')[1]

          });

        }

      }

      let course = courses.find(course => course.subject == subject && course.number == number);

      let section = {crn, name, tests, meets};

      if (course) {

        course.sections.push(section);

      } else {

        courses.push({semester_id: term.id, subject, number, title, credits, sections: [section]});

      }

    }

    let writer = fs.createWriteStream(`./json/${term.file}.json`);

    writer.write(JSON.stringify(courses));

  })

});

function formatTime(time) {

  if (!time) return;

  let parts = time.split(' ');

  let period = parts[1];
  let hours = parts[0].split(':')[0];
  let minutes = parts[0].split(':')[1];

  if (period == 'pm' && parseInt(hours) < 12) hours = parseInt(hours) + 12;

  return hours + ':' + minutes;

}
