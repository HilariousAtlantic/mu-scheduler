const readline = require('readline');
const fs = require('fs');

let data = [];

let reader = readline.createInterface({
  input: fs.createReadStream('data.csv', {encoding: 'UTF-8'})
});

reader.on('line', line => data.push(line));
reader.on('close', parse);

function parse() {
  let sections = [];
  for (var i = 0; i < data.length; i++) {
    let parts = data[i].split(',');
    let section = {};
    section.id = parts[0];
    section.subject = parts[1];
    section.number = parts[2];
    section.section = parts[3];
    section.campus = parts[4];
    section.credits = parts[5];
    section.title = parts[6].replace(';', ',');
    section.meets = [{
      days: parts[7],
      time: parts[8],
      instructor: parts[9],
      date: parts[10],
      location: parts[11]
    }];
    while (data[i+1] && data[i+1].startsWith(',')) {
      parts = data[i++].split(',');
      section.meets.push({
        days: parts[7],
        time: parts[8],
        instructor: parts[9].replace(';', ','),
        date: parts[10],
        location: parts[11]
      });
    }
    sections.push(section);
  }
  let writer = fs.createWriteStream('sections.json');
  writer.write(JSON.stringify(sections));
}
