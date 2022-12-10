function containsDuplicate(nums: number[]): boolean {
  const seen = new Set<number>();
  for (const x of nums) {
    if (seen.has(x)) return true;
    else seen.add(x);
  }

  return false;
}

export default containsDuplicate;
