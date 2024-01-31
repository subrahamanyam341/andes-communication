package rating

import (
	"bytes"
	"errors"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/subrahamanyam341/andes-communication/p2p"
	"github.com/subrahamanyam341/andes-communication/p2p/mock"
	"github.com/subrahamanyam341/andes-communication/testscommon"
	"github.com/subrahamanyam341/andes-core-16/core"
	coreMocks "github.com/subrahamanyam341/andes-core-16/data/mock"
	logger "github.com/subrahamanyam341/andes-logger-123"
)

func createMockArgs() ArgPeersRatingHandler {
	return ArgPeersRatingHandler{
		TopRatedCache: &mock.CacherStub{},
		BadRatedCache: &mock.CacherStub{},
		Logger:        &testscommon.LoggerStub{},
	}
}

func TestNewPeersRatingHandler(t *testing.T) {
	t.Parallel()

	t.Run("nil top rated cache should error", func(t *testing.T) {
		t.Parallel()

		args := createMockArgs()
		args.TopRatedCache = nil

		prh, err := NewPeersRatingHandler(args)
		assert.True(t, errors.Is(err, p2p.ErrNilCacher))
		assert.True(t, strings.Contains(err.Error(), "TopRatedCache"))
		assert.Nil(t, prh)
	})
	t.Run("nil bad rated cache should error", func(t *testing.T) {
		t.Parallel()

		args := createMockArgs()
		args.BadRatedCache = nil

		prh, err := NewPeersRatingHandler(args)
		assert.True(t, errors.Is(err, p2p.ErrNilCacher))
		assert.True(t, strings.Contains(err.Error(), "BadRatedCache"))
		assert.Nil(t, prh)
	})
	t.Run("nil logger should error", func(t *testing.T) {
		t.Parallel()

		args := createMockArgs()
		args.Logger = nil

		prh, err := NewPeersRatingHandler(args)
		assert.Equal(t, p2p.ErrNilLogger, err)
		assert.Nil(t, prh)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		prh, err := NewPeersRatingHandler(createMockArgs())
		assert.Nil(t, err)
		assert.NotNil(t, prh)
	})
}

func TestPeersRatingHandler_IncreaseRating(t *testing.T) {
	t.Parallel()

	t.Run("new peer should add to cache", func(t *testing.T) {
		t.Parallel()

		wasCalled := false
		providedPid := core.PeerID("provided pid")
		args := createMockArgs()
		args.TopRatedCache = &mock.CacherStub{
			GetCalled: func(key []byte) (value interface{}, ok bool) {
				return nil, false
			},
			PutCalled: func(key []byte, value interface{}, sizeInBytes int) (evicted bool) {
				assert.True(t, bytes.Equal(providedPid.Bytes(), key))

				wasCalled = true
				return false
			},
		}
		args.BadRatedCache = &mock.CacherStub{
			GetCalled: func(key []byte) (value interface{}, ok bool) {
				return nil, false
			},
		}
		prh, _ := NewPeersRatingHandler(args)
		assert.NotNil(t, prh)

		prh.IncreaseRating(providedPid)
		assert.True(t, wasCalled)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		cacheMap := make(map[string]interface{})
		providedPid := core.PeerID("provided pid")
		args := createMockArgs()
		puCalledCounter := 0
		args.TopRatedCache = &mock.CacherStub{
			GetCalled: func(key []byte) (value interface{}, ok bool) {
				val, found := cacheMap[string(key)]
				return val, found
			},
			PutCalled: func(key []byte, value interface{}, sizeInBytes int) (evicted bool) {
				cacheMap[string(key)] = value
				puCalledCounter++
				return false
			},
		}

		prh, _ := NewPeersRatingHandler(args)
		assert.NotNil(t, prh)

		prh.IncreaseRating(providedPid)
		val, found := cacheMap[string(providedPid.Bytes())]
		assert.True(t, found)
		assert.Equal(t, int32(0), val)

		// exceed the limit
		numOfCalls := 100
		for i := 0; i < numOfCalls; i++ {
			prh.IncreaseRating(providedPid)
		}
		val, found = cacheMap[string(providedPid.Bytes())]
		assert.True(t, found)
		assert.Equal(t, int32(maxRating), val)
		assert.Equal(t, numOfCalls+1, puCalledCounter)
	})
}

