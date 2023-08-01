package piscine

func PodiumPosition(podium [][]string) [][]string {
	n := len(podium)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		temp := podium[i]
		podium[i] = podium[j]
		podium[j] = temp
	}
	return podium
}
