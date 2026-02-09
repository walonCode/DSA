def summed(num: list[int]) -> int:
    if len(num) == 0:
        return 0
    total: int = 0
    for value in num:
        total += value

    return total


value = summed([4, 3, 5, 6, 7])

print(f"The sum is {value}")
