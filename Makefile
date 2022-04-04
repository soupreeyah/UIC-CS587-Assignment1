export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=/usr/lib64/go/pkg/ethos_$(GOARCH)
export GOLINUXINCLUDE=/usr/lib64/go/pkg/linux_$(GOARCH)

export ETHOSROOT=server/rootfs
export MINIMALTDROOT=server/minimaltdfs

.PHONY: all install
all: serverX clientX

myRpc.go: myRpc.t
	$(ETN2GO) . myRpc main $^

serverX: serverX.go myRpc.go
	ethosGo $^

clientX: clientX.go myRpc.go
	ethosGo $^

install: clean myRpc.go serverX clientX
	sudo rm -rf server
	(ethosParams server && cd server && ethosMinimaltdBuilder)
	echo 7 > server/param/sleepTime
	ethosTypeInstall myRpc
	ethosServiceInstall myRpc
	ethosDirCreate $(ETHOSROOT)/services/myRpc   $(ETHOSROOT)/types/spec/myRpc/MyRpc all
	install -D  serverX clientX	$(ETHOSROOT)/programs
	ethosStringEncode /programs/serverX > $(ETHOSROOT)/etc/init/services/serverX
	ethosStringEncode /programs/clientX > $(ETHOSROOT)/etc/init/services/clientX

clean:
	sudo rm -rf server
	rm -rf myRpc/ myRpcIndex/
	rm -f myRpc.go
	rm -f serverX
	rm -f serverX.goo.ethos
	rm -f clientX
	rm -f clientX.goo.ethos
