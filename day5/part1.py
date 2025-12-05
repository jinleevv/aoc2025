from collections import defaultdict


def main():
    f = open("day5/data.txt")
    lines = f.readlines()
    
    res = 0
    range_collections = defaultdict(int)
    nums = []

    for line in lines:
        line = line.strip("\n")
        if "-" in line:
            s, e = int(line.split("-")[0]), int(line.split("-")[1])
            range_collections[s] = max(range_collections[s], e)
        else:
            if line != "":
                nums.append(int(line))

    for num in nums:
        for start, end in range_collections.items():
            if start <= num <= end:
                res += 1
                break        

    print(res)


if __name__ == "__main__":
    main()