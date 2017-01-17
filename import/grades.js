const fs = require('fs');
const util = require('util');
const readline = require('readline');
const rq = require('request-promise-native');

//const http://grdist.miamioh.edu/php/getClasses.php?dept=%v&num=%v&inst=&from=%v&to=%v&iid=-1&did=%v&sem=&loc=O

let writer = fs.createWriteStream('./grades.csv');

let reader = readline.createInterface({
  input: fs.createReadStream('./courses.csv', {encoding: 'UTF-8'})
});

let allGrades = [];

let courses = new Set();
let subjects = new Map();


reader.on('line', line => {
  let [section, term, crn, subject, number] = line.split(',');

  if (courses.has(subject + number)) return;
  courses.add(subject + number);

  fetchDepartmentId(subject)
    .then(id => console.log(id))
    //.then(id => fetchGPAs)
    //.then(() => [...allGrades, ...grades])
    .catch((err) => {
      util.log(err);
    });
});

function fetchGPAs(subject, number) {

}

function fetchDepartmentId(subject) {
  return new Promise((resolve, reject) => {
    if (subjects.has(subject)) {
       resolve(subjects.get(subject));
    } else {
      rq({url: `http://grdist.miamioh.edu/php/getDID.php?dept=${subject}`, json: true})
        .then(body => JSON.parse(body))
        .then(data => {
          if (data.length > 0 && data[0].did > -1) {
            subjects.set(subject, data[0].did);
            resolve(data);
          } else {
            subjects.set(subject, null);
            reject(`Failed to fetch department ID for ${subject}.  Response: ${data}`);
          }
        })
        .catch(err => {
          reject(`Request for ${subject} department ID failed. ${err}`);
        });
    }
  });
}
