package contains_duplicates


func containsNearbyDuplicate(nums []int, k int) bool {




	for i:=1; i < len(nums); i++ {
		
		for j:=i; j< k; j++ {

			if nums[i-1] == nums[j] {
				return true
			} 
		}
	}
	return false
    
}