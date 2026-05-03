func firstUniqChar(s string) int {
mapChar := make(map[byte]int)
for i := 0; i < len(s); i++ {
    mapChar[s[i]]++
}
for i := 0; i < len(s); i++ {
    if mapChar[s[i]] == 1 {
        return i
    }
}