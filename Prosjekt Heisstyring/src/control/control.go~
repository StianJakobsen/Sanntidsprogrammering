//Sanntidsprogrammering!!
package control

import ( 
	"fmt"
	"udp"
	//"runtime"
	"time"
	"driver"
	//"control"
	//"os"
	"functions"
	
	
)
/*
func GoToFloor(button int,  floorChan chan int,data *udp.Data) {
	floor := <-floorChan
	if driver.GetFloorSensorSignal() == -1 {
		driver.SetMotorDirection(driver.DIRN_DOWN)
	}
	var done int
	temp:= floor	
	//polse:
	for {		
		select {
		
		case temp = <-floorChan:
			//fmt.Println("Her er temp: %d", temp)
			//fmt.Println("Her er DONE: %d", done)
			//if done == 1{
									
			//	floor = temp
			//	done = 0
			//}

		default:			

			driver.SetFloorIndicator(driver.GetFloorSensorSignal())	
			if done == 1{
				//fmt.Printf("GAA IN EHFE")				
				floor = temp
				done = 0
				
				
			}	
			//fmt.Printf("Hva er done? %d\n",done)
			driver.SetButtonLamp(button,floor,1)
			//fmt.Printf("Her er flooooooooor: %d\n", floor)
				
			if floor == driver.GetFloorSensorSignal()  {
				
				fmt.Println("Framme på:", floor)
				udp.SetStatus(data,0,floor)

				driver.SetDoorOpenLamp(true)				
				driver.SetMotorDirection(driver.DIRN_STOP)
				time.Sleep(1*time.Second)
				driver.SetDoorOpenLamp(false)
				driver.SetFloorIndicator(floor)
				driver.SetButtonLamp(button,floor,0)
				done = 1
				
				//temp = -1
				//driver.SetDoorOpenLamp(false)	
				//fmt.Println("Done: %d", done)
				break
		
			} else if floor > driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1 {   
			udp.SetStatus(data,2, floor)
			driver.SetMotorDirection(driver.DIRN_UP) 
		
			} else if floor < driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1{
			udp.SetStatus(data ,1, floor)
			driver.SetMotorDirection(driver.DIRN_DOWN)
			}

		}
	
	}
}
*/
func GoToFloor(floor int, data *udp.Data) { // Lamper for command buttons må leggas til, kall list noe annet
	//fmt.Println("control 82: går til floor floor:",floor)
	//fmt.Println("control 82: er i floor:",status.CurrentFloor)
	for { 
		driver.SetFloorIndicator(driver.GetFloorSensorSignal())
		if floor == driver.GetFloorSensorSignal() {
				driver.SetFloorIndicator(floor)
				driver.SetButtonLamp(2,floor,0)
				driver.SetMotorDirection(driver.DIRN_STOP)
				driver.SetDoorOpenLamp(true)				
				time.Sleep(1500*time.Millisecond)
				driver.SetDoorOpenLamp(false)
				
				data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor = floor
				if (floor == 0 || floor == 3) || len(data.Statuses[udp.GetIndex(udp.GetID(), data)].OrderList) == 0 {
					data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = 0
				}
				//if list == 1 {
				//	driver.SetButtonLamp((*status).ButtonList[0], floor, 0)
				//	(*status).ButtonList = functions.UpdateList((*status).ButtonList,0)
				//}
				fmt.Println("Heisen er framme på floor:", floor)
				udp.PrintData(*data)
				break
		} else if floor > driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1 {   
			driver.SetMotorDirection(driver.DIRN_UP)
			data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = 1
		} else if floor < driver.GetFloorSensorSignal() && driver.GetFloorSensorSignal() != -1 && floor != -1{
			driver.SetMotorDirection(driver.DIRN_DOWN)
			data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = -1
		}/*else if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor == -1{
			
			driver.SetMotorDirection(driver.DIRN_DOWN)
		}*/
		if driver.GetFloorSensorSignal() != -1{
			data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor = driver.GetFloorSensorSignal()
			driver.SetButtonLamp(2,data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor,0)
		}	
	}
}


