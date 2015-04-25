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
	
	//udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:39998")
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
	data.PriBroad = false
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:39998")//+strconv.Itoa(data.PrimaryQ[index])+":39998")
	bconn, err := net.DialUDP("udp",nil, udpAddr)
	checkError(err)
	// WRITE
	fmt.Print("ORDERLIST SENT: ", data.Statuses[index].OrderList)
	fmt.Println("                                  PrimeryQ: ", data.PrimaryQ)
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
	//var tempData Data
	tempData := *data
	var temp Data
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
			for i := 1; i < len(data.Statuses); i++{
				data.Statuses[i] = tempData.Statuses[i]
			}
			if len(tempData.Statuses) > len(data.Statuses){
				fmt.Println("Legger til nye element i statuses")
				data.Statuses = append(data.Statuses, tempData.Statuses[len(data.Statuses):]...)
			}
			data.PrimaryQ = tempData.PrimaryQ
			out <- data
		default:
			//fmt.Println("HØRER")
			if len(tempData.PrimaryQ) == 1{
				conn.SetReadDeadline(time.Now().Add(500*time.Millisecond))
			}
			n, err := conn.Read(buffer) // Høtt skjer om den stoppar her?
			//out<- data
			if err == nil{
				//fmt.Println("Mottok ei melding")
				//checkError(err)
				//Data = buffer
				err = json.Unmarshal(buffer[0:n], &temp)
				if functions.CheckList(tempData.PrimaryQ,temp.ID)==false {
					tempData.Statuses = append(tempData.Statuses, temp.Statuses[GetIndex(temp.ID, &temp)])
					tempData.PrimaryQ = append(tempData.PrimaryQ, temp.PrimaryQ[len(temp.Statuses)-1]) //PrimaryQ[1:]...)
				}else{
					tempData.Statuses[GetIndex(temp.ID, data)] = temp.Statuses[GetIndex(temp.ID, data)]
				}
			}
			//if update == true{
			//	out <- data
			//	update = false
			//}
			//(*data).Statuses[GetIndex((*data).Status.ID,data)] = (*data).Status // Oppdaterar mottatt status hos primary 
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

func ListenForPrimary(bconn *net.UDPConn, baddr *net.UDPAddr, in chan *Data, out chan *Data, PrimaryChan chan int) { // Bruke chan muligens fordi den skal skrive til Data
	buffer := make([]byte, 1024)
	//udpAddr, err := net.ResolveUDPAddr("udp", ":39998")
	//conn, err := net.ListenUDP("udp", udpAddr)
	//checkError(err)
	var data *Data
	var temp Data
	for {
		data = <-in 
		
		fmt.Println("Hører")
		fmt.Println("Her er gammel OrderList: ", data.Statuses[GetIndex(GetID(),data)].OrderList)	
		bconn.SetReadDeadline(time.Now().Add(5*time.Second))		
		n, err := bconn.Read(buffer)
		if err != nil && data.PrimaryQ[1] == GetID() {
			fmt.Println("Mottar ikke meldinger fra primary lenger, tar over")
			data.PrimaryQ = data.PrimaryQ[1:] // UpdateList(data.PrimaryQ,0)
			data.Statuses = data.Statuses[1:]
			//go PrimaryBroadcast(baddr, data)
			//go PrimaryListen(data, SortChan)
			// SendOrderlist(Data)
			PrimaryChan<- 1
			break
		}
		//Data = buffer
		err = json.Unmarshal(buffer[0:n], &temp)
		if(temp.PriBroad == false) {
			*data = temp	
			fmt.Println("her er primaryQen:", data.PrimaryQ)
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
func SlaveUpdate(in chan *Data, out chan *Data) { // chan muligens, bare oppdatere når det er endringar
	data := <-in
	out<- data
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187."+ strconv.Itoa(data.PrimaryQ[0]) + ":39999")
	conn, err := net.DialUDP("udp",nil, udpAddr)
	checkError(err)
	for {
		 //WRITE
		data = <-in
		data.ID = GetID()
		fmt.Println("Data.ID før sending",data.ID)
		data.Statuses[GetIndex(GetID(), data)].LastUpdate = time.Now()
		
		b,_ := json.Marshal(*data)
		// Må endre detta til å bare slette når confirmation på ordre kommer, confirmation kan vere samma som lampe lista??
		//data.Statuses[GetIndex(GetID(), data)].UpList = data.Statuses[GetIndex(GetID(), data)].UpList[:0]
		//data.Statuses[GetIndex(GetID(), data)].DownList = data.Statuses[GetIndex(GetID(), data)].DownList[:0]
		
		conn.Write(b)	
		checkError(err)
		time.Sleep(50*time.Millisecond) // bytte til bare ved endringar etterhvert
		if data.Statuses[GetIndex(GetID(), data)].Primary == true {
			break
		}
		out<- data
	}
}

// send_ch, receive_ch chan Udp_message
func UdpInit(localListenPort int, broadcastListenPort int, message_size int, data *Data, slaveListenIn chan *Data, slaveListenOut chan *Data, PrimaryChan chan int, SlaveChan chan int, ) (err error) {
	
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
		fmt.Println("Tar over som primary!")
		data.PrimaryQ = append(data.PrimaryQ, GetID())
		data.Statuses = append(data.Statuses, status)
		data.Statuses[GetIndex(GetID(), data)].Primary = true
		//PrimaryChan <- 1
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
		//SlaveUpdate(data)				
		go ListenForPrimary(broadcastListenConn, baddr, slaveListenIn, slaveListenOut, PrimaryChan)
		time.Sleep(2500*time.Millisecond) // Vente for å la Primary oppdatere PrimaryQen
		
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




