// go run networkUDP.cd ..go
//Sanntidsprogrammering!!
package udp
import ("fmt" // Using '.' to avoid prefixing functions with their package names
		// This is probably not a good idea for large projects...
	//"runtime"
	"time"
	"net"
	//"bufio"
	"os"
	"strconv"
	"driver"
	//"sort"
	"encoding/json"
	//"sort"
	"functions"
	
)


type Status struct {
	Running int
	CurrentFloor int
	NextFloor int
	Primary bool
	ID int
	LastUpdate time.Time
	//PrimaryQ [3]string
	//CommandList []int
	UpList []int  // slice = slice[:0] for å tømme slicen når sendt til primary
	DownList[]int // slice = slice[:0] for å tømme slicen når sendt til primary
	ButtonList []int
	OrderList []int // sjekke for nye ordrer når primary sender
}

type Data struct {
	//Status Status
	//Timestamp???????
	ButtonList []int // [up0,up1,up2,dwn1,dwn2,dwn3]
	PriBroad bool
	ID int
	Statuses []Status // Oppdatere den her å i UdpInit()
	PrimaryQ []int
}


func SetStatus(status *Status, running int, NextFloor int) {
	(*status).LastUpdate = time.Now()
	(*status).Running = running
	(*status).CurrentFloor = driver.GetFloorSensorSignal()
	(*status).NextFloor = NextFloor
	(*status).ID = GetID()
	
	/*
	(*data).Statuses[GetIndex(GetID(), data)].Running = running
	(*data).Statuses[GetIndex(GetID(), data)].CurrentFloor = driver.GetFloorSensorSignal()
	(*data).Statuses[GetIndex(GetID(), data)].NextFloor = NextFloor
	(*data).Statuses[GetIndex(GetID(), data)].ID = ID
	//Println(" id i func:", (*Status).ID)
	*/
}
func GetID() int {
	ut:=0
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}
 	var ipAddr string
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipAddr = ipnet.IP.String()
			}
		}
	}
	if len(ipAddr)==14{ 
		ut,_ = strconv.Atoi(ipAddr[12:14])
	}else{
		ut,_ = strconv.Atoi(ipAddr[12:15])
	}
	return ut
	
}


/////////// Primary functions ////////////

func PrimaryBroadcast(baddr *net.UDPAddr, data *Data) { // IMALIVE, oppdatere backup for alle
	var temp Data
	
	//udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:" + strconv.Itoa(broadcastPort))
	//checkError(err)
	bconn, err := net.DialUDP("udp", nil, baddr)
	checkError(err)
	fmt.Println("BROADCASTER")
	for {
		temp = *data
		temp.PriBroad = true
		//fmt.Println("SENDER")
		// WRITE
		b,_ := json.Marshal(temp)
		bconn.Write(b)
		//json.Unmarshal(b[0:len(b)], temp) 
		//Println("b: ", b)
		//Println("PrimaryQ marshalled: ", len(temp.Statuses))
		checkError(err)
		time.Sleep(500*time.Millisecond)
	}

}

func SendOrderlist(data *Data,index int) { // IMALIVE
	//fmt.Println("Går inn og SENDER ORDRER!")
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:39998")//+strconv.Itoa(data.PrimaryQ[index])+":39998")
	bconn, err := net.DialUDP("udp",nil, udpAddr)
	checkError(err)
	// WRITE
	fmt.Print("ORDERLIST SENT: ", data.Statuses[index].OrderList)
	fmt.Println("                                  PrimeryQ: ", data.PrimaryQ)
	data.PriBroad = false
	b,_ := json.Marshal(data) // nok å bare sende en gang?
	bconn.Write(b)		
	checkError(err)
}

