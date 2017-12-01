package number2text

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var ErrNumberTooBig = errors.New("internal_error.number_too_big")

func Int2Chinese(i int64) (string, error) {

	CHINESE_NEGATIVE := "负"

	ui := int64(math.Abs(float64(i)))
	words, err := uInt2Chinese(ui)

	if err != nil {
		return "", err
	}

	if i < 0 {
		words = CHINESE_NEGATIVE + words
	}

	return words, nil
}

func uInt2Chinese(i int64) (string, error) {

	CHINESE_ZERO := "零"
	CHINESE_DIGITS := []string{"", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	CHINESE_UNITS := []string{"", "十", "百", "千"}
	CHINESE_GROUP_UNITS := []string{"", "万", "亿", "兆"}

	words := []string{}
	groupIsZero := true
	needZero := false

	pi := i

	for _, e := range enumerateDigit(pi) {
		position := e["position"]
		digit := e["digit"]

		unit := position % len(CHINESE_UNITS)
		group := int(math.Floor(float64(position) / float64(len(CHINESE_UNITS))))

		if digit != 0 {
			if needZero {
				words = append(words, CHINESE_ZERO)
			}
			if digit != 1 || unit != 1 || !groupIsZero || (group == 0 && needZero) {
				words = append(words, CHINESE_DIGITS[digit])
			}
			if unit > len(CHINESE_UNITS)-1 {
				return "", ErrNumberTooBig
			}
			words = append(words, CHINESE_UNITS[unit])
		}

		groupIsZero = groupIsZero && digit == 0

		if unit == 0 && !groupIsZero {
			if group > len(CHINESE_GROUP_UNITS)-1 {
				return "", ErrNumberTooBig
			}
			words = append(words, CHINESE_GROUP_UNITS[group])
		}

		needZero = digit == 0 && (unit != 0 || groupIsZero)

		if unit == 0 {
			groupIsZero = true
		}
	}

	if pi == 0 {
		words = []string{CHINESE_ZERO}
	}

	return strings.Join(words, ""), nil

}

func Float2Chinese(f float64, n int) (string, error) {

	CHINESE_NEGATIVE := "负"

	i := n
	for ; i >= 0; i-- {
		sf := strconv.FormatFloat(f, 'f', i, 64)
		if !strings.HasSuffix(sf, "0") {
			break
		}
	}

	roundF := math.Floor(f)
	if f-roundF > 0.999999 {
		roundF = math.Ceil(f)
	}
	amplifiedF := int64(f*math.Pow10(i) + 0.5)
	amplifiedI := int64(roundF * math.Pow10(i))
	decimalF := float64(amplifiedF-amplifiedI) / math.Pow10(i)
	intF := int64(roundF)

	if f < 0.0 {
		roundF = math.Ceil(f)
		if roundF-f > 0.999999 {
			roundF = math.Floor(f)
		}
		amplifiedF = int64(f*math.Pow10(i) - 0.5)
		amplifiedI = int64(roundF * math.Pow10(i))
		decimalF = float64(amplifiedI-amplifiedF) / math.Pow10(i)
		intF = int64(roundF * -1)
	}

	intWords, err := uInt2Chinese(intF)
	if err != nil {
		return "", err
	}

	if f < 0.0 {
		intWords = CHINESE_NEGATIVE + intWords
	}

	decimalWords, err := decimal2Chinese(decimalF, i)
	if err != nil {
		return "", err
	}
	if decimalWords == "" {
		return intWords, nil
	}

	return intWords + "点" + decimalWords, nil
}

func enumerateDigit(i int64) []map[string]int {
	enum := []map[string]int{}
	position := 0
	for i > 0 {
		digit := int(i % 10)
		i = int64(math.Floor(float64(i) / 10))
		enum = append(enum, map[string]int{"position": position, "digit": digit})
		position += 1
	}

	reversedEnum := []map[string]int{}
	for i := len(enum) - 1; i >= 0; i-- {
		reversedEnum = append(reversedEnum, enum[i])
	}
	return reversedEnum
}

func decimal2Chinese(f float64, i int) (string, error) {

	CHINESE_DIGIT_MAP := map[string]string{
		"0": "零", "1": "一", "2": "二", "3": "三", "4": "四", "5": "五", "6": "六", "7": "七", "8": "八", "9": "九",
	}

	words := []string{}

	sf := strconv.FormatFloat(f, 'f', i, 64)
	if len(sf) <= 2 {
		return "", nil
	}
	sd := sf[2:]
	for _, d := range sd {
		words = append(words, CHINESE_DIGIT_MAP[string(d)])
	}
	return strings.Join(words, ""), nil
}
