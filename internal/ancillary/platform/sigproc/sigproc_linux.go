//go:build linux && cgo

package sigproc

import (
	"fmt"
	"github.com/mesh-master/foundation/internal/ancillary/platform"
	"os"
	"time"
)

/*
#include <stdio.h>
#include <stdint.h>
#include <errno.h>
#include <string.h>
#include <signal.h>

extern void sigproc_event_hander(int, int);
char *get_errno() { return strerror(errno); }

int sigqueue_int(int pid, int signal, int value) {
	union sigval sigval;
	sigval.sival_int = value;
	printf("Signal: #%d, value: [%d]\n", pid, value);
	return sigqueue(pid, signal, sigval);
}

void sighand(int signo, siginfo_t *info, void *extra) {
	int64_t v = info->si_value.sival_int;
	int code = v;
	int extra_val = (v & (~0 >> 10));
	sigproc_event_hander(code, extra_val);
	printf("Got signal: %d, value: [%d]\n", code, extra_val);
}

void init() {
	struct sigaction action;
	action.sa_flags = SA_SIGINFO;
	action.sa_sigaction = &sighand;
	printf("Init");
	if (sigaction(SIGUSR1, &action, NULL) == -1) {
		   perror("sigusr: sigaction");
		   return;
	}
}
*/
import "C"

func SignalProcess(id platform.ProcId, code SigCodeTyp, value SigValueTyp) {
	pid := os.Getpid()
	sigval := value.pack(code)
	fmt.Printf("Current PID: %d code %b sigval %b\n", pid, code, sigval)
	C.init()
	C.sigqueue_int(C.int(pid), C.SIGUSR1, C.int(sigval))
	time.Sleep(time.Millisecond * 50)
	C.fflush(C.stdout)
}
