// Code generated by mocker; DO NOT EDIT
// github.com/travisjeffery/mocker
package mock

import (
	"github.com/travisjeffery/jocko"
	"sync"
)

var (
	lockSerfBootstrap sync.RWMutex
	lockSerfCluster   sync.RWMutex
	lockSerfID        sync.RWMutex
	lockSerfJoin      sync.RWMutex
	lockSerfMember    sync.RWMutex
	lockSerfShutdown  sync.RWMutex
)

// Serf is a mock implementation of Serf.
//
//     func TestSomethingThatUsesSerf(t *testing.T) {
//
//         // make and configure a mocked Serf
//         mockedSerf := &Serf{
//             BootstrapFunc: func(node *jocko.ClusterMember,reconcileCh chan<- *jocko.ClusterMember) error {
// 	               panic("TODO: mock out the Bootstrap method")
//             },
//             ClusterFunc: func() []*jocko.ClusterMember {
// 	               panic("TODO: mock out the Cluster method")
//             },
//             IDFunc: func() int32 {
// 	               panic("TODO: mock out the ID method")
//             },
//             JoinFunc: func(addrs ...string) (int, error) {
// 	               panic("TODO: mock out the Join method")
//             },
//             MemberFunc: func(memberID int32) *jocko.ClusterMember {
// 	               panic("TODO: mock out the Member method")
//             },
//             ShutdownFunc: func() error {
// 	               panic("TODO: mock out the Shutdown method")
//             },
//         }
//
//         // TODO: use mockedSerf in code that requires Serf
//         //       and then make assertions.
//
//     }
type Serf struct {
	// BootstrapFunc mocks the Bootstrap method.
	BootstrapFunc func(node *jocko.ClusterMember, reconcileCh chan<- *jocko.ClusterMember) error

	// ClusterFunc mocks the Cluster method.
	ClusterFunc func() []*jocko.ClusterMember

	// IDFunc mocks the ID method.
	IDFunc func() int32

	// JoinFunc mocks the Join method.
	JoinFunc func(addrs ...string) (int, error)

	// MemberFunc mocks the Member method.
	MemberFunc func(memberID int32) *jocko.ClusterMember

	// ShutdownFunc mocks the Shutdown method.
	ShutdownFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// Bootstrap holds details about calls to the Bootstrap method.
		Bootstrap []struct {
			// Node is the node argument value.
			Node *jocko.ClusterMember
			// ReconcileCh is the reconcileCh argument value.
			ReconcileCh chan<- *jocko.ClusterMember
		}
		// Cluster holds details about calls to the Cluster method.
		Cluster []struct {
		}
		// ID holds details about calls to the ID method.
		ID []struct {
		}
		// Join holds details about calls to the Join method.
		Join []struct {
			// Addrs is the addrs argument value.
			Addrs []string
		}
		// Member holds details about calls to the Member method.
		Member []struct {
			// MemberID is the memberID argument value.
			MemberID int32
		}
		// Shutdown holds details about calls to the Shutdown method.
		Shutdown []struct {
		}
	}
}

// Reset resets the calls made to the mocked APIs.
func (mock *Serf) Reset() {
	lockSerfBootstrap.Lock()
	mock.calls.Bootstrap = nil
	lockSerfBootstrap.Unlock()
	lockSerfCluster.Lock()
	mock.calls.Cluster = nil
	lockSerfCluster.Unlock()
	lockSerfID.Lock()
	mock.calls.ID = nil
	lockSerfID.Unlock()
	lockSerfJoin.Lock()
	mock.calls.Join = nil
	lockSerfJoin.Unlock()
	lockSerfMember.Lock()
	mock.calls.Member = nil
	lockSerfMember.Unlock()
	lockSerfShutdown.Lock()
	mock.calls.Shutdown = nil
	lockSerfShutdown.Unlock()
}

// Bootstrap calls BootstrapFunc.
func (mock *Serf) Bootstrap(node *jocko.ClusterMember, reconcileCh chan<- *jocko.ClusterMember) error {
	if mock.BootstrapFunc == nil {
		panic("moq: Serf.BootstrapFunc is nil but Serf.Bootstrap was just called")
	}
	callInfo := struct {
		Node        *jocko.ClusterMember
		ReconcileCh chan<- *jocko.ClusterMember
	}{
		Node:        node,
		ReconcileCh: reconcileCh,
	}
	lockSerfBootstrap.Lock()
	mock.calls.Bootstrap = append(mock.calls.Bootstrap, callInfo)
	lockSerfBootstrap.Unlock()
	return mock.BootstrapFunc(node, reconcileCh)
}

// BootstrapCalled returns true if at least one call was made to Bootstrap.
func (mock *Serf) BootstrapCalled() bool {
	lockSerfBootstrap.RLock()
	defer lockSerfBootstrap.RUnlock()
	return len(mock.calls.Bootstrap) > 0
}

// BootstrapCalls gets all the calls that were made to Bootstrap.
// Check the length with:
//     len(mockedSerf.BootstrapCalls())
func (mock *Serf) BootstrapCalls() []struct {
	Node        *jocko.ClusterMember
	ReconcileCh chan<- *jocko.ClusterMember
} {
	var calls []struct {
		Node        *jocko.ClusterMember
		ReconcileCh chan<- *jocko.ClusterMember
	}
	lockSerfBootstrap.RLock()
	calls = mock.calls.Bootstrap
	lockSerfBootstrap.RUnlock()
	return calls
}