func TestPeersRatingHandler_DecreaseRating(t *testing.T) {
	t.Parallel()

	t.Run("new peer should add to cache", func(t *testing.T) {
		t.Parallel()

		wasCalled := false
		providedPid := core.PeerID("provided pid")
		args := createMockArgs()
		args.TopRatedCache = &mock.CacherStub{
			GetCalled: func(key []byte) (value interface{}, ok bool) {
				return nil, false
			},
			PutCalled: func(key []byte, value interface{}, sizeInBytes int) (evicted bool) {
				assert.True(t, bytes.Equal(providedPid.Bytes(), key))

				wasCalled = true
				return false
			},
		}
		args.BadRatedCache = &mock.CacherStub{
			GetCalled: func(key []byte) (value interface{}, ok bool) {
				return nil, false
			},
		}
		prh, _ := NewPeersRatingHandler(args)
		assert.NotNil(t, prh)

		prh.DecreaseRating(providedPid)
		assert.True(t, wasCalled)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		topRatedCacheMap := make(map[string]interface{})
		badRatedCacheMap := make(map[string]interface{})
		providedPid := core.PeerID("provided pid")
		args := createMockArgs()
		putTopCalledCounter := 0
		args.TopRatedCache = &mock.CacherStub{
			GetCalled: func(key []byte) (value interface{}, ok bool) {
				val, found := topRatedCacheMap[string(key)]
				return val, found
			},
			PutCalled: func(key []byte, value interface{}, sizeInBytes int) (evicted bool) {
				topRatedCacheMap[string(key)] = value
				putTopCalledCounter++
				return false
			},
			RemoveCalled: func(key []byte) {
				delete(topRatedCacheMap, string(key))
			},
		}
		putBadCalledCounter := 0
		args.BadRatedCache = &mock.CacherStub{
			GetCalled: func(key []byte) (value interface{}, ok bool) {
				val, found := badRatedCacheMap[string(key)]
				return val, found
			},
			PutCalled: func(key []byte, value interface{}, sizeInBytes int) (evicted bool) {
				badRatedCacheMap[string(key)] = value
				putBadCalledCounter++
				return false
			},
			RemoveCalled: func(key []byte) {
				delete(badRatedCacheMap, string(key))
			},
		}

		prh, _ := NewPeersRatingHandler(args)
		assert.NotNil(t, prh)

		// first call adds it with specific rating
		prh.DecreaseRating(providedPid)
		val, found := topRatedCacheMap[string(providedPid.Bytes())]
		assert.True(t, found)
		assert.Equal(t, int32(0), val)
		assert.Equal(t, 1, putTopCalledCounter)
		assert.Equal(t, 0, putBadCalledCounter)

		// exceed the limit
		numOfCalls := 200
		for i := 0; i < numOfCalls; i++ {
			prh.DecreaseRating(providedPid)
		}
		val, found = badRatedCacheMap[string(providedPid.Bytes())]
		assert.True(t, found)
		assert.Equal(t, int32(minRating), val)
		assert.Equal(t, 1, putTopCalledCounter)
		assert.Equal(t, numOfCalls, putBadCalledCounter)

		// move back to top tier
		for i := 0; i < numOfCalls; i++ {
			prh.IncreaseRating(providedPid)
		}
		_, found = badRatedCacheMap[string(providedPid.Bytes())]
		assert.False(t, found)

		val, found = topRatedCacheMap[string(providedPid.Bytes())]
		assert.True(t, found)
		assert.Equal(t, int32(maxRating), val)
		expectedBadPutCalled := numOfCalls + 49 // needs 49 calls from -100 to -2
		expectedTopPutCalled := 1 + numOfCalls - 49
		assert.Equal(t, expectedTopPutCalled, putTopCalledCounter)
		assert.Equal(t, expectedBadPutCalled, putBadCalledCounter)
	})
}