func ElevatorControl(data *udp.Data) {
	//time.Sleep(1*time.Second)
	//var data.Statuses[udp.GetIndex(udp.GetID(),data)] *udp.data.Statuses[udp.GetIndex(udp.GetID(),data)]
	temp := 0
	//temp = temp + 0
	
	for {
		//&&data.Statuses[udp.GetIndex(udp.GetID(),data)] = <-data.Statuses[udp.GetIndex(udp.GetID(),data)]In
	
		//if driver.GetFloorSensorSignal() != -1 {
		//	data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor = driver.GetFloorSensorSignal()
		//}
		//fmt.Println(data.data.Statuses[udp.GetIndex(udp.GetID(),data)]es[udp.GetIndex(data.PrimaryQ[i], data)].CurrentFloor) 
		//fmt.Println("control 109: OrderList",data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList)
		//if len(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList)==0 {
		//	data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, -1)
		//}
		//if len(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList) == 0{
		//	data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = 0
		//	data.Statuses[udp.GetIndex(udp.GetID(),data)]Out<-data.Statuses[udp.GetIndex(udp.GetID(),data)]
		//} 
		//for i:=0;i<len(data.Statuses[udp.GetIndex(udp.GetID(),data)].ButtonList);i++ {
		//	fmt.Println("ButtonList(i): ",data.Statuses[udp.GetIndex(udp.GetID(),data)].ButtonList[i])
		//	fmt.Println("OrderList(i): ",data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[i])
		//	driver.SetButtonLamp(data.Statuses[udp.GetIndex(udp.GetID(),data)].ButtonList[i], data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[i], 1)
		//}
		//ButtonList = ButtonList[:0]
						
		//fmt.Printf("OrderList: %d CommandList[0]: CurrentFloor: %d ID: %d \n",data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor, data.Statuses[udp.GetIndex(udp.GetID(),data)].ID)
		
		if len(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList) > 0{
				fmt.Println("OrderList: ", data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList)
				// 
				if data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[0] != data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor  {
					//fmt.Println(data.data.Statuses[udp.GetIndex(udp.GetID(),data)]es[udp.GetIndex(data.PrimaryQ[i], data)].OrderList)
					// Sjekker om heisens ordreliste
					temp = data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[0]
					data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = functions.UpdateList(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList,0)
					GoToFloor(temp,data) // vurdere å kjøre commandbuttons inni gotofloor
					temp = 0
				/*
				}else if data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[0] < data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor{
					temp = data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[0]
					data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = functions.UpdateList(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList,0)
					GoToFloor(temp,data)
					temp = 0
				*/		
				}else if data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList[0] == driver.GetFloorSensorSignal() {
					data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList=functions.UpdateList(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList,0)
					GoToFloor(driver.GetFloorSensorSignal(), data)
				}

		}	
	}
}
	
		
func GetDestination(data *udp.Data) { //returnerer bare button, orderlist oppdateres
	//time.Sleep(1*time.Second)
	for {
		time.Sleep(2*time.Millisecond) // Polling rate, mby change	
		for floor := 0; floor < driver.N_FLOORS; floor++ {
				if driver.GetButtonSignal(0,floor) == 1 && len(data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList) == 0 {
					data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList, floor)
					fmt.Println("control: 250, data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList: ", data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList) 
				}else if driver.GetButtonSignal(0,floor) == 1 && len(data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList) > 0 {
					if functions.CheckList(data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList,floor) == false {
						data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList,floor)
						fmt.Println("control: 254, data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList: ", data.Statuses[udp.GetIndex(udp.GetID(),data)].UpList) 
					}				
				}else if driver.GetButtonSignal(1,floor) == 1 && len(data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList)==0 {	
					data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList, floor)
					fmt.Println("control: 257, data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList: ", data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList)
				}else if driver.GetButtonSignal(1,floor) == 1 && len(data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList) > 0 {
					if functions.CheckList(data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList,floor) == false {
						data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList,floor)
						fmt.Println("control: 260, data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList: ", data.Statuses[udp.GetIndex(udp.GetID(),data)].DownList)
					}	
				}else if driver.GetButtonSignal(2,floor) == 1 && !functions.CheckList(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor){
					//if data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 0{
					//	
					//}
					if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor < floor && data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 1{
						fmt.Println("er")
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor)
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = functions.SortUp(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList)
						driver.SetButtonLamp(2,floor,1)
					} else if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor > floor && data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == -1{
						fmt.Println("er eg")
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor)
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = functions.SortDown(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList)
						driver.SetButtonLamp(2,floor,1)
					} else if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor < floor && data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 0{
						fmt.Println("er eg her")
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor)
						data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = 1
						driver.SetButtonLamp(2,floor,1)
					} else if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor > floor && data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 0{
						fmt.Println("er eg her?")
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor)
						data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = -1
						driver.SetButtonLamp(2,floor,1)
					} else if data.Statuses[udp.GetIndex(udp.GetID(),data)].CurrentFloor == floor && data.Statuses[udp.GetIndex(udp.GetID(),data)].Running == 0{
						fmt.Println("er eg her!!!!!")
						data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList = append(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor)
						data.Statuses[udp.GetIndex(udp.GetID(),data)].Running = 0
						driver.SetButtonLamp(2,floor,1)
					}
				} else if driver.GetButtonSignal(2,floor) == 1 && functions.CheckList(data.Statuses[udp.GetIndex(udp.GetID(),data)].OrderList, floor) {
					driver.SetButtonLamp(2,floor,1)
				}
					
		}
		/*
		if(driver.GetStopSignal() != 0) {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
		*/
	}
}
/*
func GetCommand() (int,int) {
	button := 2	
	for {	
		for floor := 0; floor < driver.N_FLOORS; floor++ {
			if(driver.GetButtonSignal(button,floor) == 1) {
				return button,floor
			}
			
		}
		if(driver.GetStopSignal() != 0) {
			driver.SetMotorDirection(driver.DIRN_STOP)
			break
		}
	}
return -1,-1
}
*/

