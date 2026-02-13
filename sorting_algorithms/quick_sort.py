def quick_sort(nums, low, high):
    if low < high:
        middle = partition(nums,low,high)
        quick_sort(nums, low, middle - 1)
        quick_sort(nums,middle + 1, high)

def partition(nums, low, high):
    pivot = nums[high]
    i = low
    for j in range(low, high):
        if nums[j] < pivot :
            nums[i], nums[j] = nums[j], nums[i]
            i += 1
    nums[i],nums[high] = nums[high],nums[i]
    return i

arr = [2,1,3]  
high = len(arr) - 1
quick_sort(arr, 0, high)
print(f"Sorted array: {arr}")