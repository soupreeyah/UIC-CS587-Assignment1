package main

import(
	"ethos/syscall"
	"ethos/altEthos"
	"log"
)

func init() {
	SetupMyRpcGetBalanceReply(getBalanceReply)
	SetupMyRpcTransferReply(transferReply)
}

var bal int64
var myAccount int64 = 1
var otherAccount int64 = 2
var amt int64 = 1000

func getBalanceReply(balance int64, status syscall.Status) (MyRpcProcedure) {
	log.Printf("Client : Received account balance: %v\n", balance)
	return nil
} 

func transferReply(status syscall.Status) (MyRpcProcedure){
	log.Printf("Client : Received transfer reply from %v, to %v, for amount: %v\n", myAccount, otherAccount, amt)
	return nil
}

func main () {

	altEthos.LogToDirectory ("test/ClientX")

	log.Printf("Client : before call \n")

	fd, status := altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}
	callbal := MyRpcGetBalance{myAccount}
	status = altEthos.ClientCall(fd, &callbal)
	if status != syscall.StatusOk {
		log.Printf("Get Balance failed: %v\n", status)
		altEthos.Exit(status)
	}
	fd, status = altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}
	callbal = MyRpcGetBalance{otherAccount}
	status = altEthos.ClientCall(fd, &callbal)
	if status != syscall.StatusOk {
		log.Printf("Get Balance failed: %v\n", status)
		altEthos.Exit(status)
	}


	fd, status = altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}
	calltrans := MyRpcTransfer{myAccount,otherAccount,amt}
	status = altEthos.ClientCall(fd, &calltrans)
	if status != syscall.StatusOk {
		log.Printf("Transfer failed: %v\n", status)
		altEthos.Exit(status)
	}
	fd, status = altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}
	calltrans = MyRpcTransfer{otherAccount,myAccount,amt}
	status = altEthos.ClientCall(fd, &calltrans)
	if status != syscall.StatusOk {
		log.Printf("Transfer failed: %v\n", status)
		altEthos.Exit(status)
	}

	
	log.Printf( "ClientX : done\n" )
}
