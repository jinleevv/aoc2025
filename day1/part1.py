def main():
    num = 50
    res = 0

    f = open('day1/data.txt')

    lines = f.readlines()
    f.close()

    for line in lines:
        instruction = line.strip("\n")
        direction = instruction[0]
        rotation = int(instruction[1:])

        if direction == "R":
            num += rotation
            num = num % 100
        else:
            num -= rotation
            num = num % 100
        
        if num == 0:
            res += 1
    
    print(res)


if __name__ == "__main__":
    main()        