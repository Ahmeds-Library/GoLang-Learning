package main
import (
	"fmt"
	"os"
	"time"
)
func main()  {
	start := time.Now()	
	defer func() {
		fmt.Println("Total execution time:", time.Since(start))
	}()
	
	content, err := os.ReadFile("../File.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := string(content)

	fmt.Println("Total characters:",len(data))
	countwords(data)
	countlines(data)
	countpunct(data)
	countspaces(data)
	countvowels(data)
	countdigits(data)
	countparagraph(data)
	countspchar(data)
	countconsonant(data)
}
//********************************<Count Words>****************************************
func countwords (a string) int {
	var words int = 0
	for i := 0; i < len(a); i++ {
		if a[i] == ' ' || a[i] == '\n' || a[i] == '\t'  {
			words++
		}
	}
	fmt.Println("Total words:", words)
	return words
}	

//********************************<Count Lines>****************************************
func countlines (a string) int {
	var lines int = 1
	for i := 0; i < len(a); i++ {
		if a[i] == '\n'{
			lines++
		}
	}
	fmt.Println("Total Lines:", lines)
	return lines
}

//********************************<Count Punctuations>****************************************
func countpunct (a string) int {
	var punc int = 0
	for i := 0; i < len(a); i++ {
		if a[i] == '.' || a[i] == ',' || a[i] == '!' || a[i] == '?' || a[i] == ';' || a[i] == ':' || a[i] == ')' || a[i] == '(' || a[i] == '-' || a[i] == '_' || a[i] == '"' || a[i] == '\''{
			punc++
		}
	}
	fmt.Println("Total Punctuations:", punc)
	return punc
}

//********************************<Count spaces>****************************************
func countspaces (a string) int {
	var spaces int = 0
	for i := 0; i < len(a); i++ {
		if a[i] == ' '{
			spaces++
		}
	}
	fmt.Println("Total Spaces:", spaces)
	return spaces
}

//********************************<Count Vowels>****************************************
func countvowels (a string) int {
	var vowels int = 0
	for i := 0; i < len(a); i++ {
		if a[i] == 'a' || a[i] == 'e' || a[i] == 'i' || a[i] == 'o' || a[i] == 'u' || a[i] == 'A' || a[i] == 'E' || a[i] == 'I' || a[i] == 'O' || a[i] == 'U' {
			vowels++
		}
	}
	fmt.Println("Total Vowels:", vowels)
	return vowels
}

//********************************<Count Digits>****************************************
func countdigits (a string) int {
	var digits int = 0
	for i := 0; i < len(a); i++ {
		if a[i] == '0' || a[i] == '1' || a[i] == '2' || a[i] == '3' || a[i] == '4' || a[i] == '5' || a[i] == '6' || a[i] == '7' || a[i] == '8' || a[i] == '9' {
			digits++
		}
	}
	fmt.Println("Total Digits:", digits)
	return digits
}	

//********************************<Count spaces>****************************************
func countparagraph (a string) int {
	var paragraph int = 1
	for i := 0; i < len(a); i++ {
		if a[i] == '\n' {
			paragraph++
		}
	}
	fmt.Println("Total Paragraph:", paragraph)
	return paragraph	
}

//********************************<Count Special characters>****************************************
func countspchar (a string) int {
	var spchar int = 0
	for i := 0; i < len(a); i++ {
		if a[i] == '@' || a[i] == '#' || a[i] == '$' || a[i] == '%' || a[i] == '^' || a[i] == '&' || a[i] == '*' || a[i] == '+' || a[i] == '=' || a[i] == '{' || a[i] == '}' || a[i] == '[' || a[i] == ']' {
			spchar++
		}
	}
	fmt.Println("Total Special Characters:", spchar)
	return spchar
}

//********************************<Count consonant>****************************************
func countconsonant (a string) int {
	var consonant int = 0
	for i := 0; i < len(a); i++ {
		if a[i] != 'a' && a[i] != 'e' && a[i] != 'i' && a[i] != 'o' && a[i] != 'u' && a[i] != 'A' && a[i] != 'E' && a[i] != 'I' && a[i] != 'O' && a[i] != 'U' && (a[i] >= 'a' && a[i] <= 'z' || a[i] >= 'A' && a[i] <= 'Z') {
			consonant++
		}
	}
	fmt.Println("Total Consonants:", consonant)
	return consonant
}