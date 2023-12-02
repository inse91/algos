package easy

import (
	"container/list"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"os"
	"slices"
	"testing"
)

func Test88(t *testing.T) {
	testCases := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{name: "1", nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, want: []int{1, 2, 2, 3, 5, 6}},
		{name: "2", nums1: []int{1}, m: 1, nums2: []int{}, n: 0, want: []int{1}},
		{name: "3", nums1: []int{0}, m: 0, nums2: []int{1}, n: 1, want: []int{1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nums1 := merge(tc.nums1, tc.m, tc.nums2, tc.n)
			assert.Equal(t, tc.want, nums1)
		})
	}

}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	firstPart := nums1[:m]
	secondPart := nums2[:n]
	nums1 = append(firstPart, secondPart...)

	slices.Sort(nums1)
	return nums1
}

func Test88V2(t *testing.T) {
	testCases := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{name: "1", nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, want: []int{1, 2, 2, 3, 5, 6}},
		{name: "2", nums1: []int{1}, m: 1, nums2: []int{}, n: 0, want: []int{1}},
		{name: "3", nums1: []int{0}, m: 0, nums2: []int{1}, n: 1, want: []int{1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nums1 := mergeV2(tc.nums1, tc.m, tc.nums2, tc.n)
			assert.Equal(t, tc.want, nums1)
		})
	}

}

func mergeV2(nums1 []int, m int, nums2 []int, n int) []int {

	res := make([]int, 0, len(nums1))

	firstPart := nums1[:m]
	secondPart := nums2[:n]

	for i1, i2, j := 0, 0, 0; i1 < len(firstPart) || i2 < len(secondPart); j++ {
		if i1 == len(firstPart) {
			res = append(res, secondPart[i2:]...)
			break
		}
		if i2 == len(secondPart) {
			res = append(res, firstPart[i1:]...)
			break
		}

		list.New()

		if firstPart[i1] <= secondPart[i2] {
			res = append(res, firstPart[i1])
			i1++
			continue
		}
		res = append(res, secondPart[i2])
		i2++
	}

	for i, v := range res {
		nums1[i] = v
	}

	return nums1
}

func TestNew(t *testing.T) {
	fmt.Println(123)

}

func OpenFile() (io.Reader, error) {

	file, err := os.Open("test_file.txt")
	if err != nil {
		return nil, err
	}

	defer func() {
		closeErr := file.Close()
		if err != nil && closeErr != nil {
			log.Println(closeErr)
		}
		if err != nil {
			return
		}
		err = closeErr
	}()
	return file, nil
}
