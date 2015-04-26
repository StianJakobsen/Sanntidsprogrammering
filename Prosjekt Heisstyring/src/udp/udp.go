// go run networkUDP.cd ..go
//Sanntidsprogrammering!!
package udp
import ("fmt" 
	"time"
	"net"
	"os"
	"strconv"
	"driver"
	"encoding/json"
	"functions"
	
)


type Status struct {
	Running int
	CurrentFloor int
	Primary bool
	ID int
	UpList []int 
	DownList[]int 
	ButtonList []int
	OrderList []int
}

type Data struct {
	ButtonList []int // [up0,up1,up2,dwn1,dwn2,dwn3]
	PriBroad bool
	ID int
	Statuses []Status
	PrimaryQ []int
}


func SetStatus(status *Status, running int, NextFloor int) {
	(*status).Running = running
	(*status).CurrentFloor = driver.GetFloorSensorSignal()
	(*status).ID = GetID()
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
	bconn, err := net.DialUDP("udp", nil, baddr)
	checkError(err)
	fmt.Println("BROADCASTER")
	for {
		temp = *data
		temp.PriBroad = true
		// WRITE
		b,_ := json.Marshal(temp)
		bconn.Write(b)
		checkError(err)
		time.Sleep(500*time.Millisecond)
	}
}

func SendOrderlist(data *Data,index int) { 
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:39998")
	bconn, err := net.DialUDP("udp",nil, udpAddr)
	checkError(err)
	// WRITE
	fmt.Print("ORDERLIST SENT: ", data.Statuses[index].OrderList)
	fmt.Println("                                  PrimeryQ: ", data.PrimaryQ)
	data.PriBroad = false
	b,_ := json.Marshal(data) 
	bconn.Write(b)		
	checkError(err)
}

func PrimaryListen(in chan *Data, out chan *Data) {
	buffer := make([]byte, 1024)
	data := <-in
	out<-data
	tempData := *data
	var receivedData Data
	tempData.ID = GetID()
	udpAddr, err := net.ResolveUDPAddr("udp", ":39999")
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {	
		select{
		case data = <-in:
			if len(tempData.Statuses) > len(data.Statuses){
				fmt.Println("Legger til nye element i statuses")
				data.Statuses = append(data.Statuses, tempData.Statuses[len(data.Statuses):]...)
			}else if len(tempData.Statuses) < len(data.Statuses){
				fmt.Println("Fjerner element i statuses")
				for i := 1; i < len(data.Statuses); i++{
					if functions.CheckList(tempData.PrimaryQ, data.PrimaryQ[i]) == false{
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
			if len(tempData.PrimaryQ) == 1{
				conn.SetReadDeadline(time.Now().Add(500*time.Millisecond))
			}
			n, err := conn.Read(buffer) // Høtt skjer om den stoppar her?
			if err == nil{
				err = json.Unmarshal(buffer[0:n], &receivedData)
				checkError(err)
				if functions.CheckList(tempData.PrimaryQ,receivedData.ID)==false {
					fmt.Printf("Går inn i Checklisttingeling: %v, ID: %d", receivedData.PrimaryQ, receivedData.ID)
					tempData.Statuses = append(tempData.Statuses, receivedData.Statuses[GetIndex(receivedData.ID, &receivedData)])
					tempData.PrimaryQ = append(tempData.PrimaryQ, receivedData.ID)
					tempData.ID = receivedData.ID
					SendOrderlist(&tempData,1)
				}else{
					
					tempData.ID = receivedData.ID
					tempData.Statuses[GetIndex(tempData.ID,&tempData)] = tempData.Statuses[GetIndex(tempData.ID, &tempData)]	
				}
				
			}

		}
		
	}
}

/////////// Slave functions //////////// 

func ListenForPrimary(bconn *net.UDPConn,baddr *net.UDPAddr ,in chan *Data, out chan *Data, PrimaryChan chan int) { // Bruke chan muligens fordi den skal skrive til Data
	buffer := make([]byte, 1024)
	var data *Data
	var temp Data
	for {
		data = <-in 
		fmt.Println("Hører")
		bconn.SetReadDeadline(time.Now().Add(5*time.Second))		
		n, err := bconn.Read(buffer)
		if err != nil && data.PrimaryQ[1] == GetID() {
			fmt.Println("Mottar ikke meldinger fra primary lenger, tar over")
			data.Statuses[1].UpList = append(data.Statuses[1].UpList,data.Statuses[0].UpList...)
			data.Statuses[1].DownList = append(data.Statuses[1].DownList,data.Statuses[0].DownList...)
			for j:=0;j<6;j++{
				if(data.ButtonList[j] == 1 && j<3) {
					data.Statuses[1].UpList = append(data.Statuses[1].UpList,j) 
				}else if(data.ButtonList[j] == 1){
					data.Statuses[1].DownList = append(data.Statuses[1].DownList,j-2)
				}
			}
			
			data.PrimaryQ = data.PrimaryQ[1:] 
			data.Statuses = data.Statuses[1:]
			go PrimaryBroadcast(baddr, data)
			go ChannelFunc(PrimaryChan)
			break
		}
		err = json.Unmarshal(buffer[0:n], &temp)
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
		}	
		out <- data
	}	
}

func SlaveUpdate(in chan *Data, out chan *Data) { 
	data := <-in
	out<- data
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:39999")
	conn, err := net.DialUDP("udp",nil, udpAddr)
	checkError(err)
	for {
		 //WRITE
		data = <-in
		data.ID = GetID()
		if(driver.GetFloorSensorSignal()!=-1){	
			data.Statuses[GetIndex(GetID(), data)].CurrentFloor = driver.GetFloorSensorSignal()
		}
		b,_ := json.Marshal(*data)
		conn.Write(b)	
		checkError(err)
		time.Sleep(150*time.Millisecond)
		if data.Statuses[GetIndex(GetID(), data)].Primary == true {
			break
		}
		out<- data
	}
}
func UdpInit(broadcastListenPort int,  data *Data, slaveListenIn chan *Data, slaveListenOut chan *Data, PrimaryChan chan int, SlaveChan chan int) (err error) {
	buffer := make([]byte, 1024)
	var status Status
	status.Primary = false
	SetStatus(&status,0,driver.GetFloorSensorSignal())
	data.PriBroad = false
	//Generating broadcast address
	baddr, err := net.ResolveUDPAddr("udp4", "129.241.187.255:"+strconv.Itoa(broadcastListenPort))
	if err != nil {
		return err
	}
	//Creating listener on broadcast connection
	broadcastListenConn, err := net.ListenUDP("udp", baddr)
	broadcastListenConn.SetReadDeadline(time.Now().Add(3*time.Second))
	n, err := broadcastListenConn.Read(buffer)
	if err != nil {
		data.ButtonList = []int{0,0,0,0,0,0}
		fmt.Println("Tar over som primary!")
		data.PrimaryQ = append(data.PrimaryQ, GetID())
		data.Statuses = append(data.Statuses, status)
		data.Statuses[GetIndex(GetID(), data)].Primary = true
		go PrimaryBroadcast(baddr,data)
		
	} else {
		err = json.Unmarshal(buffer[0:n], data)
		if functions.CheckList(data.PrimaryQ,GetID()) == false{
			data.PrimaryQ = append(data.PrimaryQ, GetID())
			data.Statuses = append(data.Statuses, status)
		}
		go ChannelFunc(SlaveChan)
		time.Sleep(2500*time.Millisecond)
		go ListenForPrimary(broadcastListenConn,baddr, slaveListenIn, slaveListenOut, PrimaryChan)
	}
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
	}
}





