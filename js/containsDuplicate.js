// LeetCode - 217. Contains Duplicate (https://leetcode.com/problems/contains-duplicate/)

var containsDuplicate = function(nums) {
  var setNumbers = new Set()
  for(let i = 0; i<= nums.length; i++){
      if(setNumbers.size == 0){
          setNumbers.add(nums[i])
          continue
      }

      if (setNumbers.has(nums[i])){
          return true
      }

      setNumbers.add(nums[i])
  }

  return false
};

console.log(containsDuplicate(nums = [1,2,3,1]))