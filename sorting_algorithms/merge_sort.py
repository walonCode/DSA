# spilting the arr in to two
def merge_sort(num:list[int]) -> list[int]:
    if len(num) < 2:
        return num
    half = len(num) // 2
    # doing the spiliting recursively
    sorted_left = merge_sort(num[:half])
    sorted_right= merge_sort(num[half:])
    #finally calling the merge function
    return merge(sorted_left, sorted_right)


def merge(first:list[int], second:list[int]) -> list[int]:
    final_list:list[int] = []
    i,j = 0,0
    #loops through both list
    while i < len(first) and j < len(second):
        if first[i] <= second[j]:
            final_list.append(first[i])
            i += 1
        else:
           final_list.append(second[j])
           j += 1
    
    # add the remaining element in too the final_list     
    final_list.extend(first[i:])
    final_list.extend(second[j:])
    return final_list
    
result = merge_sort([2,4,8,9,1,2,0,0,3,6,8 ,9,10])
print(f"Sorted result is: {result}")
