package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"log"
)

var account1 AccountType
var account2 AccountType


func init() {
	SetupMyRpcGetBalance(getBalance)
	SetupMyRpcTransfer(transfer)
}

func getBalance(account int64) (MyRpcProcedure) {
	//account1:= AccountType {1,34000}
	//account2:= AccountType {2,54000}
	
	log.Printf( "Server: getBalance \n" )

	if account==account1.accNo { 
		log.Printf("Server: Received getBalance request for: %v\n", account1.accNo)
		return &MyRpcGetBalanceReply{account1.balance,syscall.StatusOk}
	}

	if account==account2.accNo {
		log.Printf("Server: Received getBalance request for: %v\n", account2.accNo)
		return &MyRpcGetBalanceReply{account2.balance,syscall.StatusOk}
	}	
	if (account!=account1.accNo && account!=account2.accNo){
		log.Printf("Invalid account! \n")
		return nil
	}
	return &MyRpcGetBalanceReply{-1, syscall.StatusFail}
}

func transfer(fromAccount int64, toAccount int64, amount int64) (MyRpcProcedure) {
	log.Printf( "Server: Transfer \n" )
	if(fromAccount==account1.accNo && toAccount==account2.accNo){ 
		account1.balance -= amount
		account2.balance += amount		
		log.Printf("Server: Transfer request from account %v, to account %v, for amount: %v\n", fromAccount, toAccount, amount)
		return &MyRpcTransferReply{syscall.StatusOk}
	}
	if(fromAccount==account2.accNo && toAccount==account1.accNo){ 
		account2.balance -= amount
		account1.balance += amount		
		log.Printf("Server: Transfer request from account %v, to account %v, for amount: %v\n", fromAccount, toAccount, amount)
		return &MyRpcTransferReply{syscall.StatusOk}
	}
	return &MyRpcTransferReply{syscall.StatusFail} 
	//else StatusFail, will take care of other validations such as invalid account no 
}

func main () {
	account1.accNo = 1
	account1.balance = 34000
	account2.accNo = 2
	account2.balance = 54000

	status2:=altEthos.LogToDirectory("test/ServerX")
	if status2 != syscall.StatusOk {
		log.Printf("ServerX : Directory creation failed!: %v\n", status2)
		altEthos.Exit(status2)
	}

	listeningFd, status:= altEthos.Advertise("myRpc")
	if status!= syscall.StatusOk{
		log.Printf( "Advertising service failed for server: %s \n" , status)
		altEthos.Exit(status)
	}

	for {
		_, fd, status := altEthos.Import(listeningFd)
		if status!=syscall.StatusOk{
			log.Printf( "Error calling Import : %v\n" , status)
			altEthos.Exit(status)
		}	

		log.Printf("Server : new connection accepted \n")

		t := MyRpc{}
		altEthos.Handle(fd,&t)
	}

}
