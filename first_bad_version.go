package dsa

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool { return true }
func firstBadVersion(n int) int {
	start, end := 1, n
	fbv := -1
	for start <= end {
		mid := (start + end) / 2
		if isBadVersion(mid) {
			fbv = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return fbv
}
