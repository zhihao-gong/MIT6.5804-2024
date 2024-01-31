package mr

//
// RPC definitions.
//

import (
	"os"
	"strconv"
)

type RpcStatusCode int

const (
	Success RpcStatusCode = iota
	Error
	NotFound
	Unauthorized
	Forbidden
)

type RpcResult struct {
	Code    RpcStatusCode
	Message string
}

type RegisterArgs struct {
}

type RegisterReply struct {
	WorkerId string
	result   RpcResult
}

type PingArgs struct {
	WorkerId string
}

type PingReply struct {
	Result RpcResult
}

type AskForTaskArgs struct {
	WorkerId string
}

type AskForTaskReply struct {
	Task   Task
	result RpcResult
}

type ReportFinishedTaskArgs struct {
	WorkerId string
	TaskId   string
}

type ReportFinishedTaskReply struct {
	result RpcResult
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
