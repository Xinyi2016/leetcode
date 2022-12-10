// Kadane's Algorithm
// local_maximum at index i is the maximum of A[i] and the sum of A[i] and local_maximum at index i-1
function maxSubArray(nums: number[]): number {
  let local_max = 0;
  let global_max = nums[0];

  for (let i = 0; i < nums.length; i++) {
    local_max = Math.max(local_max + nums[i], nums[i]);
    global_max = Math.max(global_max, local_max);
  }
  return global_max;
}

export default maxSubArray;
