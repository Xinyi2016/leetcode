/**
 Do not return anything, modify nums1 in-place instead.
 */
 function merge(nums1: number[], m: number, nums2: number[], n: number): void {
    for (let i=m+n-1; i>=0;i--) {
        if (m == 0) {
            nums1[i] = nums2[n-1]
            n--
            continue
        }

        if (n == 0) {
            nums1[i] = nums1[m-1]
            m--
            continue
        }

        if (nums1[m-1] > nums2[n-1]) {
            nums1[i] = nums1[m-1]
            m--
        } else {
            nums1[i] = nums2[n-1]
            n--
        }

    }


}