package LL1

// Todo: write a real first calculator later

//func first(ter map[string]bool,
//	nTer map[string]bool,
//	grams map[string][][]string) (firstSet map[string][]string) {
//
//	for key, _ := range nTer {
//		firstSet[key] = []string{}
//		list := grams[key]
//		firstIndx := 0
//		if len(list) == 0 {
//			nTer[key] = false
//		}
//
//		for i, v := range list {
//			if _, ok := ter[v[firstIndx]]; ok {
//				firstSet[key] = append(firstSet[key], v[firstIndx])
//				list[i] = list[len(list)-1]
//				grams[key] = list[:len(list)-1]
//			} else {
//				if !nTer[v[firstIndx]] {
//					firstSet[key] = append(firstSet[key], firstSet[v[firstIndx]]...)
//
//				}
//			}
//		}
//	}
//
//	return firstSet
//}
