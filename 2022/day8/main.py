f = open("./input","r")
lines = f.readlines()

treeGrid = []

class Tree:
    def __init__(self, row, col, height):
        self.row = row
        self.col = col
        self.height = height

for line in lines:
    i = line.strip()
    trees = [*i]
    treeGrid.append([int(x) for x in trees])

noOfVisible = 0

def is_visible_up(tree):
    for row in range(tree.row):
        if treeGrid[row][tree.col] >= tree.height and row != tree.row:
            return False
    return True

def is_visible_down(tree):
    for row in range(tree.row, len(treeGrid)):
        if treeGrid[row][tree.col] >= tree.height and row != tree.row:
            return False
    return True

def is_visible_right(tree):
    for col in range(tree.col, len(treeGrid[tree.row])):
        if treeGrid[tree.row][col] >= tree.height and col != tree.col:
            return False
    return True

def is_visible_left(tree):
    for col in range(tree.col):
        if treeGrid[tree.row][col] >= tree.height and col != tree.col:
            return False
    return True

for i in range(len(treeGrid)):
    for j in range(len(treeGrid[i])):
        tree = Tree(i, j, treeGrid[i][j])
        if is_visible_up(tree) or is_visible_down(tree) or is_visible_right(tree) or is_visible_left(tree):
            noOfVisible += 1

print(noOfVisible)   
