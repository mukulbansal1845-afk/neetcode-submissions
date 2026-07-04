func maxProfit(arr []int) int {

	mn := arr[0]
	pro := 0

	n := len(arr)

	for i:=0;i<n;i++{

		if arr[i]<mn{
			mn = arr[i]
		}

		diff := arr[i] - mn

		if diff>pro {
			pro = diff
		}
	}

	return pro



}