func PrimaryListen(in chan *Data, out chan *Data) {
	buffer := make([]byte, 1024)
	//var tempo Status
	//var data Data
	data := <-in
	out<-data
	//updating := false
	tempData := *data
	var receivedData Data
	tempData.ID = GetID()
	udpAddr, err := net.ResolveUDPAddr("udp", ":39999")
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {	
		//fmt.Println("før case data = <-in")
		select{
		case data = <-in:
			//fmt.Println("udp: 133. Inne i PrimaryListen")
			//updating = true
			if len(tempData.Statuses) > len(data.Statuses){
				fmt.Println("Legger til nye element i statuses")
				data.Statuses = append(data.Statuses, tempData.Statuses[len(data.Statuses):]...)
			}else if len(tempData.Statuses) < len(data.Statuses){
				fmt.Println("Fjerner element i statuses")
				for i := 1; i < len(data.Statuses); i++{
					if functions.CheckList(tempData.PrimaryQ, data.PrimaryQ[i]) == false{
						//functions.UpdateList(data.PrimaryQ, data.PrimaryQ[i])
						data.Statuses = UpdateStatusList(data.Statuses, GetIndex(data.PrimaryQ[i], data))
					}
				}
			}
			for i := 1; i < len(data.Statuses); i++{
				data.Statuses[i] = tempData.Statuses[i]
			}
			data.PrimaryQ = tempData.PrimaryQ
			data.ButtonList = tempData.ButtonList
			out <- data
		default:
			//fmt.Println("HØRER")
			//if len(tempData.PrimaryQ) == 1{
			//	conn.SetReadDeadline(time.Now().Add(500*time.Millisecond))
			//}
			n, err := conn.Read(buffer) // Høtt skjer om den stoppar her?
			//out<- data
			if err == nil{
				//fmt.Println("Mottok ei melding")
				//checkError(err)
				//Data = buffer
				err = json.Unmarshal(buffer[0:n], &receivedData)
				checkError(err)
				//fmt.Printf("Går inn i Checklisttingeling: %v, ID: %d\n", receivedData.PrimaryQ, receivedData.ID)
				//fmt.Println("Her er lengden til primaryQ",len(receivedData.Statuses))
				
				if functions.CheckList(tempData.PrimaryQ,receivedData.ID)==false {
					fmt.Printf("Går inn i Checklisttingeling: %v, ID: %d", receivedData.PrimaryQ, receivedData.ID)
					tempData.Statuses = append(tempData.Statuses, receivedData.Statuses[GetIndex(receivedData.ID, &receivedData)])
					tempData.PrimaryQ = append(tempData.PrimaryQ, receivedData.ID) //PrimaryQ[1:]...)
					tempData.ID = receivedData.ID
					SendOrderlist(&tempData,1)
				}else{
					
					fmt.Println("Recieved time: ", receivedData.Statuses[GetIndex(GetID(), &receivedData)].LastUpdate)
					tempData.ID = receivedData.ID
					tempData.Statuses[GetIndex(tempData.ID,&tempData)] = tempData.Statuses[GetIndex(tempData.ID, &tempData)]
					
					
				}
				
			}
			
			
			fmt.Println("PrimaryQ: ", tempData.PrimaryQ)
		
				fmt.Println("Delay: ", functions.Delay(tempData.Statuses[GetIndex(tempData.ID, &tempData)].LastUpdate,time.Now()))
				if(functions.Delay(time.Now(),tempData.Statuses[GetIndex(tempData.ID, &tempData)].LastUpdate)>5){
					tempData.Statuses[0].UpList = append(tempData.Statuses[0].UpList, tempData.Statuses[GetIndex(tempData.ID, &tempData)].UpList...)
					tempData.Statuses[0].DownList = append(tempData.Statuses[0].DownList, tempData.Statuses[GetIndex(tempData.ID, &tempData)].DownList...)
					for j:=0;j<6;j++{
						if(tempData.ButtonList[j] == 1 && j<3) {
							tempData.Statuses[0].UpList = append(tempData.Statuses[0].UpList,j) 
							
						}else if(tempData.ButtonList[j] == 1){
							tempData.Statuses[0].DownList = append(tempData.Statuses[0].DownList,j-2)
						}
					}
					tempData.Statuses = UpdateStatusList(tempData.Statuses,GetIndex(tempData.ID,&tempData))
					tempData.PrimaryQ = functions.UpdateList(tempData.PrimaryQ,GetIndex(tempData.ID,&tempData))
					SendOrderlist(&tempData,1)
				}
			
			

		}
		
	}
}
/*
func CleanDeadSlaves(data Data){ // FIX
	for{
		data.Statuses[0].LastUpdate = time.Now()
		time.Sleep(3*time.Second)
		fmt.Println("Lendgen til primaryq i cleandeadslaves: ",len(data.PrimaryQ))
		for i:=1;i<len(data.PrimaryQ);i++{
			fmt.Println("Sjekker delay: ",functions.Delay(data.Statuses[0].LastUpdate,data.Statuses[GetIndex(data.PrimaryQ[i],data)].LastUpdate))
			if(functions.Delay(data.Statuses[0].LastUpdate,data.Statuses[GetIndex(data.PrimaryQ[i],data)].LastUpdate)>2){
				data.Statuses = UpdateStatusList(data.Statuses,GetIndex(data.PrimaryQ[i],data))
				data.PrimaryQ = functions.UpdateList(data.PrimaryQ,i)
				
			}			
		}		
	}
}
*/
/////////// Slave functions //////////// 

