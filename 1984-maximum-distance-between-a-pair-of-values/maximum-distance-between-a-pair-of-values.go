func maxDistance(nums1 []int, nums2 []int) int {
    maxDist := 0
    i, j := 0, 0
    
    // Iterate through both arrays
    for i < len(nums1) && j < len(nums2) {
        // Check if the condition is met
        if nums1[i] <= nums2[j] {
            // Update max distance if current distance is larger
            if j - i > maxDist {
                maxDist = j - i
            }
            // Try to increase the distance by moving j forward
            j++
        } else {
            // If nums1[i] > nums2[j], we need a smaller nums1[i]
            // Move i forward to potentially meet the condition
            i++
            // Optimization: Since i must be <= j, 
            // ensure j doesn't lag behind i
            if i > j {
                j = i
            }
        }
    }
    
    return maxDist
}