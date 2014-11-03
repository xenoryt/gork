package path

import (
	"container/heap"
)

//Point is a point on the graph. Should have a function that
//returns its location on the graph.
type Point interface {
	GetLoc() (int, int)
}

//Visit is called when a search algorithm determines a point is
//on the path and visits it.
type Visit func(x, y int)

type Paths func(x, y int) []Point

//Cost should be a function that returns the cost of going from
//point a to point b
type Cost func(ax, ay, bx, by int) int

//Heuristic should return an estimate of how far apart a and b are
type Heuristic func(ax, ay, bx, by int) int

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//Line will trace a line given the endpoints and visit any
//cell the line touches.
func Line(startx, starty, endx, endy int, visit Visit) {
	dx := abs(endx - startx)
	dy := abs(endy - starty)
	n := 1 + dx + dy
	x_inc := 1
	y_inc := 1

	if endx < startx {
		x_inc = -1
	}
	if endy < starty {
		y_inc = -1
	}

	err := dx - dy

	dx *= 2
	dy *= 2

	x, y := startx, starty
	for ; n > 0; n-- {
		visit(x, y)

		if err > 0 {
			x += x_inc
			err -= dy
		} else {
			y += y_inc
			err += dx
		}
	}
}

type astarNode struct {
	point    *Point
	prev     *astarNode
	estimate int
	cost     int
}

type pointQueue []*astarNode

func (pq pointQueue) Len() int { return len(pq) }
func (pq pointQueue) Less(i, j int) bool {
	return pq[i].estimate < pq[j].estimate
}
func (pq pointQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *pointQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*astarNode))
}
func (pq *pointQueue) Pop() interface{} {
	tmp := *pq
	n := len(tmp)
	point := tmp[n-1]
	*pq = tmp[0 : n-1]
	return point
}

//checks if the new node has the same point as a node in pq but
//with lower cost. If so, change cost and prev and fix the heap
//returns true if an update was done
func (pq pointQueue) update(node *astarNode) bool {
	for i := 0; i < len(pq); i++ {
		if *pq[i].point == *node.point {
			if pq[i].estimate > node.estimate {
				pq[i].estimate = node.estimate
				pq[i].prev = node.prev
				heap.Fix(&pq, i)
			}
			return true
		}
	}
	return false
}

//WeightedAStar searches for a path from a to b using weighted A*
//algorithm. The weight w must be greater or equal to 1.
//If w > 1, then the path may be suboptimal but the search will be faster.
func WeightedAStar(a, b Point, g Cost, h Heuristic, getPaths Paths, visit Visit, w int) {
	newnode := func(p Point, prev *astarNode) *astarNode {
		x, y := p.GetLoc()
		gx, gy := b.GetLoc()
		cost := 0
		heur := h(x, y, gx, gy) * w
		//fmt.Println(x, y, heur)
		if prev != nil {
			px, py := (*prev.point).GetLoc()
			cost = g(px, py, x, y)
			cost += prev.cost
		}
		return &astarNode{&p, prev, cost + heur, cost}
	}
	visitAll := func(node *astarNode) {
		for node != nil {
			//fmt.Println(node.cost)
			visit((*node.point).GetLoc())
			node = node.prev
		}
	}

	//A priority queue of all the possible paths to take
	var paths pointQueue
	//A set of points already visited
	visited := make(map[Point]bool)
	heap.Init(&paths)
	heap.Push(&paths, newnode(a, nil))

	for paths.Len() > 0 {
		path := heap.Pop(&paths).(*astarNode)
		//p1, p2 := (*path.point).GetLoc()
		//b1, b2 := b.GetLoc()
		if *path.point == b {
			visitAll(path)
			break
		}
		visited[*path.point] = true

		//get list of adjacent points
		for _, p := range getPaths((*path.point).GetLoc()) {
			if visited[p] {
				continue
			}
			node := newnode(p, path)
			//make sure this node isn't already in the queue. If it is,
			//replace it if it has a higher cost
			if !paths.update(node) {
				paths.Push(node)
			}
		}
	}
}

//AStar searches for a path from a to b using A* algorithm. The path
//found will always be the most optimal
func AStar(a, b Point, f Cost, g Heuristic, getPaths Paths, visit Visit) {
	WeightedAStar(a, b, f, g, getPaths, visit, 1)
}