func ListenForPrimary(bconn *net.UDPConn,baddr *net.UDPAddr ,in chan *Data, out chan *Data, PrimaryChan chan int) { // Bruke chan muligens fordi den skal skrive til Data
	buffer := make([]byte, 1024)
	//udpAddr, err := net.ResolveUDPAddr("udp", ":" + strconv.Itoa(broadcastPort))
	//bconn, err := net.ListenUDP("udp", udpAddr)
	//checkError(err)
	var data *Data
	var temp Data
	for {
		data = <-in 
		
		fmt.Println("Hører")
		//fmt.Println("Her er gammel OrderList: ", data.Statuses[GetIndex(GetID(),data)].OrderList)	
		bconn.SetReadDeadline(time.Now().Add(5*time.Second))		
		n, err := bconn.Read(buffer)
		//fmt.Println("Størrelsen på mottatt data: ", n)
		if err != nil && data.PrimaryQ[1] == GetID() {
			fmt.Println("Mottar ikke meldinger fra primary lenger, tar over")
			//fiks primary sin orderlist
			data.Statuses[1].UpList = append(data.Statuses[1].UpList,data.Statuses[0].UpList...)
			data.Statuses[1].DownList = append(data.Statuses[1].DownList,data.Statuses[0].DownList...)
			for j:=0;j<6;j++{
				if(data.ButtonList[j] == 1 && j<3) {
					data.Statuses[1].UpList = append(data.Statuses[1].UpList,j) 
				}else if(data.ButtonList[j] == 1){
					data.Statuses[1].DownList = append(data.Statuses[1].DownList,j-2)
				}
			}
			
			data.PrimaryQ = data.PrimaryQ[1:] // UpdateList(data.PrimaryQ,0)
			data.Statuses = data.Statuses[1:]
			go PrimaryBroadcast(baddr, data)
			//go PrimaryListen(data, SortChan)
			// SendOrderlist(Data)
			//PrimaryChan <-
			go ChannelFunc(PrimaryChan)
			break
		}
		//Data = buffer
		err = json.Unmarshal(buffer[0:n], &temp)
		fmt.Println("PrimaryQ: ", temp.PrimaryQ)
		
		if(temp.PriBroad == false) {
			if(len(data.PrimaryQ)!=len(temp.PrimaryQ)){
				*data = temp
			}else{
				for i := 0; i < len(temp.Statuses); i++{
					if temp.PrimaryQ[i] != GetID(){
						data.Statuses[i] = temp.Statuses[i]
					}
				}
				data.Statuses[GetIndex(GetID(), &temp)].UpList = append(data.Statuses[GetIndex(GetID(), &temp)].UpList, temp.Statuses[GetIndex(GetID(), &temp)].UpList...)
				data.Statuses[GetIndex(GetID(), &temp)].DownList = append(data.Statuses[GetIndex(GetID(), &temp)].DownList, temp.Statuses[GetIndex(GetID(), &temp)].DownList...)
				data.Statuses[GetIndex(GetID(), &temp)].ButtonList = temp.Statuses[GetIndex(GetID(), &temp)].ButtonList
				data.Statuses[GetIndex(GetID(), &temp)].Running = temp.Statuses[GetIndex(GetID(), &temp)].Running
				data.Statuses[GetIndex(GetID(), &temp)].OrderList = append(data.Statuses[GetIndex(GetID(), &temp)].OrderList, temp.Statuses[GetIndex(GetID(), &temp)].OrderList...)
					
			}
			
			
			fmt.Println("her er primaryQen:", data.PrimaryQ)
			fmt.Println("Her er lengden til primaryQ",len(data.Statuses))
			fmt.Println("Her er ny OrderList: ", data.Statuses[GetIndex(GetID(),data)].OrderList)
			fmt.Println("Index: ", GetIndex(GetID(),data))
		}	
		out <- data
		
		
		// Printf("Rcv %d bytes: %s\n",n, buffer)	
	}	
}

