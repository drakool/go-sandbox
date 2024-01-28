package contains_duplicates

import "testing"


func TestDuplicates(t *testing.T){
   want:=true
      //intSlice := []int{5, 2, 6, 3, 1, 4}

   arr := []int{1,2,3,1}
   if containsNearbyDuplicate(arr , 3) == want{
      t.Errorf("containsNearbyDuplicate returned false\n")
   }
}
