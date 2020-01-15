# go-loadbalancer-example
ufak bir loadbalancer ornegi // loadbalancer example by Go
go-loadbalancer-example
ufak bir loadbalancer ornegi // loadbalancer example by Go

go run main.go // loadbalancerımız caliscak

servers: cmd'lerden bu sekilde 3 tane server ayaga kaldırdım

npx http-server -p 5001
npx http-server -p 5002 
npx http-server -p 5003

-- curl localhost:8080


curl atınca art arda, 
yükün sırayla 1-2-3-1-2-3.... seklinde dagıtıldıgını goruyoruz.


//notlar / notes
// godoc is important! 

The Dial function connects to a server
conn, err := net.Dial("tcp", "golang.org:80")
if err != nil {
	// handle error
}

//The Listen function creates servers:

ln, err := net.Listen("tcp", ":8080")
if err != nil {
	// handle error
}
for {
	conn, err := ln.Accept()
	if err != nil {
		// handle error
	}
	go handleConnection(conn)
}