/*
func SlaveAlive(data *Data) {
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187."+ strconv.Itoa((*data).PrimaryQ[0]) + ":"+strconv.Itoa(GetID()+30000)
	conn, err := net.DialUDP("udp",nil, udpAddr)
	checkError(err)
	for {
		 //WRITE
		(*data).ID = GetID()
		fmt.Println("Data.ID før sending",(*data).ID)
		
		
		b,_ := json.Marshal(*data)
		// Må endre detta til å bare slette når confirmation på ordre kommer, confirmation kan vere samma som lampe lista??
		(*data).Statuses[GetIndex(GetID(), data)].UpList = (*data).Statuses[GetIndex(GetID(), data)].UpList[:0]
		(*data).Statuses[GetIndex(GetID(),data)].DownList = (*data).Statuses[GetIndex(GetID(), data)].DownList[:0]
		
		conn.Write(b)	
		checkError(err)
		time.Sleep(150*time.Millisecond) // bytte til bare ved endringar etterhvert
		if (*data).Statuses[GetIndex(GetID(), data)].Primary == true {
			break
		}
	}
}
*/
func SlaveUpdate(in chan *Data, out chan *Data, timeChan chan time.Time) { // chan muligens, bare oppdatere når det er endringar
	data := <-in
	out<- data
	//init := 0
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:39999")
	conn, err := net.DialUDP("udp",nil, udpAddr)
	checkError(err)
	
	for {
		 //WRITE
		data = <-in
		data.ID = GetID()
		data.Statuses[GetIndex(data.ID,data)].LastUpdate = <-timeChan
		fmt.Println("Data.ID før sending",data.ID)
		
		if(driver.GetFloorSensorSignal()!=-1){	
			data.Statuses[GetIndex(GetID(), data)].CurrentFloor = driver.GetFloorSensorSignal()
		}
		//if len(data.Statuses[GetIndex(data.ID,data)].UpList)>0||len(data.Statuses[GetIndex(data.ID,data)].DownList)>0||init==0 {
		b,_ := json.Marshal(*data)
		//init = 1
		fmt.Println("Sender denne UpList: ", data.Statuses[GetIndex(data.ID,data)].UpList)
		fmt.Println("Sender denne DownList: ", data.Statuses[GetIndex(data.ID,data)].DownList)
		// Må endre detta til å bare slette når confirmation på ordre kommer, confirmation kan vere samma som lampe lista??
		//data.Statuses[GetIndex(GetID(), data)].UpList = data.Statuses[GetIndex(GetID(), data)].UpList[:0]
		//data.Statuses[GetIndex(GetID(), data)].DownList = data.Statuses[GetIndex(GetID(), data)].DownList[:0]
		
		conn.Write(b)	
		checkError(err)
		time.Sleep(150*time.Millisecond) // bytte til bare ved endringar etterhvert
		if data.Statuses[GetIndex(GetID(), data)].Primary == true {
			break
		//}
		}
		out<- data
	}
}

