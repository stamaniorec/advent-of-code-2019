package graph

import (
	"../camera"
)

// Node represents a cell in the viewport grid
type Node struct {
	Row, Col int
	IsIntersection bool
}

// Graph is an adjacency list representation of nodes
type Graph map[Node][]Node

// BuildGraph builds a graph out of scaffold cells in the camera viewport
func BuildGraph(viewport camera.Viewport, u Node) Graph {
	g := Graph(make(map[Node][]Node))
	visited := make(map[Node]bool)

	buildGraph(viewport, u, visited, g)

	return g
}

// FindEulerPath returns a Eulerian path in the graph
// Checks whether the graph is Eulerian are omitted - I've verified it manually,
// since I've seen the camera viewport.
// Fleury's algorithm has been used.
func FindEulerPath(g Graph, u Node) []Node {
	var path []Node
	path = append(path, u)

	findEulerPath(g, u, &path)

	return path
}

func buildGraph(viewport camera.Viewport, u Node, visited map[Node]bool, g Graph) {
	visited[u] = true

	neighborsCoordinates := []struct {
		Row int
		Col int
	}{
		{
			u.Row - 1,
			u.Col,
		},
		{
			u.Row,
			u.Col - 1,
		},
		{
			u.Row,
			u.Col + 1,
		},
		{
			u.Row + 1,
			u.Col,
		},
	}

	for _, coord := range neighborsCoordinates {
		isWithinBounds :=
			coord.Row >= 0 && coord.Col >= 0 &&
				coord.Row < viewport.Rows && coord.Col < viewport.Cols

		if isWithinBounds {
			isScaffold := string(viewport.Grid[coord.Row][coord.Col]) == "#"
			if isScaffold {
				isIntersection := camera.IsScaffoldIntersection(viewport, coord.Row, coord.Col)
				v := Node{coord.Row, coord.Col, isIntersection}

				addDirectedEdge(g, u, v)
				if _, loop := visited[v]; !loop {
					buildGraph(viewport, v, visited, g)
				}
			}
		}
	}
}

func findEulerPath(g Graph, u Node, path *[]Node) {
	for _, v := range g[u] {
		if isNonBridge(g, u, v) {
			*path = append(*path, v)

			removeUndirectedEdge(g, u, v)
			findEulerPath(g, v, path)
		}
	}
}

func addDirectedEdge(g Graph, from, to Node) {
	g[from] = append(g[from], to)
}

func addUndirectedEdge(g Graph, node, neighbor Node) {
	addDirectedEdge(g, node, neighbor)
	addDirectedEdge(g, neighbor, node)
}

func removeDirectedEdge(g Graph, from, to Node) {
	var newAdjList []Node
	for _, adj := range g[from] {
		if adj != to {
			newAdjList = append(newAdjList, adj)
		}
	}
	g[from] = newAdjList
}

func removeUndirectedEdge(g Graph, node, neighbor Node) {
	removeDirectedEdge(g, node, neighbor)
	removeDirectedEdge(g, neighbor, node)
}

// dfsCount returns the number of reachable verticies by performing DFS
func dfsCount(g Graph, u Node, visited map[Node]bool) int {
	if _, isVisited := visited[u]; isVisited {
		return 0
	}

	visited[u] = true

	count := 1
	for _, v := range g[u] {
		count += dfsCount(g, v, visited)
	}
	return count
}

// isNonBridge finds if the edge (u,v) is a non-bridge by removing the edge
// and comparing the number of reachable vertices before and after the removal
func isNonBridge(g Graph, u, v Node) bool {
	if len(g[u]) == 1 {
		return true
	}

	visited := make(map[Node]bool)
	reachableBefore := dfsCount(g, u, visited)

	removeUndirectedEdge(g, u, v)
	visited = make(map[Node]bool)
	reachableAfter := dfsCount(g, u, visited)
	addUndirectedEdge(g, u, v)

	return reachableBefore <= reachableAfter
}

func findPathDfs(g Graph, u Node, visited map[Node]bool, curPath []Node, actualPath *[]Node) {
	visited[u] = true
	curPath = append(curPath, u)

	for _, v := range g[u] {
		_, isVisited := visited[v]
		if !isVisited || v.IsIntersection {
			findPathDfs(g, v, visited, curPath, actualPath)
		}
	}

	delete(visited, u)
}

func FindPath(g Graph, u Node) []Node {
	var curPath, actualPath []Node
	visited := make(map[Node]bool)
	findPathDfs(g, u, visited, curPath, &actualPath)
	return curPath
}
