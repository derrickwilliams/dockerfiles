package ui

import (
	"strings"
)

const BORDER_CHAR string = ""
const CHARS_PER_SECTION int = 4

func TopBorder(horizontalLen int) {
	return strings.Times(BORDER_CHAR, (CHARS_PER_SECTION * 4) + );
}