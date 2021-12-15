import networkx as nx
import numpy as np

risk = np.genfromtxt(fname="input.txt", delimiter=1, dtype=int)
len_x, len_y = len(risk[:,0]), len(risk[0])

g = nx.grid_2d_graph(len_x, len_y, create_using=nx.DiGraph)
for u, (x, y) in g.edges:
    g[u][(x, y)]['weight'] = risk[x, y]
print('[Task 1]: Shortest path:', nx.shortest_path_length(g, (0, 0), (len_x-1, len_y-1), 'weight'))