func TestPeersRatingHandler_GetTopRatedPeersFromList(t *testing.T) {
	t.Parallel()

	t.Run("asking for 0 peers should return empty list", func(t *testing.T) {
		t.Parallel()

		prh, _ := NewPeersRatingHandler(createMockArgs())
		assert.NotNil(t, prh)

		res := prh.GetTopRatedPeersFromList([]core.PeerID{"pid"}, 0)
		assert.Equal(t, 0, len(res))
	})
	t.Run("nil provided list should return empty list", func(t *testing.T) {
		t.Parallel()

		prh, _ := NewPeersRatingHandler(createMockArgs())
		assert.NotNil(t, prh)

		res := prh.GetTopRatedPeersFromList(nil, 1)
		assert.Equal(t, 0, len(res))
	})
	t.Run("no peers in maps should add them to cachers and return them", func(t *testing.T) {
		t.Parallel()

		prh, _ := NewPeersRatingHandler(createMockArgs())
		assert.NotNil(t, prh)

		providedListOfPeers := []core.PeerID{"pid 1", "pid 2"}
		res := prh.GetTopRatedPeersFromList(providedListOfPeers, 5)
		assert.Equal(t, providedListOfPeers, res)
	})
	t.Run("one peer in top rated, asking for one should work", func(t *testing.T) {
		t.Parallel()

		providedPid := core.PeerID("provided pid")
		args := createMockArgs()
		args.TopRatedCache = &mock.CacherStub{
			LenCalled: func() int {
				return 1
			},
			KeysCalled: func() [][]byte {
				return [][]byte{providedPid.Bytes()}
			},
			HasCalled: func(key []byte) bool {
				return bytes.Equal(key, providedPid.Bytes())
			},
		}
		prh, _ := NewPeersRatingHandler(args)
		assert.NotNil(t, prh)

		providedListOfPeers := []core.PeerID{providedPid}
		res := prh.GetTopRatedPeersFromList(providedListOfPeers, 1)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, providedPid, res[0])
	})
	t.Run("one peer in each, asking for two should work", func(t *testing.T) {
		t.Parallel()

		providedTopPid := core.PeerID("provided top pid")
		providedBadPid := core.PeerID("provided bad pid")
		args := createMockArgs()
		args.TopRatedCache = &mock.CacherStub{
			LenCalled: func() int {
				return 1
			},
			KeysCalled: func() [][]byte {
				return [][]byte{providedTopPid.Bytes()}
			},
			HasCalled: func(key []byte) bool {
				return bytes.Equal(key, providedTopPid.Bytes())
			},
		}
		args.BadRatedCache = &mock.CacherStub{
			LenCalled: func() int {
				return 1
			},
			KeysCalled: func() [][]byte {
				return [][]byte{providedBadPid.Bytes()}
			},
			HasCalled: func(key []byte) bool {
				return bytes.Equal(key, providedBadPid.Bytes())
			},
		}
		prh, _ := NewPeersRatingHandler(args)
		assert.NotNil(t, prh)

		providedListOfPeers := []core.PeerID{providedTopPid, providedBadPid}
		expectedListOfPeers := []core.PeerID{providedTopPid, providedBadPid}
		res := prh.GetTopRatedPeersFromList(providedListOfPeers, 2)
		assert.Equal(t, expectedListOfPeers, res)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		providedPid1, providedPid2, providedPid3 := core.PeerID("provided pid 1"), core.PeerID("provided pid 2"), core.PeerID("provided pid 3")
		args := createMockArgs()
		args.TopRatedCache = coreMocks.NewCacherMock()
		args.Logger = &testscommon.LoggerStub{
			GetLevelCalled: func() logger.LogLevel {
				return logger.LogTrace // coverage
			},
		}
		prh, _ := NewPeersRatingHandler(args)
		assert.NotNil(t, prh)

		extraPid := core.PeerID("extra pid")
		providedListOfPeers := []core.PeerID{providedPid1, providedPid2, providedPid3, extraPid}
		res := prh.GetTopRatedPeersFromList(providedListOfPeers, 4)
		assert.Equal(t, providedListOfPeers, res)
	})
}

func TestPeersRatingHandler_MultiplePIDsShouldWork(t *testing.T) {
	t.Parallel()

	args := createMockArgs()
	args.TopRatedCache = coreMocks.NewCacherMock()
	args.BadRatedCache = coreMocks.NewCacherMock()

	prh, _ := NewPeersRatingHandler(args)
	assert.NotNil(t, prh)

	numOps := 200
	var wg sync.WaitGroup
	wg.Add(numOps)
	for i := 0; i < numOps; i++ {
		go func(idx int) {
			switch idx % 8 {
			case 0:
				prh.IncreaseRating("pid1")
			case 1:
				prh.IncreaseRating("pid2")
			case 2:
				prh.IncreaseRating("pid3")
			case 3:
				prh.IncreaseRating("pid4")
			case 4:
				prh.DecreaseRating("pid1")
			case 5:
				prh.DecreaseRating("pid2")
			case 6:
				prh.DecreaseRating("pid3")
			case 7:
				prh.DecreaseRating("pid4")
			default:
				assert.Fail(t, "should not get other values")
			}
			wg.Done()
		}(i)
		time.Sleep(time.Millisecond * 10)
	}
	wg.Wait()
}

func TestPeersRatingHandler_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	var prh *peersRatingHandler
	assert.True(t, prh.IsInterfaceNil())

	prh, _ = NewPeersRatingHandler(createMockArgs())
	assert.False(t, prh.IsInterfaceNil())
}
