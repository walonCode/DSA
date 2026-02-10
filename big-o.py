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


def binary_search(target:int, arr:list[int]) -> bool:
    arr = sorted(arr)
    low = 0
    high = len(arr) - 1
    while low <= high:
        median = (low + high) // 2
        if arr[median] == target:
            return True
        elif arr[median] < target:
            low = median + 1
        else:
            high = median - 1
    return False

print(binary_search(10, [i for i in range(200)])) 

def count_name(list_of_list:list[list[str]], target:str) -> int :
    count = 0
    for values in list_of_list:
        for name in values:
            if name == target:
                count += 1
    
    return count

all_names:list[list[str]] = [
    ["walon", "cosmogirl", "billjane321"],
    ["cosmokiller", "gr8", "cosmojane3"],
    ["walon", "walon"]
]    
print(count_name(all_names,"walon"))    