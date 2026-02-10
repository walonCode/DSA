def bubble_sort(num:list[int]) -> list[int]:
    swapping = True
    end = len(num)
    while swapping:
        swapping = False
        for i in range(1, end):
            # check if the left is greater the right
            if num[i - 1] > num[i]:
                #python tuple swap
               num[i], num[i - 1] = num[i -1], num[i]
               swapping = True
        #since the end is already sorted no need for the iteration to reach there again       
        end -= 1       
    return num      


sorted_list = bubble_sort([5,7,3,6,8])
print(f"Sorted list: {sorted_list}")