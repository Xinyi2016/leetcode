function dominantIndex(nums: number[]): number {
  let max = 0;
  let maxIndex = 0;
  let isTwice = false;

  nums.forEach((n, i) => {
    if (n > max) {
      if (n >= max * 2) {
        isTwice = true;
      } else {
        isTwice = false;
      }
      max = n
      maxIndex = i;
    } else {
      if (max < n * 2) {
        isTwice = false;
      }
    }
  })

  return isTwice ? maxIndex : -1;
}

export default dominantIndex;
