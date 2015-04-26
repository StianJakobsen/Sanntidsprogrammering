//Sanntidsprogrammering!!
package functions
import (
	"time"
	"math"
	"sort"
)




func UpdateList(OrderList []int, j int) []int {
	temp := make([]int, len(OrderList)-1)
	for i:= 0; i<len(OrderList);i++ {
		if i<j {
			temp[i] = OrderList[i]
		} else if i>j {
			temp[i-1] = OrderList[i]
		}
	}
	return temp
}





func SortUp(UpList []int)  []int{ //Sorterer listen UpList i stigende rekkefølge og fjerner like tall og -1
	sort.Ints(UpList)
	temp := make([]int,1)
	temp[0] = UpList[0]
	
	
	counter := 0
	for i:= 1;i<len(UpList); i++ {
		if UpList[i] > temp[counter] {
			counter ++
			temp = append(temp,UpList[i])
		}
	}
	return temp
}
func CheckList(list []int, check int) bool{ // Sjekker om listen list inneholder heltallet check
	for i:=0;i<len(list);i++{
		if(list[i] == check){
			return true
		}
	}
	return false
}



func SortDown(DownList []int)  []int{
	DownList = SortUp(DownList)
	sort.Sort(sort.Reverse(sort.IntSlice(DownList)))
	return DownList
	
} 
func Delay(SlaveTime time.Time, PrimeTime time.Time) int{
	temp := SlaveTime.Sub(PrimeTime)
	return int(math.Floor(temp.Seconds()))
}

	