// Cluster calls ClusterFunc.
func (mock *Serf) Cluster() []*jocko.ClusterMember {
	if mock.ClusterFunc == nil {
		panic("moq: Serf.ClusterFunc is nil but Serf.Cluster was just called")
	}
	callInfo := struct {
	}{}
	lockSerfCluster.Lock()
	mock.calls.Cluster = append(mock.calls.Cluster, callInfo)
	lockSerfCluster.Unlock()
	return mock.ClusterFunc()
}

// ClusterCalled returns true if at least one call was made to Cluster.
func (mock *Serf) ClusterCalled() bool {
	lockSerfCluster.RLock()
	defer lockSerfCluster.RUnlock()
	return len(mock.calls.Cluster) > 0
}

// ClusterCalls gets all the calls that were made to Cluster.
// Check the length with:
//     len(mockedSerf.ClusterCalls())
func (mock *Serf) ClusterCalls() []struct {
} {
	var calls []struct {
	}
	lockSerfCluster.RLock()
	calls = mock.calls.Cluster
	lockSerfCluster.RUnlock()
	return calls
}

// ID calls IDFunc.
func (mock *Serf) ID() int32 {
	if mock.IDFunc == nil {
		panic("moq: Serf.IDFunc is nil but Serf.ID was just called")
	}
	callInfo := struct {
	}{}
	lockSerfID.Lock()
	mock.calls.ID = append(mock.calls.ID, callInfo)
	lockSerfID.Unlock()
	return mock.IDFunc()
}

// IDCalled returns true if at least one call was made to ID.
func (mock *Serf) IDCalled() bool {
	lockSerfID.RLock()
	defer lockSerfID.RUnlock()
	return len(mock.calls.ID) > 0
}

// IDCalls gets all the calls that were made to ID.
// Check the length with:
//     len(mockedSerf.IDCalls())
func (mock *Serf) IDCalls() []struct {
} {
	var calls []struct {
	}
	lockSerfID.RLock()
	calls = mock.calls.ID
	lockSerfID.RUnlock()
	return calls
}

// Join calls JoinFunc.
func (mock *Serf) Join(addrs ...string) (int, error) {
	if mock.JoinFunc == nil {
		panic("moq: Serf.JoinFunc is nil but Serf.Join was just called")
	}
	callInfo := struct {
		Addrs []string
	}{
		Addrs: addrs,
	}
	lockSerfJoin.Lock()
	mock.calls.Join = append(mock.calls.Join, callInfo)
	lockSerfJoin.Unlock()
	return mock.JoinFunc(addrs...)
}

// JoinCalled returns true if at least one call was made to Join.
func (mock *Serf) JoinCalled() bool {
	lockSerfJoin.RLock()
	defer lockSerfJoin.RUnlock()
	return len(mock.calls.Join) > 0
}

// JoinCalls gets all the calls that were made to Join.
// Check the length with:
//     len(mockedSerf.JoinCalls())
func (mock *Serf) JoinCalls() []struct {
	Addrs []string
} {
	var calls []struct {
		Addrs []string
	}
	lockSerfJoin.RLock()
	calls = mock.calls.Join
	lockSerfJoin.RUnlock()
	return calls
}

// Member calls MemberFunc.
func (mock *Serf) Member(memberID int32) *jocko.ClusterMember {
	if mock.MemberFunc == nil {
		panic("moq: Serf.MemberFunc is nil but Serf.Member was just called")
	}
	callInfo := struct {
		MemberID int32
	}{
		MemberID: memberID,
	}
	lockSerfMember.Lock()
	mock.calls.Member = append(mock.calls.Member, callInfo)
	lockSerfMember.Unlock()
	return mock.MemberFunc(memberID)
}

// MemberCalled returns true if at least one call was made to Member.
func (mock *Serf) MemberCalled() bool {
	lockSerfMember.RLock()
	defer lockSerfMember.RUnlock()
	return len(mock.calls.Member) > 0
}

// MemberCalls gets all the calls that were made to Member.
// Check the length with:
//     len(mockedSerf.MemberCalls())
func (mock *Serf) MemberCalls() []struct {
	MemberID int32
} {
	var calls []struct {
		MemberID int32
	}
	lockSerfMember.RLock()
	calls = mock.calls.Member
	lockSerfMember.RUnlock()
	return calls
}

// Shutdown calls ShutdownFunc.
func (mock *Serf) Shutdown() error {
	if mock.ShutdownFunc == nil {
		panic("moq: Serf.ShutdownFunc is nil but Serf.Shutdown was just called")
	}
	callInfo := struct {
	}{}
	lockSerfShutdown.Lock()
	mock.calls.Shutdown = append(mock.calls.Shutdown, callInfo)
	lockSerfShutdown.Unlock()
	return mock.ShutdownFunc()
}

// ShutdownCalled returns true if at least one call was made to Shutdown.
func (mock *Serf) ShutdownCalled() bool {
	lockSerfShutdown.RLock()
	defer lockSerfShutdown.RUnlock()
	return len(mock.calls.Shutdown) > 0
}

// ShutdownCalls gets all the calls that were made to Shutdown.
// Check the length with:
//     len(mockedSerf.ShutdownCalls())
func (mock *Serf) ShutdownCalls() []struct {
} {
	var calls []struct {
	}
	lockSerfShutdown.RLock()
	calls = mock.calls.Shutdown
	lockSerfShutdown.RUnlock()
	return calls
}
