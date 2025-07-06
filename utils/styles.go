package utils

// prompt texts
const (
	PromptTargetURL string = BoldText + CyanText + " target URL: " + DefaultText
)

// Prompt prefixes
const (
	IndicatorSuccess string = BoldText + GreenText + "[✓]" + DefaultText
	IndicatorFailure string = BoldText + RedText + "[x]" + DefaultText

	IndicatorReachable   string = BoldText + GreenText + "●" + DefaultText
	IndicatorUnreachable string = BoldText + RedText + "●" + DefaultText
)

// styles
const (
	DefaultText = "\033[0m"
	BoldText    = "\033[1m"
	ClearLine   = "\033[1A\033[2K\r"
)

// Colors
const (
	BlackText   = "\033[30m"
	RedText     = "\033[31m"
	GreenText   = "\033[32m"
	YellowText  = "\033[33m"
	BlueText    = "\033[34m"
	MagentaText = "\033[35m"
	CyanText    = "\033[36m"
	WhiteText   = "\033[37m"
)

// counter
const (
	ProgressBarPrefix = "[00:00:00] "
	ProgressBarFilled = "█"
	ProgressBarEmpty  = "░"
)
