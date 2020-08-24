package dest_city

func destCity(paths [][]string) string {
	result := ""
	startMap := make(map[string]bool)
	for _, path := range paths {
		startMap[path[0]] = true
	}
	for _, path := range paths {
		if ok, _ := startMap[path[1]]; ok == false {
			return path[1]
		}
	}
	return result
}
