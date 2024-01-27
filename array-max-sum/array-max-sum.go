package main


import "fmt"

func main(){

	 arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	 //fmt.Println("windows size = ", 3)

 	arr= []int{ 1, 4, 2, 10, 2, 3, 1, 0, 20 }
    k := 4
	 // ret := maxArraySum(arr)
	 ret := maxSum(arr,len(arr), k)
	 fmt.Printf("The maximum value for the array %v is %d\n", arr, ret)



	 ret = maxArray(arr, 4)
	 fmt.Printf("The maximum value for the array %v is %d\n", arr, ret)


}

// Returns maximum sum in
    // a subarray of size k using brute force technique.
    func maxSum(arr []int,  n int,  k int) int{
        // Initialize result
        var max_sum int = 0
 
        // Consider all blocks starting with i.
        for i := 0; i < n - k + 1; i++ {
            var current_sum int = 0
            for j := 0; j < k; j++ {
                current_sum = current_sum + arr[i + j];
            }
 
            // Update result if required.

            max_sum = Max(current_sum, max_sum);
        }
 
        return max_sum;
    }

// Uses sliding fixed window (subtract i-1 and add i+k+1)
func maxArray(arr []int, k int) int{
	max_sum:=0
	counter :=0


	// Compute sum of first window of size k
	for i:=0; i < k; i++{
		max_sum +=arr[i]
	}
	counter = k-1

	// Compute sums of remaining windows by
        // removing first element of previous
        // window and adding last element of
        // current window.
    window_sum :=max_sum
    for i:=k; i < len(arr); i++{
    	window_sum +=arr[i]-arr[i-k]
    	max_sum =max(max_sum, window_sum)
    	counter++
    }

    fmt.Println("iterations done ", counter)
    return max_sum


}



func maxArraySum(array []int, k int) int{

	var max_so_far int 
    var max_ending_here int
	for index, val := range array{
		max_ending_here += val
		if( max_so_far < max_ending_here){
			max_so_far = max_ending_here

		}
		// if( max_ending_here < 0 ){
		// 	max_ending_here = 0
		// }


		fmt.Printf("i = %v, max_sum = %v, max_ending_here = %v\n", index, max_so_far, 
																		max_ending_here)
	}
	return max_ending_here
}





    func Max (x,y int)int{
    	if x< y {
    		return y
    	}
    	return x
    }