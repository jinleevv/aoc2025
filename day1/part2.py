def main():
    num = 50
    res = 0

    f = open('day1/data.txt')

    lines = f.readlines()
    f.close()

    for line in lines:
        instruction = line.strip("\n")
        direction, rotation = instruction[0], int(instruction[1:])

        res += rotation // 100
        remainder = rotation % 100

        if direction == "R":
            if num + remainder >= 100:
                res += 1
            num = (num + remainder) % 100
        else:
            if num != 0 and num - remainder <= 0:
                res += 1
            num = (num - remainder) % 100

    print(res)



if __name__ == "__main__":
    main()        