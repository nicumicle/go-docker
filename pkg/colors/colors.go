package colors

func WriteRed(text string) string {
	return WriteCustom("\033[31m", text)
}

func WriteGreen(text string) string {
	return WriteCustom("\033[32m", text)
}

func WriteYellow(text string) string {
	return WriteCustom("\033[33m", text)
}

func WriteBlue(text string) string {
	return WriteCustom("\033[34m", text)
}

func WritePurple(text string) string {
	return WriteCustom("\033[35m", text)
}

func WriteCyan(text string) string {
	return WriteCustom("\033[36m", text)
}

func WriteGray(text string) string {
	return WriteCustom("\033[37m", text)
}

func WriteWhite(text string) string {
	return WriteCustom("\033[97m", text)
}

func WriteCustom(color string, text string) string {
	return color + text + colorReset()
}

func colorReset() string {
	return "\033[0m"
}
