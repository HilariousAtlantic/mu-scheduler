const fs = require('fs');
const util = require('util');
const sleep = require('sleep');
const readline = require('readline');
const rq = require('request-promise-native');

let courses = [];
let coursesSet = new Set();

let allGPAs = [];

let semesters = {
  "10": "Fall",
  "15": "Winter",
  "20": "Spring",
  "30": "Summer"
};

importGrades();

function importGrades() {
  let reader = readline.createInterface({
    input: fs.createReadStream('./courses.csv', {encoding: 'UTF-8'})
  });

  reader.on('line', line => {
    let [section, term, crn, subject, number] = line.split(',');

    coursesSet.add(subject + number);
  });

  reader.on('close', () => {
    coursesSet.forEach(course => courses.push({
      subject: course.substr(0, 3),
      number: course.substr(3)
    }));

    fetchGPAs(courses.pop());
  });
}

function fetchGPAs(course) {
  let url = `http://grdist.miamioh.edu/php/getClasses.php?sem=&loc=O&inst=&from=2000&to=2016&iid=-1&did=-1&dept=${course.subject}&num=${course.number}`;

  console.log(`Fetching grades for ${course.subject + course.number} from ${url}`);

  rq({url: url, json: true})
    .then(GPAs  => {
      saveGPAs(GPAs);
      if (courses.length > 0) {
        fetchGPAs(courses.pop());
      }
    })
    .catch(err => {
      util.log(`Request for ${subject+number} grades failed. ${err}`);
    })
}

function saveGPAs(GPAs) {
  let writer = fs.createWriteStream('./grades.csv');

  GPAs.forEach(gpa => {
    if (gpa.campus != "O") return;

    writer.write([
      gpa.avggpa,
      fixProfessorName(gpa.name),
      gpa.NameShort,
      gpa.number,
      gpa.Year,
      semesters[gpa.Semester],
      gpa.Section,
    ].join(',') + '\n');
  });
}

/**
 * Format a professor name as `first last MI` * instead of `last first MI`
 */
function fixProfessorName(name) {
  [lastName, firstName, middleInitial] = name.trim().split(/\s+/);

  if (middleInitial) {
    return [firstName, lastName, middleInitial].join(' ');
  }

  return [firstName, lastName].join(' ');
}

/**
 * Fix the incorrectly formatted professor names that are already in the CSV file.
 */
function fixExistingProfessorNames() {
  let lines = fs.readFileSync('./grades.csv').toString().split('\n');

  lines = lines.map(line => {
    if (line.trim().length == 0) return '';

    let splitted = line.split(',');
    splitted[1] = fixProfessorName(splitted[1]);

    return splitted.join(',')
  });

  let fixed = lines.join('\n');

  fs.truncate('./grades.csv', 0, () => {
    fs.writeFile('./grades.csv', fixed, err => {
      if (err) {
        console.log(err);
      } else {
        console.log('Professor names fixed');
      }
    });
  });
}
