## API Documenation

GET /semesters

```
[
  {
    "id": 0,
    "season": "spring",
    "year": 2017,
    "name": "Spring 2017"
  },
  .
  .
  .
]
```

GET /courses

```
[
  {
    "credits": "3",
    “id”: 0,
    "semester_id": 0,
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
  "credits": "3",
  “id”: 0,
  “sections”: [
    {
      “id”: 323423,
      “course_id: “0”,
      “section”: “A”,
      “meets”: [
        {
          “days”: “MWF”,
          “start_time”: 60,
          “end_time”: 150,
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
  ],
  "semester_id": 0,
  “subject: “CSE”
  “number”: 201,
  “name”: “Intro to Software Engineering”
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
        “start_time”: 60,
        “end_time”: 150,
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
