package editdistance

import (
	"fmt"
	"testing"
)

/*
https://leetcode.com/problems/edit-distance/
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:

Insert a character
Delete a character
Replace a character


Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')


Constraints:

0 <= word1.length, word2.length <= 500
word1 and word2 consist of lowercase English letters.

*/

func minDistance(word1 string, word2 string) int {
	lw1 := len(word1)
	lw2 := len(word2)

	if lw1 == lw2 {

	} else if lw1 > lw2 {

	} else { //lw1 < lw2

	}
	return 0
}

func insert(w string, position int, c string, min int) (string, int) {
	return w[:position] + c + w[position:], min + 1
}

func replace(w string, position int, c string, min int) (string, int) {
	return w[:position-1] + c + w[position:], min + 1
}

func delete(w string, position int, min int) (string, int) {
	return w[:position-1] + w[position:], min + 1
}

func TestA(t *testing.T) {
	a := "hello"
	fmt.Println(a[:2]) //he
	fmt.Println(a[2:]) //llo
}

func TestMinDistanct(t *testing.T) {
	min := minDistance("horse", "ros")
	fmt.Println(min)
}

/*
inten tion
execu tion


horse
ros

rorse



หาว่า length เท่ากันไหม อย่างน้อยก็จะได้ min เท่า length

ถ้า word 2 น้อยกว่า word 1 ก็ ไล่ดูว่า เรียงแล้วรึยัง

case 1 ไม่มีตัวที่ตรงกันเลย

case 2 มีตัวที่ตรงกัน
*/