func CostFunction(in chan *udp.Data, out chan *udp.Data) {
	handled := 0
	var DownList []int
	var UpList []int
	var data *udp.Data
	data = <-in
	for {
		//fmt.Println("control 243, handled: ",handled)
		handled = 0
		//fmt.Println("status.UpList i CostFunction: ",(*data).Statuses[udp.GetIndex((*data).PrimaryQ[0], data)].UpList)
		//fmt.Println("Lengden til statuses: ", len(data.Statuses))
		//fmt.Println("PrimaryQ: ", data.PrimaryQ)
		/*if len(data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList) > 0 {
		if(data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList[0] == -1){
			data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList = UpdateList(data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList, 0)
		}
		}*/
		if len(UpList) == 0 && len(DownList) == 0{
			if handled == 0{
				out<- data
				data = <-in
				//fmt.Println("control: 383. Er inne i costfunction")
			}else{
				out<-data // Tømt UpList og DownList 
				data = <-in // Venter
			}
		}else if len(UpList) > 0 || len(DownList) > 0{
			fmt.Printf("Could not handle all orders, will try again after recieving new.\nUpList: %v DownList: %v\n", UpList, DownList)
			out<-data
			data = <-in
		}
		for k := 0; k < len(data.PrimaryQ);k++ {
			if udp.GetIndex(data.PrimaryQ[k],data) != -1 {
				//fmt.Printf("PrimaryQ: %v Lengde Statuses: %v\n", data.PrimaryQ, len(data.Statuses))
				//fmt.Println("Har den samme info om status.Uplist her: ", data.Statuses[k].UpList)
				DownList = append(DownList,data.Statuses[k].DownList...)
				//data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].DownList = data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].DownList[:0]

				UpList = append(UpList,data.Statuses[k].UpList...)
				//data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].UpList = data.Statuses[udp.GetIndex(data.PrimaryQ[k], data)].UpList[:0]
			}
		}

	//if len(UpList) > 0 || len(DownList) > 0{
	//	//fmt.Println("status.UpList i CostFunction: ",data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].UpList)
	//	fmt.Println("Lengden til statuses: ", len(data.Statuses))
	//	fmt.Println("PrimaryQ: ", data.PrimaryQ)
	//	fmt.Println("control 258: OppList i cost function: ", UpList)
	//	fmt.Println("control 259: Down List i cost function: ", DownList)
	//	time.Sleep(2*time.Second)
	//}
	
	if len(UpList) > 0 {
		UpList = functions.SortUp(UpList)
	}
	if len(DownList) > 0 {
		DownList = functions.SortDown(DownList)
	}
	
	//fmt.Println("OrderList: ", data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList)
	/*if(len(data.PrimaryQ) > 1){
		data.Statuses[1].OrderList = append(data.Statuses[1].OrderList,3)
		
		udp.SendOrderlist(data,1)
	}*/
	//fmt.Println("Sjekk om UPLIST oppdateres riktig: ", UpList)
	//fmt.Println("Sjekk om DOWNLIST oppdateres riktig: ", DownList)
	//fmt.Println(DownList)
	for down:=0; down<len(DownList);down++ { // Kanskje feil å sette running inni her? 
		
		if handled == 1{
			handled = 0
			down = 0	
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i samme floor, og står stille
			if DownList[down] == data.Statuses[i].CurrentFloor && data.Statuses[i].Running == 0 {
				
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,DownList[down])
				data.Statuses[i].OrderList = functions.SortDown(data.Statuses[i].OrderList)
				data.Statuses[i].Running = 0
				fmt.Println("control 280: Heis i samma floor og står stille. Downlist:", DownList)
				DownList = functions.UpdateList(DownList,down) //Må modifiseres
				data.Statuses[i].DownList = functions.UpdateList(data.Statuses[i].DownList, down) 
				data.ButtonList[2+DownList[down]] =1
				if i != 0 {
					udp.SendOrderlist(data,i) // , i)
				}else{
					
				}
				handled = 1
				break
				//pluss noe mer, som å åpne døra
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i etasjen over og på veg nedover
			if data.Statuses[i].CurrentFloor == DownList[down]+1 && data.Statuses[i].Running == -1 && handled != 1 {
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,DownList[down])
				fmt.Println("control 370: Heis i etasjen over og på vei nedover. Downlist:", DownList)
				data.Statuses[i].OrderList = functions.SortDown(data.Statuses[i].OrderList)
				DownList = functions.UpdateList(DownList,down)
				data.Statuses[i].DownList = functions.UpdateList(data.Statuses[i].DownList, down)
				data.Statuses[i].ButtonList = append(data.Statuses[i].ButtonList,1)
				data.ButtonList[2+DownList[down]] =1
				if i != 0 {
					udp.SendOrderlist(data, i) // , i)
				}
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis i etasjen over og står stille
			if data.Statuses[i].CurrentFloor == DownList[down]+1 && data.Statuses[i].Running == 0 && handled != 1 {
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,DownList[down])
				data.Statuses[i].OrderList = functions.SortDown(data.Statuses[i].OrderList)
				data.Statuses[i].Running = -1
				fmt.Println("control 385: Heis i etasjen over og står stille")
				DownList = functions.UpdateList(DownList,down)
				data.Statuses[i].DownList = functions.UpdateList(data.Statuses[i].DownList, down)
				data.Statuses[i].ButtonList = append(data.Statuses[i].ButtonList,1)
				data.ButtonList[2+DownList[down]] =1
				if i != 0 {
					udp.SendOrderlist(data,i) // , i)
				}
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ { // Heis på veg nedover
			if data.Statuses[i].CurrentFloor > DownList[down] && data.Statuses[i].Running == -1  && handled != 1{
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,DownList[down])
				DownList = functions.UpdateList(DownList,down)
				data.Statuses[i].DownList = functions.UpdateList(data.Statuses[i].DownList, down)
				fmt.Println("control 398: Heis på vei nedover og er over floor. Downlist:", DownList)
				data.Statuses[i].OrderList = functions.SortDown(data.Statuses[i].OrderList)
				data.Statuses[i].ButtonList = append(data.Statuses[i].ButtonList,1)
				data.ButtonList[2+DownList[down]] =1
				if i != 0 {
					udp.SendOrderlist(data,i) // , i)
				}
				handled = 1
				break 
			}
		}

		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[i].Running == 0 {
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,DownList[down])
				fmt.Println("control 437: heis står stille generelt. Downlist:", DownList)
				data.ButtonList[2+DownList[down]] =1
				if DownList[down] > data.Statuses[i].CurrentFloor{
					data.Statuses[i].OrderList = functions.SortUp(data.Statuses[i].OrderList)
					data.Statuses[i].Running = 1
				}else if DownList[down] < data.Statuses[i].CurrentFloor{
					data.Statuses[i].OrderList = functions.SortDown(data.Statuses[i].OrderList)
					data.Statuses[i].Running = -1
				}
				data.Statuses[i].ButtonList = append(data.Statuses[i].ButtonList,1)
				DownList = functions.UpdateList(DownList,down)
				data.Statuses[i].DownList = functions.UpdateList(data.Statuses[i].DownList, down)
				if i != 0 {
					udp.SendOrderlist(data,i) // , i)
				}
				handled = 1
				break 
			}
		}

		/*if handled == 0 {
			data.Statuses[data.PrimaryQ[0]].OrderList = append(data.Statuses[data.PrimaryQ[0]].OrderList,DownList[down])
			data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList = SortDown(data.Statuses[udp.GetIndex(data.PrimaryQ[0], data)].OrderList)
			DownList = UpdateList(DownList,down)
			fmt.Println("h")
			handled = 1 
		}*/
	}

