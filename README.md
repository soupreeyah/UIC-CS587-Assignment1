Assignment 1: Ethos

Creation of a AccountServer and one or more AccountClients in Go. 
The AccountServer will initialize a set of accounts, and then process requests from the clients, updating the accounts stored in the file system.
All transfers and errors are written to logs.

To execute: 


Open one terminal


make install

cd server //here server is the folder in which the filesystem is created 
sudo -E ethosRun -t (batch mode)


or


sudo -E ethosRun


Then open another terminal and navigate to server.


etAl server.ethos

For new instances:


et server.ethos
