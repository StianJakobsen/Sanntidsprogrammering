//Sanntidsprogrammering!!
package functions
import ("fmt" // Using '.' to avoid prefixing functions with their package names
		// This is probably not a good idea for large projects...
	//"runtime"
	//"time"
	//"net"
	//"bufio"
	//"os"
	//"strconv"
	"driver"
	//"sort"
	//"encoding/json"
	"sort"
	"time"
	"math"
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
	/*
	sort.Ints(DownList)
	if(len(DownList)>0){ 
		temp := make([]int,1)
		//fmt.Println("DownList i SortDown: ",DownList)
		temp[0] = DownList[len(DownList)-1]
		counter := 0
		for i:= (len(DownList)-1); i>=1; i-- {
			
			if DownList[i] < temp[counter] {
				counter ++
				temp = append(temp,DownList[i])
			}
		}
		return temp
	}else{
		return DownList
	}
	*/
} 
func Delay(SlaveTime time.Time, PrimeTime time.Time) int{
	temp := SlaveTime.Sub(PrimeTime)
	return int(math.Floor(temp.Seconds()))
}
func printData(data Data) {
	vektor1 := []string{"-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-"}
	//vektor2 := vektor1
	for j:=driver.N_FLOORS;j>-1;j--{	
		
		vektor1 = []string{"-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-"}
		for i:=0;i<len(data.PrimaryQ);i++{
			if(data.Statuses[i].CurrentFloor == j){
				vektor1[2+10*i]= "#"
			}
		}
		fmt.Print(j)
		fmt.Print(": ")
		fmt.Println(vektor1)
		
	//fmt.Println(vektor2)
	}



}
	

