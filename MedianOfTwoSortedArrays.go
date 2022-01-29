/*
  Runtime: 8ms faster than 97% of submissions
  Memory Usage: 5.2 MB less than 99.83% of submissions 
*/

func solution(nums1 []int, nums2 []int) float64 {
	var i, j, k, a, b int
	n1 := len(nums1)
	n2 := len(nums2)
	center := ((n1 + n2) / 2) + 1

	for i < n1 && j < n2 && k < center {
		if nums1[i] < nums2[j] {
			//res[k] = nums1[i]
			b = a
			a = nums1[i]
			k++
			i++
		} else {
			//res[k] = nums2[j]
			b = a
			a = nums2[j]
			k++
			j++
		}
	}

	for i < n1 && k < center {
		//res[k] = nums1[i]
		b = a
		a = nums1[i]
		k++
		i++
	}

	for j < n2 && k < center {
		//res[k] = nums2[j]
		b = a
		a = nums2[j]
		k++
		j++
	}

	if (n1+n2)%2 == 0 {
		return float64(a+b) / float64(2)
	} else {
		return float64(a)
	}
}
