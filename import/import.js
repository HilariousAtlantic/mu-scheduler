const readline = require('readline');
const fs = require('fs');

let data = [];

let reader = readline.createInterface({
  input: fs.createReadStream(process.argv[2] || 'sections.csv', {encoding: 'UTF-8'})
});

reader.on('line', line => {

  if (line != ',,,,,,,,,,' && !line.startsWith('CRN')) {
    data.push(line)
  }

});

reader.on('close', parse);

function parse() {

  let courses = [];

  for (var i = 0; i < data.length; i++) {

    let parts = data[i].split(',');

    let section = {

      crn: parts[0],
      subject: parts[1],
      number: parts[2],
      name: parts[3],
      campus: parts[4],
      credits: parts[5],
      title: parts[6].replace(';', ','),
      tests: [],
      meets: [{

        days: parts[7],
        start_time: standardizeTime(parts[8].split('-')[0]),
        end_time: standardizeTime(parts[8].split('-')[1]),
        instructor: parts[9].replace(';', ','),
        start_date: parts[10].split('-')[0],
        end_date: parts[10].split('-')[1],
        location: parts[11]

      }]

    };

    while (data[i+1] && data[i+1].startsWith(',')) {
      parts = data[i++].split(',');
      if (parts[10].split('-')[0] == parts[10].split('-')[1]) {

        section.tests.push({
          start_time: standardizeTime(parts[8].split('-')[0]),
          end_time: standardizeTime(parts[8].split('-')[1]),
          date: parts[10].split('-')[0],
          location: parts[11]
        });

      } else {

        section.meets.push({
          days: parts[7],
          start_time: standardizeTime(parts[8].split('-')[0]),
          end_time: standardizeTime(parts[8].split('-')[1]),
          instructor: parts[9].replace(';', ','),
          start_date: parts[10].split('-')[0],
          end_date: parts[10].split('-')[1],
          location: parts[11]
        });

      }

    }

    let {subject, number, title, credits, crn, name, meets, campus} = section;
    let course = courses.find(course => course.subject == subject && course.number == number);

    if (course) {

      course.sections.push({crn, section, meets, campus})

    } else {

      courses.push({subject, number, title, credits, sections: [{crn, name, meets, campus}]})

    }
  }

  let writer = fs.createWriteStream(process.argv[3] || 'courses.json');
  writer.write(JSON.stringify(courses));
}

function standardizeTime(time) {

  if (!time) return;

  let parts = time.split(' ');

  let period = parts[1];
  let hours = parts[0].split(':')[0];
  let minutes = parts[0].split(':')[1];

  if (period == 'pm' && parseInt(hours) < 12) hours = parseInt(hours) + 12;

  return hours + ':' + minutes;

}
