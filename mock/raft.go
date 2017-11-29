// Code generated by mocker; DO NOT EDIT
// github.com/travisjeffery/mocker
package mock

import (
	"github.com/travisjeffery/jocko"
	"sync"
)

var (
	lockRaftAddr      sync.RWMutex
	lockRaftApply     sync.RWMutex
	lockRaftBootstrap sync.RWMutex
	lockRaftIsLeader  sync.RWMutex
	lockRaftLeaderID  sync.RWMutex
	lockRaftShutdown  sync.RWMutex
)

// Raft is a mock implementation of Raft.
//
//     func TestSomethingThatUsesRaft(t *testing.T) {
//
//         // make and configure a mocked Raft
//         mockedRaft := &Raft{
//             AddrFunc: func() string {
// 	               panic("TODO: mock out the Addr method")
//             },
//             ApplyFunc: func(cmd jocko.RaftCommand) error {
// 	               panic("TODO: mock out the Apply method")
//             },
//             BootstrapFunc: func(serf jocko.Serf,serfEventCh <-chan *jocko.ClusterMember,commandCh chan<- jocko.RaftCommand) error {
// 	               panic("TODO: mock out the Bootstrap method")
//             },
//             IsLeaderFunc: func() bool {
// 	               panic("TODO: mock out the IsLeader method")
//             },
//             LeaderIDFunc: func() string {
// 	               panic("TODO: mock out the LeaderID method")
//             },
//             ShutdownFunc: func() error {
// 	               panic("TODO: mock out the Shutdown method")
//             },
//         }
//
//         // TODO: use mockedRaft in code that requires Raft
//         //       and then make assertions.
//
//     }
type Raft struct {
	// AddrFunc mocks the Addr method.
	AddrFunc func() string

	// ApplyFunc mocks the Apply method.
	ApplyFunc func(cmd jocko.RaftCommand) error

	// BootstrapFunc mocks the Bootstrap method.
	BootstrapFunc func(serf jocko.Serf, serfEventCh <-chan *jocko.ClusterMember, commandCh chan<- jocko.RaftCommand) error

	// IsLeaderFunc mocks the IsLeader method.
	IsLeaderFunc func() bool

	// LeaderIDFunc mocks the LeaderID method.
	LeaderIDFunc func() string

	// ShutdownFunc mocks the Shutdown method.
	ShutdownFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// Addr holds details about calls to the Addr method.
		Addr []struct {
		}
		// Apply holds details about calls to the Apply method.
		Apply []struct {
			// Cmd is the cmd argument value.
			Cmd jocko.RaftCommand
		}
		// Bootstrap holds details about calls to the Bootstrap method.
		Bootstrap []struct {
			// Serf is the serf argument value.
			Serf jocko.Serf
			// SerfEventCh is the serfEventCh argument value.
			SerfEventCh <-chan *jocko.ClusterMember
			// CommandCh is the commandCh argument value.
			CommandCh chan<- jocko.RaftCommand
		}
		// IsLeader holds details about calls to the IsLeader method.
		IsLeader []struct {
		}
		// LeaderID holds details about calls to the LeaderID method.
		LeaderID []struct {
		}
		// Shutdown holds details about calls to the Shutdown method.
		Shutdown []struct {
		}
	}
}

// Reset resets the calls made to the mocked APIs.
func (mock *Raft) Reset() {
	lockRaftAddr.Lock()
	mock.calls.Addr = nil
	lockRaftAddr.Unlock()
	lockRaftApply.Lock()
	mock.calls.Apply = nil
	lockRaftApply.Unlock()
	lockRaftBootstrap.Lock()
	mock.calls.Bootstrap = nil
	lockRaftBootstrap.Unlock()
	lockRaftIsLeader.Lock()
	mock.calls.IsLeader = nil
	lockRaftIsLeader.Unlock()
	lockRaftLeaderID.Lock()
	mock.calls.LeaderID = nil
	lockRaftLeaderID.Unlock()
	lockRaftShutdown.Lock()
	mock.calls.Shutdown = nil
	lockRaftShutdown.Unlock()
}

