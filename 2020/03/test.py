import fileinput

slopes = [(1,1),(3,1),(5,1),(7,1),(1,2)]

G = []
for line in fileinput.input():
    G.append(list(line.strip()))

ans = 1

for (dc,dr) in slopes:
    r = 0
    c = 0
    score = 0
    while r+1 < len(G):
        c += dc
        r += dr
        if G[r][c%len(G[r])] == "#":
            score += 1
    ans *= score

print(ans)