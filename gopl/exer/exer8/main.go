package exer8

import(
	"log"
	"net"
	"flag"
	"os"
)

func main() {

	portPtr := flag.Int("port", 0, "an int")
	flag.Parse()


	listener, err := net.Listen("tcp", "localhost:%d")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}		
