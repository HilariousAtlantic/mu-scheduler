package main

//without comments
//goodSchedules := make([[]Section]Course,0)
/*
func findGoodSchedules(courses *[]Course, selectedSections *[]Section) {

	if len(*selectedSections) == len(*courses) {
		GoodSchedules = append(GoodSchedules, *selectedSections)
		return
	}
SKIPCOURSE:
	for _, course := range *courses {

		for _, selectedSection := range *selectedSections {
			if selectedSection.CourseID == course.ID {
				continue SKIPCOURSE
			}
		}
		for _, section := range getSectionsFromCourse(*course) {
			for _, selectedSection := range *selectedSections {

				if !(section.startTime > selectedSection.startTime &&
					section.startTime < selectedSection.endTime ||
					selectedSection.startTime > section.startTime &&
						selectedSection.startTime < section.startTime) {
					*selectedSections = append(*selectedSections, section)
					findGoodSchedules(courses, selectedSections)
				}
			}
			return
		}
	}
}
*/
func getCourseTree(ids string) []*Course {

	courses := getCoursesFromIDString(ids)
	sections := getSectionsFromCourses(courses)
	meets := getMeetsFromSections(sections)

	for _, section := range sections {
		if len(meets) == 0 {
			break
		}
		for _, meet := range meets {
			if meet.SectionID == section.ID {
				section.Meets = append(section.Meets, *meet)
			}

		}
	}
	for _, course := range courses {
		for _, section := range sections {
			if section.CourseID == course.ID {
				course.Sections = append(course.Sections, *section)
			}
		}
	}
	return courses
}
