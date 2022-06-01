//go:build linux && cgo

package sigproc

import (
	"fmt"
	"github.com/go-serv/service/internal/ancillary/platform"
	"os"
	"time"
)

// #include <stdio.h>
// #include <errno.h>
// #include <string.h>
// #include <signal.h>
//
// extern void sigproc_event_hander(int);
// 	char *get_errno() { return strerror(errno); }
//
// int sigqueue_int(int pid, int signal, int value) {
// 		union sigval sigval;
// 		sigval.sival_int = value;
//		printf("Signal: #%d, value: [%d]\n", pid, value);
// 		return sigqueue(pid, signal, sigval);
// }
//
// void sighand(int signo, siginfo_t *info, void *extra)
//{
//       //void *ptr_val = info->si_value.sival_ptr;
//       int int_val = info->si_value.sival_int;
//		 sigproc_event_hander(int_val);
//       printf("Got signal: %d, value: [%d]\n", signo, int_val);
//}
// void init() {
//        struct sigaction action;
//
//        action.sa_flags = SA_SIGINFO;
//        action.sa_sigaction = &sighand;
//		//Test();
//
//        if (sigaction(SIGUSR2, &action, NULL) == -1) {
//                perror("sigusr: sigaction");
//                return;
//        }
// }
//
import "C"

func SignalProcess(id platform.ProcId, sig ProcSignalType) {
	pid := os.Getpid()
	fmt.Printf("Current PID: %d\n", pid)
	C.init()
	C.sigqueue_int(C.int(pid), C.SIGUSR2, C.int(2))
	C.fflush(C.stdout)
	time.Sleep(time.Millisecond * 50)
}
