i = open("input").read()

print(i)

a, b = [], []
for l in i.strip().split("\n"):
    vals = l.split("   ")
    a.append(int(vals[0]))
    b.append(int(vals[1]))

a.sort()
b.sort()

c = 0
for i, _ in enumerate(a):
    c += abs(a[i]-b[i])

print(c)
