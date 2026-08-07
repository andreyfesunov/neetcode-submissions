type Solution struct{}

const DELIMITER_CHAR rune = 256
const DELIMITER = string(DELIMITER_CHAR)

func (s *Solution) Encode(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    var sb strings.Builder

    sb.WriteString(DELIMITER)

    for _, value := range strs {
        sb.WriteString(value)
        sb.WriteString(DELIMITER)
    }

    return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
    if encoded == "" {
        return []string{}
    }

    k := strings.Split(encoded, DELIMITER)

    return k[1:len(k)-1]
}
