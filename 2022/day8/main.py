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

def scenic_score_up(tree):
    score = 0
    for row in range(tree.row, -1, -1):
        if (row != tree.row):
            score += 1
            if treeGrid[row][tree.col] > tree.height:
                return score
    return score

def scenic_score_down(tree):
    score = 0
    for row in range(tree.row, len(treeGrid)):
        if (row != tree.row):
            score += 1
            if treeGrid[row][tree.col] > tree.height:
                return score
    return score

def scenic_score_right(tree):
    score = 0
    for col in range(tree.col, len(treeGrid[tree.row])):
        if (col != tree.col):
            score += 1
            if treeGrid[tree.row][col] > tree.height:
                return score
    return score

def scenic_score_left(tree):
    score = 0
    for col in range(tree.col, -1, -1):
        if (col != tree.col):
            score += 1
            if treeGrid[tree.row][col] > tree.height:
                return score
    return score

maxScore = 0

for i in range(len(treeGrid)):
    for j in range(len(treeGrid[i])):
        tree = Tree(i, j, treeGrid[i][j])
        score = scenic_score_up(tree) * scenic_score_down(tree) * scenic_score_right(tree) * scenic_score_left(tree)
        if score > maxScore:
            maxScore = score

print(maxScore)   
