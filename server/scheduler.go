package main


//without comments
//goodSchedules := make([[]Section]Course,0)
findGoodSchedules(courses *[]Course, selectedSections *[]Section){

    if len(selectedSections) == len(courses) {
        append(goodSchedules, selectedSections)
        return
    }
    for _, course := range courses{

        for _, selectedSection := range selectedSections{
            if selectedSection.courseID == course.ID{
                continue
            }
        }
        for _, section := range course{
            for _, selectedSection := range selectedSections{

                if !(section.startTime > selectedSection.startTime
                    && section.startTime < selectedSection.endTime ||
                    selectedSection.startTime > section.startTime
                    && selectedSection.startTime < section.startTime){
                        append(selectedSections, section)
                        findGoodSchedules(courses,selectedSections)
                    }
            }
            return
        }
    }
}
