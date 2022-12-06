f = open("./input","r")
lines = f.readlines()

intersectCount = 0

for line in lines:
    line = line.strip("\\n")
    pair = line.split(",")
    rangeA = pair[0].split("-")
    rangeB = pair[1].split("-")

    elfA = []
    elfB = []
    for section in range(int(rangeA[0]), int(rangeA[1])+1):
        elfA.append(section)
    for section in range(int(rangeB[0]), int(rangeB[1])+1):
        elfB.append(section)

    aMinSect = elfA[0]
    aMaxSect = elfA[len(elfA)-1]
    
    bMinSect = elfB[0]
    bMaxSect = elfB[len(elfB)-1]

    if aMinSect <= bMinSect and aMaxSect >= bMinSect or bMinSect <= aMinSect and bMaxSect >= aMinSect:
        intersectCount += 1

print(intersectCount)
