import fileinput
import re
from copy import deepcopy

L = [list(l.strip()) for l in list(fileinput.input())]

R = len(L)
C = len(L[0])

def solve(L, p1):
    while True:
        newL = deepcopy(L)

        change = False
        for r in range(R):
            for c in range(C):
                nocc = 0
                for dr in [-1, 0, 1]:
                    for dc in [-1, 0, 1]:
                        if dc == 0 and dr == 0:
                            continue
                        rr = r+dr
                        cc = c+dc

                        while 0<=rr<R and 0<=cc<C and L[rr][cc]=="." and (not p1):
                            rr = rr+dr
                            cc = cc+dc
                        if 0<=rr<R and 0<=cc<C and L[rr][cc] == "#":
                            nocc += 1
                if L[r][c] == 'L':
                    if nocc == 0:
                        newL[r][c] = '#'
                        change = True
                elif L[r][c] == "#" and nocc >= 5:
                    newL[r][c] = 'L'
                    change = True

        if not change:
            break
        L = deepcopy(newL)

    ans = 0
    for r in range(R):
        for c in range(C):
            if L[r][c] == "#":
                ans += 1
    return ans

print(solve(L, False))

