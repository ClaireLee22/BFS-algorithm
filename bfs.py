from collections import deque

def shortestPath(edges, start, target):
    graph = buildGraph(edges)

    queue = deque([start])

    visited = {start: None} # key: node, value: previous node

    while queue:
        current = queue.popleft()
        if current == target:
            break

        # look at all its neighbors
        for neighbor in graph[current]:
            # if neighbor is not in visisted
            if neighbor not in visited:
                visited[neighbor] = current
                queue.append(neighbor)

    return backtrackPath(visited, target)

def buildGraph(edges):
    graph = {}
    for edge in edges:
        a, b = edge
        if a not in graph:
            graph[a] = []
        if b not in graph:
            graph[b] = []
        
        graph[a].append(b)
        graph[b].append(a)
    return graph

def backtrackPath(visited, end):
    path = []
    current = end
    while current is not None:
        path.append(current)
        current = visited[current]
    
    return path[::-1]


if __name__ == "__main__":
    # undirected graph
    edges = [
        [0, 1],
        [0, 3],
        [1, 2], 
        [1, 5],
        [2, 3],
        [2, 4],
        [2, 5],
        [3, 4],
        [4, 5],
    ]

    # start node: 0, target node: 5
    print("The shortest path: ", shortestPath(edges, 0, 5))

    """
    output: ('The shortest path: ', [0, 1, 5])
    """