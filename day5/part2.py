from collections import defaultdict


def main():
    f = open("day5/data.txt")
    lines = f.readlines()
    
    res = 0
    range_collections = defaultdict(int)

    for line in lines:
        line = line.strip("\n")
        if "-" in line:
            s, e = int(line.split("-")[0]), int(line.split("-")[1])
            range_collections[s] = max(range_collections[s], e)
        else:
            break
    
    sorted_range = sorted(range_collections.items())
    merged = []
    for start, end in sorted_range:
        if not merged or merged[-1][1] < start:
            merged.append([start, end])
        else:
            merged[-1][1] = max(merged[-1][1], end)

    for start, end in merged:
        num = end - start + 1
        res += num

    print(res)


if __name__ == "__main__":
    main()