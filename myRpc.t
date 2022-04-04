MyRpc interface{
	GetBalance(account int64) (balance int64, status Status)
	Transfer(fromAccount int64, toAccount int64, amount int64) (status Status)
}

AccountType struct{
	accNo int64
	balance int64
}

