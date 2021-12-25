from copy import deepcopy
from dataclasses import dataclass
import itertools


INPUT_FILE = "input.txt"

terrain = []


@dataclass
class Point:
    char: str
    x: int
    y: int

    def __repr__(self) -> str:
        return self.char


def get_neighbour(terrain: list[list[Point]], point: Point) -> Point:
    if point.char == ">":
        x = point.x + 1 if point.x + 1 < len(terrain[point.y]) else 0
        return terrain[point.y][x]
    if point.char == "v":
        y = point.y + 1 if point.y + 1 < len(terrain) else 0
        return terrain[y][point.x]
    raise Exception('Point must be either ">" or "v".')


with open(INPUT_FILE) as file:
    line_no = 0
    for line in file.readlines():
        terrain.append([Point(char=point, x=i, y=line_no) for i, point in enumerate(line.strip())])
        line_no += 1

# Naive, just wanted to write some "brainless" code for the last day
def task1():
    previous_terrain = None
    for step in itertools.count(start=1):
        if step % 20 == 0:
            print(f"Step {step}...")
        terrain_copy = deepcopy(terrain)
        for horizontal_line in terrain_copy:
            for point in horizontal_line:
                if point.char == ">":
                    if (n := get_neighbour(terrain_copy, point)).char == ".":
                        terrain[point.y][point.x].char = "."
                        terrain[n.y][n.x].char = ">"

        terrain_copy = deepcopy(terrain)
        for vertical_index in range(len(terrain_copy[0])):
            vertical_line = [terrain_copy[y][vertical_index] for y in range(len(terrain_copy))]
            for point in vertical_line:
                if point.char == "v":
                    if (n := get_neighbour(terrain_copy, point)).char == ".":
                        terrain[point.y][point.x].char = "."
                        terrain[n.y][n.x].char = "v"
        if previous_terrain == terrain:
            return step
        previous_terrain = deepcopy(terrain)


print(f"Sea cucumbers stay in place after step: {task1()}")