for up:=0; up<len(UpList);up++ {
		//fmt.Println("Up: ",up)
		//fmt.Println(data.PrimaryQ)
		if handled == 1{
			handled = 0
			up = 0	
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if UpList[up] == data.Statuses[i].CurrentFloor && data.Statuses[i].Running == 0 {
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,UpList[up])
				data.Statuses[i].OrderList = functions.SortUp(data.Statuses[i].OrderList)
				data.Statuses[i].NextFloor = UpList[up]
				fmt.Println("control 387: heis i samme etasjen og står stille. UpList:", UpList)
				data.Statuses[i].ButtonList = append(data.Statuses[i].ButtonList,0)
				UpList = functions.UpdateList(UpList,up) //Må modifiseres
				data.Statuses[i].UpList = functions.UpdateList(data.Statuses[i].UpList, up)
				data.
				if i != 0 {
					udp.SendOrderlist(data,i) // , i)
				}
				handled = 1
				break 
				//pluss noe mer, som å åpne døra
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[i].CurrentFloor == UpList[up]-1 && data.Statuses[i].Running == 1 && handled != 1 {
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,UpList[up])
				data.Statuses[i].OrderList = functions.SortUp(data.Statuses[i].OrderList)
				data.Statuses[i].NextFloor = UpList[up]
				fmt.Println("control 402: heis i etasjen under og på vei oppover. UpList:", UpList)
				data.Statuses[i].ButtonList = append(data.Statuses[i].ButtonList,0)
				UpList = functions.UpdateList(UpList,up)
				data.Statuses[i].UpList = functions.UpdateList(data.Statuses[i].UpList, up)
				if i != 0 {
					udp.SendOrderlist(data,i) // , i)
				}
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[i].CurrentFloor == UpList[up]-1 && data.Statuses[i].Running == 0 && handled != 1 {
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,UpList[up])
				data.Statuses[i].OrderList = functions.SortUp(data.Statuses[i].OrderList)
				data.Statuses[i].NextFloor = UpList[up]
				data.Statuses[i].Running = 1
				fmt.Println("control 417: heis i etasjen under og står stille. UpList:", UpList)
				UpList = functions.UpdateList(UpList,up)
				data.Statuses[i].UpList = functions.UpdateList(data.Statuses[i].UpList, up)
				data.Statuses[i].ButtonList = append(data.Statuses[i].ButtonList,0)
				if i != 0 {
					udp.SendOrderlist(data,i) // , i)
				}
				handled = 1
				break 
			}
		}
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			if data.Statuses[i].CurrentFloor < UpList[up] && data.Statuses[i].Running == 1  && handled != 1{
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,UpList[up])
				data.Statuses[i].OrderList = functions.SortUp(data.Statuses[i].OrderList)
				fmt.Println("control 430: floor over heis.currentfloor og på vei oppover. UpList:", UpList)
				data.Statuses[i].NextFloor = UpList[up]
				UpList = functions.UpdateList(UpList,up)
				data.Statuses[i].UpList = functions.UpdateList(data.Statuses[i].UpList, up)
				data.Statuses[i].ButtonList = append(data.Statuses[i].ButtonList,0)
				if i != 0 {
					udp.SendOrderlist(data,i) // , i)
				}
				handled = 1
				break 
			}
		}

		fmt.Println("Her er handled: ",handled)		
		for i := 0; i < len(data.PrimaryQ) && handled == 0; i++ {
			fmt.Println("RUNNING RUNNING RUNNING",data.Statuses[i].Running  )
			if data.Statuses[i].Running == 0 {
				data.Statuses[i].OrderList = append(data.Statuses[i].OrderList,UpList[up])
				
				fmt.Println("control 473: heisen står stille. UpList:", UpList)
				
				data.Statuses[i].Running = 1
				//fmt.Println("Sjekker GetIndex: ", i)
				//fmt.Println("Sjekker Statuses: ", len(data.Statuses))
				//fmt.Println("Sjekker CurrentFloor: ", data.Statuses[i].CurrentFloor)
				//fmt.Println("Sjekker UpList: ", UpList)
				if UpList[up] > data.Statuses[i].CurrentFloor{
					data.Statuses[i].OrderList = functions.SortUp(data.Statuses[i].OrderList)
					data.Statuses[i].Running = 1
				}else if UpList[up] < data.Statuses[i].CurrentFloor{
					data.Statuses[i].OrderList = functions.SortDown(data.Statuses[i].OrderList)
					data.Statuses[i].Running = -1
				}
				UpList = functions.UpdateList(UpList,up)
				data.Statuses[i].UpList = functions.UpdateList(data.Statuses[i].UpList, up)
				data.Statuses[i].ButtonList = append(data.Statuses[i].ButtonList,0)
				if i != 0 {
					udp.SendOrderlist(data,i) // , i)
				}
				
				handled = 1
				break 
			}
		}
		/*
		if handled == 0 {
			data.Statuses[data.PrimaryQ[0]].OrderList = append(data.Statuses[data.PrimaryQ[0]].OrderList,UpList[up])
			UpList = UpdateList(UpList,up)
			handled = 1
			fmt.Println("p") 
		}*/
	}
	}
}









