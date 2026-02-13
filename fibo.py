def fibo(n:int) -> int:
    if n == 0 or n == 1:
        return n
    return fibo(n -1)+fibo(n-2)    
    
print(f"Fibo of 10 is: {fibo(10)}")    


def new_fibo(n):
    if n == 0 or n == 1:
        return n
    grandparent = 0
    parent = 1
    current = 0
    for i in range(n - 1):
        current = parent + grandparent
        grandparent = parent
        parent = current
    return current
    
print(f"New fibo of 10 is: {new_fibo(10)}")