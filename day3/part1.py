def main():
    f = open('day3/data.txt')
    lines = f.readlines()
    f.close()
    res = 0
    for line in lines:
        line = line.strip('\n')
        
        l, r = 0, 1
        lineRes = int(line[l] + line[r])
        
        while r < len(line):
            if l == r:
                r += 1
                continue
            lineRes = max(lineRes, int(line[l] + line[r]))
            if line[l] < line[r]:
                l += 1
            else:
                r += 1
        res += lineRes
        print(lineRes)

    print(res)


if __name__ == "__main__":
    main()