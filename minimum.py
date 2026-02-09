def find_minimun(nums: list[int]) -> int:
    if len(nums) == 0:
        return 0
    least: int = nums[0]
    for value in nums:
        if value < least:
            least = value

    return least


value = find_minimun([-12, 1222, 122])
print(f"Smallest value is {value}")
