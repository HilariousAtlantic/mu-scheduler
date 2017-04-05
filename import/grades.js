#!/usr/bin/env node

const fs = require('fs');
const util = require('util');
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
  let url = `http://10.36.0.48/php/getClasses.php?sem=&loc=O&inst=&from=2000&to=2016&iid=-1&did=-1&dept=${course.subject}&num=${course.number}`;

  process.stdout.write(`Fetching grades for ${course.subject + course.number}... `);

  rq({url: url, json: true})
    .then(GPAs => {
      process.stdout.write("Done\n");
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
  let writer = fs.createWriteStream('./grades.csv', { flags: 'a' });

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

function fixProfessorName(name) {
  [lastName, firstName, middleInitial] = name.trim().split(/\s+/);

  return [firstName, lastName].join(' ');
}

