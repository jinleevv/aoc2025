def main():
    f = open('day3/data.txt')
    lines = f.readlines()
    f.close()
    res = 0
    for line in lines:
        line = line.strip('\n')

        remove = len(line) - 12
        stack = []
        for digit in line:
            while remove > 0 and stack and stack[-1] < digit:
                stack.pop()
                remove -= 1
            stack.append(digit)

        num = int("".join(stack[:12]))
        res += num
        

    print(res)


if __name__ == "__main__":
    main()