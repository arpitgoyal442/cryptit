package encrypt


// A--> D    B-->E

func Nimbus(s string)string{


	ans:=""

	for _,c:=range s{

		ascii:=int(c)

		chr:=string(ascii+3)
		ans+=chr
	}



	return ans
	

}