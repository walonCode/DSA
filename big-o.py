def find_max(nums: list[int]) -> int:
    if len(nums) == 0:
        return 0
    max = nums[0]
    for value in nums:
        if value > max:
            max = value
    return max


max = find_max([10, 200, 3000, 4])
print(f"The max is {max}")


def get_avg_brand_followers(all_handler: list[list[str]], brand_name: str) -> float:
    if len(all_handler) == 0:
        return 0
    total = 0
    for list in all_handler:
        for brand in list:
            if brand_name in brand:
                total += 1
    return total / len(all_handler)


all_handles = [
    ["cosmofan1010", "cosmogirl", "billjane321"],
    ["cosmokiller", "gr8", "cosmojane3"],
    ["iloveboots", "paperthin"]
]
brand_name = "cosmo"

average = get_avg_brand_followers(all_handles, brand_name)
print(f"Average is {average:.2f}")