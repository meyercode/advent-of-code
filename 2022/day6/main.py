f = open("./input","r")
datastream = [*f.readlines()[0]]

windowMin = 0
windowMax = 3

startPacketFound = False

while not startPacketFound:
    window = {}
    for i in range(windowMin, windowMax+1):
        if datastream[i] in window:
            windowMin += 1
            windowMax += 1
            break
        else:
            window[datastream[i]] = True
        if len(window) == 4:
            startPacketFound = True

print(windowMax+1)
