package main

// Сколько существует способов задать переменную типа slice или map?


// slice: (слайс под капотом ссылка на underlying массив)
// 1. var s []string
// 2. var s = []string{"Go", "Python", "Java"}
// 3. s := []string{}
// 5. s := make([]string, 0)
// 5. s1 := s2[:]

// map: (мапа под капотом ссылка на хеш таблицу, которая равна nil, и для работы её работы нужна инициализациия)
// 1. var m = map[string]int{}
//    m["Alex"] = 18
// 2. m := map[string]int{}
//	  m["Alex"] = 18
// 3. m := make(map[string]int)
