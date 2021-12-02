import fileinput

lines = list(fileinput.input())

for line in lines:
    print(eval(line))