// Addr calls AddrFunc.
func (mock *Raft) Addr() string {
	if mock.AddrFunc == nil {
		panic("moq: Raft.AddrFunc is nil but Raft.Addr was just called")
	}
	callInfo := struct {
	}{}
	lockRaftAddr.Lock()
	mock.calls.Addr = append(mock.calls.Addr, callInfo)
	lockRaftAddr.Unlock()
	return mock.AddrFunc()
}

// AddrCalled returns true if at least one call was made to Addr.
func (mock *Raft) AddrCalled() bool {
	lockRaftAddr.RLock()
	defer lockRaftAddr.RUnlock()
	return len(mock.calls.Addr) > 0
}

// AddrCalls gets all the calls that were made to Addr.
// Check the length with:
//     len(mockedRaft.AddrCalls())
func (mock *Raft) AddrCalls() []struct {
} {
	var calls []struct {
	}
	lockRaftAddr.RLock()
	calls = mock.calls.Addr
	lockRaftAddr.RUnlock()
	return calls
}

// Apply calls ApplyFunc.
func (mock *Raft) Apply(cmd jocko.RaftCommand) error {
	if mock.ApplyFunc == nil {
		panic("moq: Raft.ApplyFunc is nil but Raft.Apply was just called")
	}
	callInfo := struct {
		Cmd jocko.RaftCommand
	}{
		Cmd: cmd,
	}
	lockRaftApply.Lock()
	mock.calls.Apply = append(mock.calls.Apply, callInfo)
	lockRaftApply.Unlock()
	return mock.ApplyFunc(cmd)
}

// ApplyCalled returns true if at least one call was made to Apply.
func (mock *Raft) ApplyCalled() bool {
	lockRaftApply.RLock()
	defer lockRaftApply.RUnlock()
	return len(mock.calls.Apply) > 0
}

// ApplyCalls gets all the calls that were made to Apply.
// Check the length with:
//     len(mockedRaft.ApplyCalls())
func (mock *Raft) ApplyCalls() []struct {
	Cmd jocko.RaftCommand
} {
	var calls []struct {
		Cmd jocko.RaftCommand
	}
	lockRaftApply.RLock()
	calls = mock.calls.Apply
	lockRaftApply.RUnlock()
	return calls
}

// Bootstrap calls BootstrapFunc.
func (mock *Raft) Bootstrap(serf jocko.Serf, serfEventCh <-chan *jocko.ClusterMember, commandCh chan<- jocko.RaftCommand) error {
	if mock.BootstrapFunc == nil {
		panic("moq: Raft.BootstrapFunc is nil but Raft.Bootstrap was just called")
	}
	callInfo := struct {
		Serf        jocko.Serf
		SerfEventCh <-chan *jocko.ClusterMember
		CommandCh   chan<- jocko.RaftCommand
	}{
		Serf:        serf,
		SerfEventCh: serfEventCh,
		CommandCh:   commandCh,
	}
	lockRaftBootstrap.Lock()
	mock.calls.Bootstrap = append(mock.calls.Bootstrap, callInfo)
	lockRaftBootstrap.Unlock()
	return mock.BootstrapFunc(serf, serfEventCh, commandCh)
}

// BootstrapCalled returns true if at least one call was made to Bootstrap.
func (mock *Raft) BootstrapCalled() bool {
	lockRaftBootstrap.RLock()
	defer lockRaftBootstrap.RUnlock()
	return len(mock.calls.Bootstrap) > 0
}

