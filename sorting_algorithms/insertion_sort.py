def insertion_sort(num:list[int]) -> list[int]:
    for i in range(0, len(num)):
        j = i
        while j > 0 and num[j - 1] > num[j ]:
            num[j],num[j -1] = num[j - 1], num[j]
            j -= 1 
    return num        
    

sorted = insertion_sort([3,2,1,5,7,9,4,5,3,1,2])
print(f"Sorted list is: {sorted}")