package main

import (
	"log"
	"sync"
	"time"
)

var windowTime int64 = 5
var bucketTime int64 = 1

type Window struct {
	Buckets map[int64]*Bucket
	Mutex   *sync.RWMutex
}

type Bucket struct {
	Value int64
}

func NewWindow() *Window {
	w := &Window{
		Buckets: make(map[int64]*Bucket),
		Mutex:   &sync.RWMutex{},
	}
	
	return w
}

func (w *Window) getCurrentBucket() *Bucket {
	now := time.Now().Unix()
	var bucket *Bucket
	var ok bool
	
	if bucket, ok = w.Buckets[now]; !ok {
		bucket = &Bucket{}
		w.Buckets[now] = bucket
	}
	
	return bucket
}

func (w *Window) removeOldBucket() {
	now := time.Now().Unix() - windowTime
	for timeStamp := range w.Buckets {
		if timeStamp <= now {
			delete(w.Buckets, timeStamp)
		}
	}
}

func (w *Window) Increment(i int64) {
	if i == 0 {
		return
	}
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	
	b := w.getCurrentBucket()
	b.Value = i
	w.removeOldBucket()
}

func (w *Window) Sum(now int64) int64 {
	sum := int64(0)
	w.Mutex.RLock()
	defer w.Mutex.RUnlock()
	
	for timeStamp, bucket := range w.Buckets {
		if timeStamp >= now-windowTime {
			sum += bucket.Value
		}
	}
	return sum
}

func (w *Window) Avg(now int64) int64 {
	return w.Sum(now) / (windowTime / bucketTime)
}

func main() {
	w := NewWindow()
	for _, x := range []int64{4, 23, 2, 6, 12, 8, 17, 41, 6, 8} {
		w.Increment(x)
		time.Sleep(time.Second * time.Duration(bucketTime))
	}
	log.Println(w.Avg(time.Now().Unix()))
}
