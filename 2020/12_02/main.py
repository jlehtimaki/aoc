import fileinput


wx = 10
wy = 1

DX = [0,1,0,-1]
DY = [1,0,-1,0]

x = 0
y = 0


dir_ = 1
L = [l.strip() for l in fileinput.input()]
for line in L:
    cmd = line[0]
    n = int(line[1:])
    if cmd == "N":
        wy += n
    elif cmd == "S":
        wy -= n
    elif cmd == "E":
        wx += n
    elif cmd == "W":
        wx -= n
    elif cmd == "L":
        for _ in range(n//90):
            wx,wy = -wy,wx
    elif cmd == "R":
        for _ in range(n//90):
            wx,wy = wy,-wx
    elif cmd == "F":
        x += n * wx
        y += n * wy
    else:
        assert False
    print(line,cmd,n,dir_,x,y)

print(abs(x)+abs(y))