// BootstrapCalls gets all the calls that were made to Bootstrap.
// Check the length with:
//     len(mockedRaft.BootstrapCalls())
func (mock *Raft) BootstrapCalls() []struct {
	Serf        jocko.Serf
	SerfEventCh <-chan *jocko.ClusterMember
	CommandCh   chan<- jocko.RaftCommand
} {
	var calls []struct {
		Serf        jocko.Serf
		SerfEventCh <-chan *jocko.ClusterMember
		CommandCh   chan<- jocko.RaftCommand
	}
	lockRaftBootstrap.RLock()
	calls = mock.calls.Bootstrap
	lockRaftBootstrap.RUnlock()
	return calls
}

// IsLeader calls IsLeaderFunc.
func (mock *Raft) IsLeader() bool {
	if mock.IsLeaderFunc == nil {
		panic("moq: Raft.IsLeaderFunc is nil but Raft.IsLeader was just called")
	}
	callInfo := struct {
	}{}
	lockRaftIsLeader.Lock()
	mock.calls.IsLeader = append(mock.calls.IsLeader, callInfo)
	lockRaftIsLeader.Unlock()
	return mock.IsLeaderFunc()
}

// IsLeaderCalled returns true if at least one call was made to IsLeader.
func (mock *Raft) IsLeaderCalled() bool {
	lockRaftIsLeader.RLock()
	defer lockRaftIsLeader.RUnlock()
	return len(mock.calls.IsLeader) > 0
}

// IsLeaderCalls gets all the calls that were made to IsLeader.
// Check the length with:
//     len(mockedRaft.IsLeaderCalls())
func (mock *Raft) IsLeaderCalls() []struct {
} {
	var calls []struct {
	}
	lockRaftIsLeader.RLock()
	calls = mock.calls.IsLeader
	lockRaftIsLeader.RUnlock()
	return calls
}

// LeaderID calls LeaderIDFunc.
func (mock *Raft) LeaderID() string {
	if mock.LeaderIDFunc == nil {
		panic("moq: Raft.LeaderIDFunc is nil but Raft.LeaderID was just called")
	}
	callInfo := struct {
	}{}
	lockRaftLeaderID.Lock()
	mock.calls.LeaderID = append(mock.calls.LeaderID, callInfo)
	lockRaftLeaderID.Unlock()
	return mock.LeaderIDFunc()
}

// LeaderIDCalled returns true if at least one call was made to LeaderID.
func (mock *Raft) LeaderIDCalled() bool {
	lockRaftLeaderID.RLock()
	defer lockRaftLeaderID.RUnlock()
	return len(mock.calls.LeaderID) > 0
}

// LeaderIDCalls gets all the calls that were made to LeaderID.
// Check the length with:
//     len(mockedRaft.LeaderIDCalls())
func (mock *Raft) LeaderIDCalls() []struct {
} {
	var calls []struct {
	}
	lockRaftLeaderID.RLock()
	calls = mock.calls.LeaderID
	lockRaftLeaderID.RUnlock()
	return calls
}

// Shutdown calls ShutdownFunc.
func (mock *Raft) Shutdown() error {
	if mock.ShutdownFunc == nil {
		panic("moq: Raft.ShutdownFunc is nil but Raft.Shutdown was just called")
	}
	callInfo := struct {
	}{}
	lockRaftShutdown.Lock()
	mock.calls.Shutdown = append(mock.calls.Shutdown, callInfo)
	lockRaftShutdown.Unlock()
	return mock.ShutdownFunc()
}

// ShutdownCalled returns true if at least one call was made to Shutdown.
func (mock *Raft) ShutdownCalled() bool {
	lockRaftShutdown.RLock()
	defer lockRaftShutdown.RUnlock()
	return len(mock.calls.Shutdown) > 0
}

// ShutdownCalls gets all the calls that were made to Shutdown.
// Check the length with:
//     len(mockedRaft.ShutdownCalls())
func (mock *Raft) ShutdownCalls() []struct {
} {
	var calls []struct {
	}
	lockRaftShutdown.RLock()
	calls = mock.calls.Shutdown
	lockRaftShutdown.RUnlock()
	return calls
}