// send_ch, receive_ch chan Udp_message
func UdpInit(localListenPort int, broadcastListenPort int, message_size int, data *Data, slaveListenIn chan *Data, slaveListenOut chan *Data, PrimaryChan chan int, SlaveChan chan int) (err error) {
	
	buffer := make([]byte, message_size)
	var status Status
	//data.Statuses = append(data.Statuses, temp)
	status.Primary = false
	//(*data).ID = GetID()	
	SetStatus(&status,0,driver.GetFloorSensorSignal())
	data.PriBroad = false
	//InitStatus(*Status)
	//Println("SE HER::::: ", (Status).ID)
	
	//Generating broadcast address
	baddr, err = net.ResolveUDPAddr("udp4", "129.241.187.255:"+strconv.Itoa(broadcastListenPort))
	if err != nil {
		return err
	}

	//Generating localaddress
	tempConn, err := net.DialUDP("udp4", nil, baddr)
	defer tempConn.Close()
	tempAddr := tempConn.LocalAddr()
	laddr, err = net.ResolveUDPAddr("udp4", tempAddr.String())
	laddr.Port = localListenPort

	//Creating local listening connections
	localListenConn, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		return err
	}

	//Creating listener on broadcast connection
	broadcastListenConn, err := net.ListenUDP("udp", baddr)
	if err != nil {
		localListenConn.Close()
		return err
	}

	//go udp_receive_server(localListenConn, broadcastListenConn, message_size receive_ch)
	//go udp_transmit_server(localListenConn, broadcastListenConn ,send_ch)

	//Setting first primary
	broadcastListenConn.SetReadDeadline(time.Now().Add(3*time.Second))
	n, err := broadcastListenConn.Read(buffer)
	if err != nil {
		data.ButtonList = []int{0,0,0,0,0,0}
		fmt.Println("Tar over som primary!")
		data.PrimaryQ = append(data.PrimaryQ, GetID())
		data.Statuses = append(data.Statuses, status)
		data.Statuses[GetIndex(GetID(), data)].Primary = true
		//go ChannelFunc(PrimaryChan)
		go PrimaryBroadcast(baddr,data)
		
	} else {
		err = json.Unmarshal(buffer[0:n], data)
		fmt.Println("PrimaryQ før checklist: ", data.PrimaryQ)
		fmt.Println("Checklist i control 349: ",functions.CheckList(data.PrimaryQ,GetID()) == false)
		if functions.CheckList(data.PrimaryQ,GetID()) == false{
			fmt.Println("Funkar checklist?")
			data.PrimaryQ = append(data.PrimaryQ, GetID())
			data.Statuses = append(data.Statuses, status)
		}
		
		//(*data).Statuses = temp.Statuses
		
		//(*data).PrimaryQ[1:] = SortUp((*data).PrimaryQ[1:])
		fmt.Println("PrimaryQ: ", data.PrimaryQ)
		fmt.Println("Statuselen: ", len(data.Statuses))
		//(*Data).PrimaryQ = append((*Data).PrimaryQ, string(buffer))
		//SlaveChan<- 1
		go ChannelFunc(SlaveChan)
		//go SlaveUpdate(slaveListenIn , slaveListenOut )				
		time.Sleep(2500*time.Millisecond) // Vente for å la Primary oppdatere PrimaryQen
		go ListenForPrimary(broadcastListenConn,baddr, slaveListenIn, slaveListenOut, PrimaryChan)
		
		
	}
	


	//	fmt.Printf("Generating local address: \t Network(): %s \t String(): %s \n", laddr.Network(), laddr.String())
	//	fmt.Printf("Generating broadcast address: \t Network(): %s \t String(): %s \n", baddr.Network(), baddr.String())
	return err
}

func GetIndex(ID int, data *Data) int { 
	for i:=0; i<len(data.PrimaryQ); i++ {
		if data.PrimaryQ[i] == ID {
			return i
		}
	}
	return -1
}


func checkError(err error) {
	if err != nil {
		fmt.Println("Noe gikk galt %v", err)
		return
	}
}

func ChannelFunc(Channel chan int) {
	Channel <-1
}

func UpdateStatusList(OrderList []Status, j int) []Status {
	temp := make([]Status, len(OrderList)-1)
	for i:= 0; i<len(OrderList);i++ {
		if i<j {
			temp[i] = OrderList[i]
		} else if i>j {
			temp[i-1] = OrderList[i]
		}
	}
	return temp
}

func PrintData(data Data) {
	vektor1 := []string{"-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-","-"}
	//vektor2 := vektor1
	for j:=driver.N_FLOORS-1;j>-1;j--{	
		
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

/*
func SendCommandList() { // Bare sende siste tal for simplicity
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:30169") // Broadcast (endre ip nettverket du sitter på)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	currentStruct := TellerStruct{teller}

	for {
		b,_ := json.Marshal(currentStruct)
		conn.Write(b)	
		Println("S
		
		
		ent: ",currentStruct.Teller) 		
		currentStruct.Teller = currentStruct.Teller + 1
		time.Sleep(1*time.Second)
	}
}*/





/*
func SendCommand(floorChan chan int) {
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:30169") // Broadcast (endre ip nettverket du sitter på)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	currentStruct := TellerStruct{teller}

	for {
		b,_ := json.Marshal(currentStruct)
		conn.Write(b)	
		Println("Sent: ",currentStruct.Teller) 		
		currentStruct.Teller = currentStruct.Teller + 1
		time.Sleep(1*time.Second)
	}

}*/




