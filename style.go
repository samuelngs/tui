package tui

// Position type
type Position string

// Positions
const (
	PositionRelative Position = "relative"
	PositionAbsolute          = "absolute"
)

// FlexDirection type
type FlexDirection string

// FlexDirections
const (
	FlexColumn FlexDirection = "column"
	FlexRow                  = "row"
)

// FlexPosition type
type FlexPosition string

// FlexPositions
const (
	FlexStart  FlexPosition = "flex-start"
	FlexEnd                 = "flex-end"
	FlexCenter              = "center"
)

// Style represents the component style
type Style struct {

	// flexbox styles
	Flex           int
	FlexDirection  FlexDirection
	JustifyContent FlexPosition

	// sizing styles
	Height, Width, MaxWidth, MaxHeight, MinWidth, MinHeight int

	// text styles
	FontSize, FontWeight int
	Color                string

	// margins
	Margin, MarginTop, MarginLeft, MarginBottom, MarginRight, MarginHorizontal, MarginVertical int

	// paddings
	Padding, PaddingTop, PaddingLeft, PaddingBottom, PaddingRight, PaddingHorizontal, PaddingVertical int

	// positions
	Position                 Position
	Top, Left, Bottom, Right int

	// borders
	BorderTopWidth, BorderLeftWidth, BorderBottomWidth, BorderRightWidth int
	BorderTopColor, BorderLeftColor, BorderBottomColor, BorderRightColor string
	BorderTopStyle, BorderLeftStyle, BorderBottomStyle, BorderRightStyle string
}

// GetWindowSize returns the window bound
func (v *Style) GetWindowSize() *Bound {
	return getWinSize()
}

// GetWindowHeight to returns the window height
func (v *Style) GetWindowHeight() uint {
	return v.GetWindowSize().Height
}

// GetWindowWidth to returns the window height
func (v *Style) GetWindowWidth() uint {
	return v.GetWindowSize().Width
}
