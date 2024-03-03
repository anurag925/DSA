# i still dont know how it works im so bad in coding

def maxsum(arr, n):
    curr, ma = arr[0], arr[0]
    for i in range(1, n):
        curr = curr + arr[i]
        if arr[i] > curr:
            curr = arr[i]
        if curr > ma:
            ma = curr
    return ma

#Complete this function
#Function to find maximum circular subarray sum.
def circularSubarraySum(arr,n):
    subArraySum = maxsum(arr, n)
    if subArraySum < 0:
        return subArraySum
    arrSum = 0
    for i in range(n):
        arrSum += arr[i]
        arr[i] = -arr[i]
        
    subArraySum2 = maxsum(arr, n)
    return max(arrSum + subArraySum2, subArraySum)