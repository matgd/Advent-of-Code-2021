# Tested in Python{3.6, 3.9}

from collections import Counter
from copy import deepcopy

with open("input.txt") as file:
    starting_polymer = file.readline().strip()
    file.readline()  # skip blank line
    pair_insertions = file.readlines()
    
pairs = [pair.strip().split(' -> ') for pair in pair_insertions]
mapping = { k: f'{k[0]}{v}{k[1]}' for k,v in pairs }

def task1(polymer):
    for _ in range(10):  # 10 steps
        polymer = ''.join([mapping.get(f'{s1}{s2}', f'{s1}{s2}')[:2] for s1, s2 in zip(polymer, polymer[1:])] + [polymer[-1]])

    count = Counter(polymer).most_common()
    print("[Task 1]:", count[0][1] - count[-1][1])

def task2(polymer):  # Thanks ThePrimeagen for the idea!
    polymer_dict = {}
    for s1, s2 in zip(polymer, polymer[1:]):
        polymer_dict.setdefault(f'{s1}{s2}', 0)
        polymer_dict[f'{s1}{s2}'] += 1

    insertion_products = {}
    for k, v in mapping.items():
        insertion_products[k] = [v[:2], v[1:]]

    for _ in range(40):
        new_polymer_dict = deepcopy(polymer_dict)
        for k, v in polymer_dict.items():
            a, b = insertion_products[k]
            new_polymer_dict.setdefault(a, 0)
            new_polymer_dict.setdefault(b, 0)
            new_polymer_dict[a] += v
            new_polymer_dict[b] += v
        for k, v in polymer_dict.items():
            new_polymer_dict[k] -= v
            if new_polymer_dict[k] == 0:
                del new_polymer_dict[k]
        polymer_dict = new_polymer_dict
    
    count_map = {}
    for k, v in polymer_dict.items():  # Count only the first one
        count_map.setdefault(k[0], 0)
        count_map[k[0]] += v
    count_map.setdefault(polymer[-1], 0)  # Last from initial polymer is still last
    count_map[polymer[-1]] += 1

    count = sorted([(v, k) for k, v in count_map.items()], key=lambda x: x[0])
    print("[Task 2]:", count[-1][0] - count[0][0])

if __name__ == '__main__':
    task1(starting_polymer)
    task2(starting_polymer)
