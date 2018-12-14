package main

import (
	"flag"
	"debug/elf"
	"log"
	"fmt"
	"strings"
	"syscall"
	"io/ioutil"
	"strconv"
)

func GetRdmaWorkerTid(chunk_pid int) (tid int, err error) {

	taskPath := "/proc/"+strconv.Itoa(chunk_pid)+"/task"
	files, err := ioutil.ReadDir(taskPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		statusFile := taskPath + "/" + f.Name() + "/status"
		//log.Println("statusFile=" + statusFile)
		data, err := ioutil.ReadFile(statusFile)
		if err != nil {
			log.Fatal("Error in reading " + statusFile)
		}
		if(strings.Contains(string(data), "rdma-worker")) {
			return strconv.Atoi(f.Name())
		}
	}

	return 0, nil;
}

func GetErrorInjectionFunAddr(exePath string) (funAddr uintptr, err error){
	exe, err := elf.Open(exePath);
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	syms, _:= exe.Symbols()
	for _, sym := range syms {
		if (strings.Contains(sym.Name, "ErrorInjection")) {
			fmt.Print(sym.Name + "=")
			fmt.Println(sym.Value)
			funAddr = (uintptr)(sym.Value)
		}
	}
	return funAddr, nil;
}

func setBreakpoint(pid int, breakpoint uintptr) []byte {
	original := make([]byte, 1)
	_, err := syscall.PtracePeekData(pid, breakpoint, original)
	if err != nil {
		log.Fatal(err)
	}
	_, err = syscall.PtracePokeData(pid, breakpoint, []byte{0xCC})
	if err != nil {
		log.Fatal(err)
	}
	return original
}
func clearBreakpoint(pid int, breakpoint uintptr, original []byte) {
	_, err := syscall.PtracePokeData(pid, breakpoint, original)
	if err != nil {
		log.Fatal(err)
	}
}

func DoErrorInjection(tid int, funAddr uintptr, op_code uint, not_submit bool, sct uint, cnt uint) (err error) {
	err = syscall.PtraceAttach(tid);
	if err != nil {
		log.Fatal(err)
		return err
	}
	var ws syscall.WaitStatus
	wpid, err := syscall.Wait4(tid, &ws, syscall.WALL, nil)
	if wpid == -1 {
		log.Fatal(err)
		return err
	}
	if ws.Exited() {
		return nil
	}
	orig_code := setBreakpoint(tid, funAddr)

	err = syscall.PtraceCont(tid, 0)
	if err != nil {
		log.Fatal(err)
	}
	_, err = syscall.Wait4(tid, &ws, syscall.WALL, nil)
	clearBreakpoint(tid, funAddr, orig_code)

	var regs syscall.PtraceRegs
	err = syscall.PtraceGetRegs(tid, &regs)
	if err != nil {
		log.Fatal(err)
	}
	regs.Rip = uint64(funAddr);

	//rdi, rsi, rdx, rcx, r8, r9
	regs.Rsi = 1
	regs.Rdx = (uint64)(op_code)
	if not_submit {
		regs.Rcx = 1
	} else {
		regs.Rcx = 0
	}
	regs.R8 = (uint64)(sct)
	regs.R9 = (uint64)(cnt)
	err = syscall.PtraceSetRegs(tid, &regs)
	if err != nil {
		log.Fatal(err)
	}
	syscall.PtraceDetach(tid)

	return nil
}

// <pid>  uint8_t op_code, bool not_submit, uint8_t sct,
func main()  {
	chunk_pid := flag.Int("pid", 0, "chunk pid")
	op_code := flag.Uint("op_code", 2, "op_code")
	not_submit := flag.Bool("not_submit", false, "not_submit")
	sct := flag.Uint("sct", 2, "sct")
	cnt := flag.Uint("cnt", 1, "cnt")

	flag.Parse();

	funAddr, err := GetErrorInjectionFunAddr("/proc/"+strconv.Itoa(*chunk_pid)+"/exe");
	if err != nil {
		return;
	}
	tid, err := GetRdmaWorkerTid(*chunk_pid)
	log.Println("RdmaWorker Tid =" + strconv.Itoa(tid))

	err = DoErrorInjection(tid, funAddr, *op_code, *not_submit, *sct, *cnt);

	if err != nil {
		log.Println("ErrorInjection Failed!")
	} else{
		log.Println("ErrorInjection Sucessfully!")
	}

}