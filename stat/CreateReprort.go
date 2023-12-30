package main

//func createReportMain(detalization []string, connections []Connection) {
//
//	jsonData := createReport(detalization, connections)
//
//	err := writeJSONToFile(jsonData, "report.json")
//	if err != nil {
//		fmt.Println("Ошибка записи в файл:", err)
//		return
//	}
//}
//
//func writeJSONToFile(data interface{}, filename string) error {
//	file, err := os.Create(filename)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	encoder := json.NewEncoder(file)
//	encoder.SetIndent("", "  ")
//
//	if err := encoder.Encode(data); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func findURLByID(id int, connections []Connection) string {
//	for _, conn := range connections {
//		if conn.ID == id {
//			return conn.URL
//		}
//	}
//	return ""
//}
//
//func createReport(detalization []string, connections []Connection) map[string]interface{} {
//	report := make(map[string]interface{})
//
//	for _, connection := range connections {
//		if connection.PID == 0 {
//			continue
//		}
//
//		ip := connection.SourceIP
//		time := connection.Time[11:]
//		url := findURLByID(connection.PID, connections)
//
//		currLevel := report
//		for _, level := range detalization {
//
//			if level == "SourceIP" {
//				if _, ok := currLevel[ip]; !ok {
//					currLevel[ip] = make(map[string]interface{})
//					if _, ok := currLevel["Sum"]; !ok {
//						currLevel["Sum"] = 0
//					}
//				}
//				currLevel = currLevel[ip].(map[string]interface{})
//			} else if level == "TimeInterval" {
//				if _, ok := currLevel[time]; !ok {
//					currLevel[time] = make(map[string]interface{})
//					if _, ok := currLevel["Sum"]; !ok {
//						currLevel["Sum"] = 0
//					}
//				}
//				currLevel = currLevel[time].(map[string]interface{})
//			} else if level == "URL" {
//				if _, ok := currLevel[url]; !ok {
//					currLevel[url] = make(map[string]interface{})
//					if _, ok := currLevel["Sum"]; !ok {
//						currLevel["Sum"] = 0
//					}
//				}
//				currLevel = currLevel[url].(map[string]interface{})
//			}
//
//			if _, ok := currLevel["Sum"]; !ok {
//				currLevel["Sum"] = 0
//			}
//			currLevel["Sum"] = currLevel["Sum"].(int) + 1
//		}
//	}
//
//	delete(report, "Sum")
//
//	return report
//}
