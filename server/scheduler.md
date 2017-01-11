//schedule = list of N different sections of different classes
//course = list of N different sections of the same class
//    ex. course CSE 148 = {section a, section b, section c...}
//courses = list of courses to be scheduled
//selectedSections = current list of sections we are considering 
//goodSchedules = list of schedules we have decided have no overlaps
goodSchedules := make([]Schedule,0)
findGoodSchedules(courses *[]Course, selectedSections *[]Section){
    //base case. If we have found the same # of sections
    //as classes we wanted to look through, then we have found
    //a good schedule.
    if len(selectedSections) == len(courses) {
        append(goodSchedules, selectedSections)
        return
    }
    for _, course := range courses{
        
        //check if we already selected this course. if we did, ignore 
        //the course entirely (skip to the next one)
        for _, selectedSection := range selectedSections{
            if selectedSection.courseID == course.ID{
                continue //goes to next course
            }
        }
        
        for _, section := range course{
            //check to make sure there is no overlap
            for _, selectedSection := range selectedSections{
                if !(section.startTime > selectedSection.startTime
                    && section.startTime < selectedSection.endTime || 
                    selectedSection.startTime > section.startTime
                    && selectedSection.startTime < section.startTime){
                        append(selectedSections, section)
                        findGoodSchedules(courses,selectedSections)
                    }
            }
            //this means we made it all the way through the courses
            //without finding a time that works (or we already searched
            //the ones that work, so we return (recurse back)
            return
        }
    }
        
}
