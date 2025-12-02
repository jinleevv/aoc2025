def repeat_sequence(num):
    num = str(num)
    n = len(num)
    for i in range(1, n):
        pattern = num[:i]
        if pattern * 2 == num:
            return True
    return False

def main():
    f = open("day2/data.txt")
    data = f.readlines()
    ranges = data[0].split(",")

    res = 0    
    for ran in ranges:
        start, end = int(ran.split("-")[0]), int(ran.split("-")[1])
        for i in range(start, end):
            # repeat sequence twice check
            if repeat_sequence(i):
                res += i
    
    print(res)


if __name__ == "__main__":
    main()