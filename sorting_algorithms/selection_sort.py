def selection_sort(nums:list[int]) -> list[int]:
    for i in range(len(nums)):
        smallest_idx = i
        for j in range(i + 1, len(nums)):
            if nums[smallest_idx] > nums[j]:
                smallest_idx = j
        nums[smallest_idx],nums[i] = nums[i], nums[smallest_idx]
    return nums
        
sorted = selection_sort([2,1,4,6,4,3,2,8,9,2,4,5,6])
print(f"Sorted list: {sorted}")