# pidusage
Cross-platform process cpu % and memory usage of a PID for golang

Ideas from https://github.com/soyuka/pidusage but just use Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/struCoder/pidusage)](https://goreportcard.com/report/github.com/struCoder/pidusage)
[![GoDoc](https://godoc.org/github.com/struCoder/pidusage?status.svg)](https://godoc.org/github.com/struCoder/pidusage)

## API

```golang
import (
  "os"
	"github.com/struCoder/pidusage"
)

func printStat() {
	sysInfo, err := pidusage.GetStat(os.Process.Pid)
}
```

## How it works

A check on the `runtime.GOOS` is done to determine the method to use.

### Linux
We use `/proc/{pid}/stat` in addition to the the `PAGE_SIZE` and the `CLK_TCK` direclty from `getconf()` command. Uptime comes from `proc/uptime`

Cpu usage is computed by following [those instructions](http://stackoverflow.com/questions/16726779/how-do-i-get-the-total-cpu-usage-of-an-application-from-proc-pid-stat/16736599#16736599). It keeps an history of the current processor time for the given pid so that the computed value gets more and more accurate. Don't forget to do `unmonitor(pid)` so that history gets cleared.
Cpu usage does not check the child process tree!

Memory result is representing the RSS (resident set size) only by doing `rss*pagesize`, where `pagesize` is the result of `getconf PAGE_SIZE`.

### On darwin, freebsd, solaris
We use a fallback with the `ps -o pcpu,rss -p PID` command to get the same informations.

Memory usage will also display the RSS only, process cpu usage might differ from a distribution to another. Please check the correspoding `man ps` for more insights on the subject.

### On AIX
AIX is tricky because I have no AIX test environement, at the moment we use: `ps -o pcpu,rssize -p PID` but `/proc` results should be more accurate! If you're familiar with the AIX environment and know how to get the same results as we've got with Linux systems.

### Windows
Next version will support
