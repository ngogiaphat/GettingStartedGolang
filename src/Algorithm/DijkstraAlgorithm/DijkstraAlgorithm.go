package main
import (
    "container/heap"
    "fmt"
    "math"
)
//Declare a struct to store the vertex of the graph
type Vertex struct {
    id    int     //ID of vertex
    dist  float64 //Distance to current peak 
    index int     //Index in heap
}
//Declare a slice to store the graph's adjacency list
type EdgeList []map[int]float64
//The function to find the shortest path between two vertices in the graph
func Dijkstra(start int, end int, graph EdgeList) []int {
    n := len(graph)
    dist := make([]float64, n)
    prev := make([]int, n)
    visited := make([]bool, n)

    for i := 0; i < n; i++ {
        dist[i] = math.Inf(1)
    }
    dist[start] = 0
    //Create a heap containing the vertices and the distance to that vertex
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    for i := 0; i < n; i++ {
        v := Vertex{i, dist[i], -1}
        heap.Push(&pq, &v)
    }
    for pq.Len() > 0 {
        //Get the vertex with the minimum distance to start
        u := heap.Pop(&pq).(*Vertex)
        visited[u.id] = true
        //If that vertex is the end vertex then stop
        if u.id == end {
            break
        }
        //Update distance to adjacent vertices of u
        for v, w := range graph[u.id] {
            if !visited[v] {
                alt := u.dist + w
                if alt < dist[v] {
                    dist[v] = alt
                    prev[v] = u.id
                    //Update the distance of vertex v in the heap
                    for _, item := range pq {
                        if item.id == v {
                            item.dist = alt
                            heap.Fix(&pq, item.index)
                            break
                        }
                    }
                }
            }
        }
    }
    //Trace the shortest path
    path := []int{end}
    for prev[path[0]] != start {
        path = append([]int{prev[path[0]]}, path...)
    }
    path = append([]int{start}, path...)
    return path
}
//Defines a type PriorityQueue, which is a slice of pointers to Vertex
type PriorityQueue []*Vertex
//Implement methods of heap.Interface for type PriorityQueue
func (pq PriorityQueue) Len()