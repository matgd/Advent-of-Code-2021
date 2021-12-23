import networkx as nx
import numpy as np


def task1():
    risk = np.genfromtxt(fname="input.txt", delimiter=1, dtype=int)
    len_x, len_y = len(risk[:,0]), len(risk[0])

    g = nx.grid_2d_graph(len_x, len_y, create_using=nx.DiGraph)
    for u, (x, y) in g.edges:
        g[u][(x, y)]['weight'] = risk[x, y]
    print('[Task 1]: Shortest path:', nx.shortest_path_length(g, (0, 0), (len_x-1, len_y-1), 'weight'))

def task2():
    risk = np.genfromtxt(fname="input.txt", delimiter=1, dtype=int)
    len_x, len_y = len(risk[:,0]), len(risk[0])
    
    risk_full = risk
    ones = np.ones((len_x, len_y), dtype=np.uint8)
    
    for i in range(1, 5):
        risk_full = np.concatenate((risk_full, risk + ones*i), axis=1)

    first_row = risk_full
    five_seg_ones = np.ones((len_x, len_y*5), dtype=np.uint8)
    for i in range(1, 5):
        risk_full = np.concatenate((risk_full, first_row + five_seg_ones*i), axis=0)

    risk_full[risk_full > 9] -= 9
    g = nx.grid_2d_graph(len_x*5, len_y*5, create_using=nx.DiGraph)
    for u, (x, y) in g.edges:
        g[u][(x, y)]['weight'] = risk_full[x, y]
    print('[Task 2]: Shortest path:', nx.shortest_path_length(g, (0, 0), (len_x*5-1, len_y*5-1), 'weight'))

if __name__ == '__main__':
    task1()
    task2()
