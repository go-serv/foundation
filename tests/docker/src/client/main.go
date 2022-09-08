package main

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/app/net_parcel/client"
	"github.com/go-serv/foundation/pkg/y/netparcel"
	"github.com/go-serv/foundation/service/net"
	src "go-server-tests-endpoints"
	"log"
	"os"
)

const primeNum = 7

var pingTask = func(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	run := func(task job.TaskInterface) {
		cc := j.GetValue().(netparcel.NetParcelClientInterface)
		//pingOpts := client.PingOptions{}
		//pingOpts.TimeoutMs = 2000
		//cc.WithOptions(pingOpts)
		res, err := cc.Ping(primeNum)
		task.Assert(err)
		if res != primeNum {
			log.Fatalf("plain TCP call, expected %d got %v", primeNum, res)
		}
		task.Done()
	}
	return nil, run, nil
}

func clientX509Pair() *net.X509PemPair {
	certFile := os.Getenv(src.EnvCertClientCertFile)
	certKeyFile := os.Getenv(src.EnvCertClientKeyFile)
	if len(certFile) == 0 || len(certKeyFile) == 0 {
		panic("Client certificate file, env variable not set")
	}
	return &net.X509PemPair{certFile, certKeyFile}
}

func main() {
	var (
		cj job.JobInterface
	)
	certs := make([]*net.X509PemPair, 1)
	unsafeEp := net.NewTcp4Endpoint(src.ServerAddr, src.UnsafePort)
	cj = client.NewClientJob(unsafeEp)
	cj.AddTask(pingTask)
	<-cj.Run()
	if _, err := cj.GetInterruptedBy(); err != nil {
		log.Fatalf("plain TCP call failed with %v", err)
	}

	certs[0] = clientX509Pair()
	secureEp := net.NewTcp4Endpoint(src.ServerAddr, src.TlsAnyPort)
	secureEp.WithNoTrustedPartiesTlsProfile("", certs)
	cj = client.NewClientJob(secureEp)
	cj.AddTask(pingTask)
	<-cj.Run()
	if _, err := cj.GetInterruptedBy(); err != nil {
		log.Fatalf("TLS (no trust) call failed with %v", err)
	}
}
