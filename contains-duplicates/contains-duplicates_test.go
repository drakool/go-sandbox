package contains_duplicates

import "testing"


func TestDuplicates(t *testing.T){
   want:=true
   if containsNearbyDuplicate(nums []int{1,2,3,1}, 3) == false{
      t.Errorf('containsNearbyDuplicate returned false\n')
   }
}
