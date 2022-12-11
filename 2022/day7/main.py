f = open("./input","r")
lines = f.readlines()

dirs = {}
fullPath = []

for line in lines:
    line = line.strip()
    tokens = line.split(" ")
    if tokens[0] == "$":
        if tokens[1] == "cd":
            if tokens[2] == "..":
                parent = fullPath.pop()
            else:
                fullPath.append(tokens[2])
    elif tokens[0] == "dir":
        pass
    elif isinstance(int(tokens[0]), int):
        dirId = ""
        for dir in fullPath:
            dirId += dir
            dirs[dirId] = (dirs[dirId] if dirId in dirs else 0) + int(tokens[0])
    else:
        print("unknown input: " + tokens[0])
        break

print("Out of instructions. Current dir: " + str(fullPath))

smallestRequired = "/"
totalFree = 70000000 - dirs["/"]

for dir, size in dirs.items():
    if dirs[dir] + totalFree >= 30000000:
        if dirs[dir] < dirs[smallestRequired]:
            smallestRequired = dir

print("Smallest dir that can be deleted: " + str(dirs[smallestRequired]))
