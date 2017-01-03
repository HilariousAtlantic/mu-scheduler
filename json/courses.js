const readline = require('readline');
const fs = require('fs');

let data = [];

let reader = readline.createInterface({
  input: fs.createReadStream('data.csv', {encoding: 'UTF-8'})
});

reader.on('line', line => data.push(line));
reader.on('close', parse);

function parse() {
  let courses = [];
  for (var i = 0; i < data.length; i++) {
    if (!data[i].startsWith(',')) {
      let parts = data[i].split(',');
      let course = {};
      course.subject = parts[1];
      course.number = parts[2];
      course.title = parts[6].replace('+', ',');
      if (!duplicate(courses, course)) {
        courses.push(course);
      }
    }
  }
  let writer = fs.createWriteStream('courses.json');
  writer.write(JSON.stringify(courses));
}

function duplicate(courses, course) {
  return courses.filter(({subject, number}) => subject == course.subject && number == course.number).length > 0;
}
