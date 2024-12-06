i = open("input").read()

# pt 1
print(sum(abs(a-b) for a, b in zip(*[sorted(x) for x in list(map(list, list(zip(*[list(map(int, line.split())) for line in i.splitlines()]))))])))

# pt 2
l, r = list(map(list, list(zip(*[list(map(int, line.split())) for line in i.splitlines()]))))
print(sum(x*r.count(x) for x in l))

# alt
a = []
freq = {}
for l in i.strip().split("\n"):
    vals = l.split("   ")
    a.append(int(vals[0]))
    n = int(vals[1])
    freq[n] = freq.get(n, 0) + 1

c = 0
for v in a:
    c += v * freq.get(v, 0)

print(c)
