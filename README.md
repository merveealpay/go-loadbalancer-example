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
