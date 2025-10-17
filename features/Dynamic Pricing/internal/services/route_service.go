package services

import (
	"container/heap"
	"math"
)

// Very small in-memory sample graph (demo). Replace with OSM/OSRM or PostGIS-based routing in prod.
type edge struct {
	to     string
	weight float64
}

var sampleGraph = map[string][]edge{
	"A": {{"B", 1}, {"C", 4}},
	"B": {{"A", 1}, {"C", 2}, {"D", 5}},
	"C": {{"A", 4}, {"B", 2}, {"D", 1}},
	"D": {{"B", 5}, {"C", 1}},
}

func ComputeOptimalRoute(start, end string) ([]string, float64) {
	dist := map[string]float64{}
	prev := map[string]string{}
	for k := range sampleGraph {
		dist[k] = math.Inf(1)
	}
	dist[start] = 0
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{value: start, priority: 0})
	for pq.Len() > 0 {
		u := heap.Pop(pq).(*Item).value
		if u == end {
			break
		}
		for _, e := range sampleGraph[u] {
			if dist[u]+e.weight < dist[e.to] {
				dist[e.to] = dist[u] + e.weight
				prev[e.to] = u
				heap.Push(pq, &Item{value: e.to, priority: dist[e.to]})
			}
		}
	}
	path := []string{}
	u := end
	if _, ok := prev[u]; !ok && u != start {
		return []string{}, math.Inf(1)
	}
	for {
		path = append([]string{u}, path...)
		if u == start {
			break
		}
		u = prev[u]
	}
	return path, dist[end]
}

type Item struct {
	value    string
	priority float64
	index    int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
