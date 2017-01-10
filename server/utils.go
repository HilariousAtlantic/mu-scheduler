package main

import(
	"fmt"
	"strconv"
)



//09:30 am-11:00 am
func getStartTime( unparsed string){
	[]string spDash = SplitN(unparsed,"-",2)
	fmt.println(spDash[0])
	[]string spColon = SplitN(spDash[0],":",2)
	fmt.println(spColon)
	hours := strconv.ParseInt(spColon[0])
	[]string startTime := SplitN(spColon[1]," ", 2)
	minutes := strconv(startTime[0])

	

}
