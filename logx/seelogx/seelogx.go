package main



import log "github.com/cihub/seelog"

func main() {
	defer log.Flush()
	log.Info("Hello from Seelog!")
	log.Infof("seelogx k0:%s, v0:%s","k0","v0")
}