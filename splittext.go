package gofpdf

//	"strings"

// SplitText splits UTF-8 encoded text into several lines using the current
// font. Each line has its length limited to a maximum width given by w. This
// function can be used to determine the total height of wrapped text for
// vertical placement purposes.
func (f *Fpdf) SplitText(txt string, maxWidth float64) []string {
	var lines []string
	var currentLine string
	words := []rune(txt)                  // 将文本转换为 UTF-8 字符切片
	maxWidth = maxWidth - (2 * f.cMargin) // 扣除左右边距

	for _, char := range words {
		charStr := string(char)

		if char == '\n' {
			// 遇到换行符，保存当前行并换行
			lines = append(lines, currentLine)
			currentLine = "" // 开启新的一行
			continue
		}

		lineWidth := f.GetStringWidth(currentLine + charStr)
		if lineWidth > maxWidth {
			// 当前行已满，换行
			lines = append(lines, currentLine)
			currentLine = charStr // 将当前字符作为新行的起点
		} else {
			// 当前行未满，继续添加字符
			currentLine += charStr
		}
	}

	// 添加最后一行
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

// func (f *Fpdf) SplitText(txt string, w float64) (lines []string) {
// 	cw := f.currentFont.Cw
// 	// fmt.Printf("宽度 - 'A': %d, '中': %d, ' ': %d\n", cw['A'], cw['中'], cw[' '])
// 	// Ceil 改成 Floor 呢
// 	wmax := int(math.Floor((w - 2*f.cMargin) * 1000 / f.fontSize))
// 	s := []rune(txt)
// 	fmt.Println("s", txt, s)
// 	nb := len(s)
// 	for nb > 0 && s[nb-1] == '\n' {
// 		nb--
// 	}
// 	s = s[0:nb]
// 	sep := -1
// 	i := 0
// 	j := 0
// 	l := 0
// 	for i < nb {
// 		fmt.Println(c, cw[c])
// 		c := s[i]
// 		l += cw[c]

// 		// 更新分割点
// 		if unicode.IsSpace(c) || isChinese(c) {
// 			sep = i
// 		}

// 		// 判断是否需要分割
// 		if c == '\n' || l > wmax {
// 			if sep == -1 {
// 				sep = i
// 			}
// 			lines = append(lines, string(s[j:sep+1]))
// 			i = sep + 1
// 			sep = -1
// 			j = i
// 			l = 0
// 		} else {
// 			i++
// 		}
// 	}
// 	if i != j {
// 		lines = append(lines, string(s[j:i]))
// 	}
// 	return lines
// }

// func (f *Fpdf) SplitText(txt string, w float64) (lines []string) {
// 	cw := f.currentFont.Cw
// 	wmax := int(math.Ceil((w - 2*f.cMargin) * 1000 / f.fontSize))
// 	s := []rune(txt)
// 	nb := len(s)
// 	for nb > 0 && s[nb-1] == '\n' {
// 		nb--
// 	}
// 	s = s[0:nb]
// 	sep := -1
// 	i := 0
// 	j := 0
// 	l := 0
// 	for i < nb {
// 		c := s[i]
// 		l += cw[c]

// 		// 更新分割点
// 		if unicode.IsSpace(c) || isChinese(c) {
// 			sep = i
// 		}

// 		// 判断是否需要分割
// 		if c == '\n' || l > wmax {
// 			if sep == -1 {
// 				sep = i
// 			}
// 			// 修复分割范围，确保字符完整
// 			lines = append(lines, string(s[j:sep+1]))
// 			i = sep + 1
// 			sep = -1
// 			j = i
// 			l = 0
// 		} else {
// 			i++
// 		}
// 	}
// 	if i != j {
// 		lines = append(lines, string(s[j:i]))
// 	}
// 	return lines
// }

// func (f *Fpdf) SplitText(txt string, w float64) (lines []string) {
// 	cw := f.currentFont.Cw
// 	wmax := int(math.Ceil((w - 2*f.cMargin) * 1000 / f.fontSize))
// 	s := []rune(txt) // Return slice of UTF-8 runes
// 	nb := len(s)
// 	for nb > 0 && s[nb-1] == '\n' {
// 		nb--
// 	}
// 	s = s[0:nb]
// 	sep := -1
// 	i := 0
// 	j := 0
// 	l := 0
// 	for i < nb {
// 		c := s[i]
// 		l += cw[c]
// 		if unicode.IsSpace(c) || isChinese(c) {
// 			sep = i
// 		}
// 		if c == '\n' || l > wmax {
// 			if sep == -1 {
// 				if i == j {
// 					i++
// 				}
// 				sep = i
// 			} else {
// 				i = sep + 1
// 			}
// 			lines = append(lines, string(s[j:sep]))
// 			sep = -1
// 			j = i
// 			l = 0
// 		} else {
// 			i++
// 		}
// 	}
// 	if i != j {
// 		lines = append(lines, string(s[j:i]))
// 	}
// 	return lines
// }
