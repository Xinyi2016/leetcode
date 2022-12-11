function twoSum(nums: number[], target: number): number[] {
    const res: Map<number, number> = new Map();

    for (let i=0;i<nums.length;i++) {
        const mis = res.get(nums[i])
        if (mis != undefined) {
            return [mis, i];
        }
        res.set(target-nums[i], i);
    }

    return [];

}

export default twoSum;