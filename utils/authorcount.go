package utils

func AuthorCounter(currentAuthors *[]string, newAuthors BookcountData, c chan uint32) {

	var uniqueamount uint32 = 0

	//Converting BookcountData.Results into list of strings

	for i := 0; i < len(newAuthors.Results); i++ {
		for j := 0; j < len(newAuthors.Results[i].Authors); j++ {

			currentAuthor := newAuthors.Results[i].Authors[j].Name
			duplicate := false

			//For each value of newdata, we compare for all the values of currentData
			for k := 0; k < len(*currentAuthors); k++ {
				if (*currentAuthors)[k] == currentAuthor {
					duplicate = true
				}
			}

			if !duplicate {
				/*
					If unique (does not exists any duplicates), then we append this value to
					currentData and increase our uniqueamount
				*/
				*currentAuthors = append(*currentAuthors, currentAuthor)
				uniqueamount++
			}
		}
	}

	//At the end, we put this uniqueamount into c: c <- uniqueamount
	c <- uniqueamount
}
