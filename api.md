## API Documenation

GET /courses

```

  [
      {
       “id”: 0,
       “subject: “CSE”
       “number”: 201,
       “name”: “Intro to Software Engineering”
      },
      .
      .
      .
  ]

```

GET /courses/course_id

```

  {
   “id”: 0,
   “subject: “CSE”
   “number”: 201,
   “name”: “Intro to Software Engineering”
   “sections”: [
      {
       “id”: 323423,
       “course_id: “0”,
       “section”: “A”,
       “meets”: [
          {
          “days”: “MWF”,
          “start_time”: "14:30",
          “end_time”: "15:50,
          “instructor”: “Ann Sobel”,
          “location”: “BEN 009”,
          "start_date: "2016-08-29",
          "end_date": "2016-12-15"
          },
          .
          .
          .
       ]
      },
      .
      .
      .
   ]
  }

```

GET /sections

```

  [
      {
       “id”: 323423,
       “course_id: “0”,
       “section”: “A”,
       “meets”: [
          {
           “days”: “MWF”,
           “start_time”: 1600,
           “end_time”: 1900,
           “instructor”: “Ann Sobel”,
           “location”: “BEN 009”,
           "start_date: "2016-08-29",
           "end_date": "2016-12-15"
          },
          .
          .
          .
       ]
      },
      .
      .
      .
  ]

```


