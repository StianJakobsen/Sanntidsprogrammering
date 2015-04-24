// go run networkUDP.cd ..go
package udp
import (."fmt" // Using '.' to avoid prefixing functions with their package names
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
)


type Status struct {
	Running int
	CurrentFloor int
	NextFloor int
	Primary bool
	ID int
	//PrimaryQ [3]string
	CommandList []int
	UpList []int  // slice = slice[:0] for å tømme slicen når sendt til primary
	DownList[]int // slice = slice[:0] for å tømme slicen når sendt til primary
	//PriList[]int
	OrderList []int // sjekke for nye ordrer når primary sender
}

type Data struct {
	Status Status
	Statuses map[int]Status // Oppdatere den her å i UdpInit()
	PrimaryQ []int
}


func SetStatus(Data *Data, running int, NextFloor int) {
	ID := GetID()
	(*Data).Status.Running = running
	(*Data).Status.CurrentFloor = driver.GetFloorSensorSignal()
	(*Data).Status.NextFloor = NextFloor
	(*Data).Status.ID = ID
	//Println(" id i func:", (*Status).ID)
}
func GetID() int {
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
	ut,_ := strconv.Atoi(ipAddr[12:15])
	return ut
}


/////////// Primary functions ////////////

func PrimaryBroadcast(baddr *net.UDPAddr, Data *Data) { // data []byte
	//udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:39998")
	//checkError(err)
	bconn, err := net.DialUDP("udp", nil, baddr)
	checkError(err)
	for {
		Println("SENDER")
		time.Sleep(2500*time.Millisecond)
		
		// WRITE
		b,_ := json.Marshal(Data)
		bconn.Write(b)
		
		checkError(err)
	}

}

func PrimaryListen(data *Data) {
	buffer := make([]byte, 1024)
	var temp Data
	udpAddr, err := net.ResolveUDPAddr("udp", ":39998")
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {		
		n, err := conn.Read(buffer)
		checkError(err)
		//Data = buffer
		err = json.Unmarshal(buffer[0:n], temp)		
		(*data).Statuses[(*data).Status.ID] = (*data).Status // Oppdaterar mottatt status hos primary 
	}
}

/////////// Slave functions //////////// 

func ListenForPrimary(bconn *net.UDPConn, Data *Data) { // Bruke chan muligens fordi den skal skrive til Data
	buffer := make([]byte, 1024)
	//udpAddr, err := net.ResolveUDPAddr("udp", ":39998")
	//conn, err := net.ListenUDP("udp", udpAddr)
	//checkError(err)
	for {
		Println("Hører")
		bconn.SetReadDeadline(time.Now().Add(5*time.Second))		
		n, err := bconn.Read(buffer)
		checkError(err)
		//Data = buffer
		err = json.Unmarshal(buffer[0:n], Data)		
		Println("her er primaryQen:", Data.PrimaryQ)		
		// Printf("Rcv %d bytes: %s\n",n, buffer)	
	}	
}



func SlaveUpdate(data *Data) { // chan muligens, bare oppdatere når det er endringar
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187."+ strconv.Itoa((*data).PrimaryQ[0]) + ":39998")
	conn, err := net.DialUDP("udp",nil, udpAddr)
	checkError(err)
	for {
		time.Sleep(2500*time.Millisecond) // bytte til bare ved endringar etterhvert

		 //WRITE
		b,_ := json.Marshal(data)
		conn.Write(b)	
		checkError(err)
	}
}

// send_ch, receive_ch chan Udp_message
func UdpInit(localListenPort int, broadcastListenPort int, message_size int, data *Data) (err error) {
	buffer := make([]byte, message_size)
	var temp Data

	//(*Status).Primary = false
	(*data).Status.Primary = false	
	SetStatus(data,0,driver.GetFloorSensorSignal())	
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
		Println("Tar over som primary!")
		(*data).Status.Primary = true
		(*data).PrimaryQ = append((*data).PrimaryQ, GetID()) 
		go PrimaryBroadcast(baddr,data)
		//go PrimaryListen()	
	} else {
		err = json.Unmarshal(buffer[0:n], temp)
		(*data).PrimaryQ = temp.PrimaryQ	
		//(*Data).PrimaryQ = append((*Data).PrimaryQ, string(buffer))
		go ListenForPrimary(broadcastListenConn, data)
		go SlaveUpdate(data)
	}
	


	//	fmt.Printf("Generating local address: \t Network(): %s \t String(): %s \n", laddr.Network(), laddr.String())
	//	fmt.Printf("Generating broadcast address: \t Network(): %s \t String(): %s \n", baddr.Network(), baddr.String())
	return err
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
		Println("Sent: ",currentStruct.Teller) 		
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

func checkError(err error) {
	if err != nil {
		Println("Noe gikk galt %v", err)
		return
	}
}











