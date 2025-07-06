package utils

// prompt texts
const (
	PromptTargetURL string = BoldText + CyanText + " target URL: " + DefaultText
	PromptOption    string = BoldText + CyanText + " option: " + DefaultText
)

// Prompt prefixes
const (
	IndicatorPending string = BoldText + RedText + "[+]" + DefaultText
	IndicatorSuccess string = BoldText + GreenText + "[✓]" + DefaultText

	IndicatorReachable   string = BoldText + GreenText + "[●]" + DefaultText
	IndicatorUnreachable string = BoldText + RedText + "[●]" + DefaultText
)
