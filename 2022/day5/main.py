f = open("./input","r")
lines = f.readlines()

stacks = [
    ['D', 'T', 'R', 'B', 'J', 'L', 'W', 'G'],
    ['S', 'W', 'C'],
    ['R', 'Z', 'T', 'M'],
    ['D', 'T', 'C', 'H', 'S', 'P', 'V'],
    ['G', 'P', 'T', 'L', 'D', 'Z'],
    ['F', 'B', 'R', 'Z', 'J', 'Q', 'C', 'D'],
    ['S', 'B', 'D', 'J', 'M', 'F', 'T', 'R'],
    ['L', 'H', 'R', 'B', 'T', 'V', 'M'],
    ['Q', 'P', 'D', 'S', 'V'],
]

instructions = [lines[i] for i in range(10, len(lines))]

for instruction in instructions:
    for stack in stacks:
        print(stack)
    instruction = instruction.strip("\\n")
    values = instruction.split(" ")
    
    src = stacks[int(values[3])-1]
    dest = stacks[int(values[5])-1]
    tempStack = []
    for i in range(0, int(values[1])):
        if len(src) > 0:
            crate = src.pop()
            tempStack.append(crate)
    print(f"{instruction}: {tempStack}")
    while len(tempStack) > 0:
        dest.append(tempStack.pop())

print(stacks)
top = ""
for stack in stacks:
    if len(stack) > 0:
        top += stack[len(stack)-1]
print(top)
