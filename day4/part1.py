def main():
    f = open("day4/data.txt")
    lines = f.readlines()

    directions = [(1, 0), (0, 1), (-1, 0), (0, -1), (1, 1), (-1, 1), (-1, -1), (1, -1)]
    matrix = [list(line.strip("\n")) for line in lines]
    res = 0
    ROWS, COLS = len(matrix), len(matrix[0])

    for r in range(ROWS):
        for c in range(COLS):
            if matrix[r][c] == "@":
                count = 0
                for dir in directions:
                    dr, dc = r + dir[0], c + dir[1]
                    if dr < 0 or dc < 0 or dr >= ROWS or dc >= COLS:
                        continue
                    if matrix[dr][dc] == "@":
                        count += 1
                if count < 4:
                    res += 1
    
    print(res)


if __name__ == "__main__":
    main()