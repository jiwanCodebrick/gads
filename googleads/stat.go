package v201809

import "time"

type CallStatItem struct {
	Requests  int
	Cached    int
	MemCached int
	TotalTime time.Duration
	ReqTime   time.Duration
	CacheTime time.Duration
}

type CallStat struct {
	CallStatItem
	ServiceStat map[string]*CallStatItem
}

var (
	stat = CallStat{
		ServiceStat: map[string]*CallStatItem{},
	}
)

func (s *CallStat) count(service string, cached, mem bool, t time.Duration) {
	// calc values
	reqDuration := time.Second * 0
	cacheDuration := time.Second * 0
	cachedcnt := 0
	memcachedcnt := 0
	if cached {
		cacheDuration += t
		cachedcnt++
		if mem {
			memcachedcnt++
		}
	} else {
		reqDuration += t
	}
	// update total stat
	s.Requests++
	s.Cached += cachedcnt
	s.MemCached += memcachedcnt
	s.TotalTime += reqDuration + cacheDuration
	s.ReqTime += reqDuration
	s.CacheTime += cacheDuration
	// init service stat
	if _, ok := s.ServiceStat[service]; !ok {
		s.ServiceStat[service] = &CallStatItem{}
	}
	// update service stat
	s.ServiceStat[service].Requests++
	s.ServiceStat[service].Cached += cachedcnt
	s.ServiceStat[service].MemCached += memcachedcnt
	s.ServiceStat[service].TotalTime += reqDuration + cacheDuration
	s.ServiceStat[service].ReqTime += reqDuration
	s.ServiceStat[service].CacheTime += cacheDuration
}

func GetStat() *CallStat {
	return &stat
}
