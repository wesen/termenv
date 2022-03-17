package termenv

import (
	"fmt"
	"strings"
)

// Sequence definitions.
const (
	// Cursor positioning.
	CursorUpSeq              = "%dA"
	CursorDownSeq            = "%dB"
	CursorForwardSeq         = "%dC"
	CursorBackSeq            = "%dD"
	CursorNextLineSeq        = "%dE"
	CursorPreviousLineSeq    = "%dF"
	CursorHorizontalSeq      = "%dG"
	CursorPositionSeq        = "%d;%dH"
	EraseDisplaySeq          = "%dJ"
	EraseLineSeq             = "%dK"
	ScrollUpSeq              = "%dS"
	ScrollDownSeq            = "%dT"
	SaveCursorPositionSeq    = "s"
	RestoreCursorPositionSeq = "u"
	ChangeScrollingRegionSeq = "%d;%dr"
	InsertLineSeq            = "%dL"
	DeleteLineSeq            = "%dM"

	// Explicit values for EraseLineSeq.
	EraseLineRightSeq  = "0K"
	EraseLineLeftSeq   = "1K"
	EraseEntireLineSeq = "2K"

	// Mouse.
	EnableMousePressSeq       = "?9h" // press only (X10)
	DisableMousePressSeq      = "?9l"
	EnableMouseSeq            = "?1000h" // press, release, wheel
	DisableMouseSeq           = "?1000l"
	EnableMouseHiliteSeq      = "?1001h" // highlight
	DisableMouseHiliteSeq     = "?1001l"
	EnableMouseCellMotionSeq  = "?1002h" // press, release, move on pressed, wheel
	DisableMouseCellMotionSeq = "?1002l"
	EnableMouseAllMotionSeq   = "?1003h" // press, release, move, wheel
	DisableMouseAllMotionSeq  = "?1003l"

	// Screen.
	RestoreScreenSeq = "?47l"
	SaveScreenSeq    = "?47h"
	AltScreenSeq     = "?1049h"
	ExitAltScreenSeq = "?1049l"

	// Session.
	SetWindowTitleSeq     = "2;%s\007"
	SetForegroundColorSeq = "10;%s\007"
	SetBackgroundColorSeq = "11;%s\007"
	SetCursorColorSeq     = "12;%s\007"
	ShowCursorSeq         = "?25h"
	HideCursorSeq         = "?25l"
)

// Reset the terminal to its default style, removing any active styles.
func (o Output) Reset() {
	fmt.Fprint(o.tty, CSI+ResetSeq+"m")
}

// SetForegroundColor sets the default foreground color.
func (o Output) SetForegroundColor(color Color) {
	fmt.Fprintf(o.tty, OSC+SetForegroundColorSeq, color)
}

// SetBackgroundColor sets the default background color.
func (o Output) SetBackgroundColor(color Color) {
	fmt.Fprintf(o.tty, OSC+SetBackgroundColorSeq, color)
}

// SetCursorColor sets the cursor color.
func (o Output) SetCursorColor(color Color) {
	fmt.Fprintf(o.tty, OSC+SetCursorColorSeq, color)
}

// RestoreScreen restores a previously saved screen state.
func (o Output) RestoreScreen() {
	fmt.Fprint(o.tty, CSI+RestoreScreenSeq)
}

// SaveScreen saves the screen state.
func (o Output) SaveScreen() {
	fmt.Fprint(o.tty, CSI+SaveScreenSeq)
}

// AltScreen switches to the alternate screen buffer. The former view can be
// restored with ExitAltScreen().
func (o Output) AltScreen() {
	fmt.Fprint(o.tty, CSI+AltScreenSeq)
}

// ExitAltScreen exits the alternate screen buffer and returns to the former
// terminal view.
func (o Output) ExitAltScreen() {
	fmt.Fprint(o.tty, CSI+ExitAltScreenSeq)
}

// ClearScreen clears the visible portion of the terminal.
func (o Output) ClearScreen() {
	fmt.Fprintf(o.tty, CSI+EraseDisplaySeq, 2)
	o.MoveCursor(1, 1)
}

// MoveCursor moves the cursor to a given position.
func (o Output) MoveCursor(row int, column int) {
	fmt.Fprintf(o.tty, CSI+CursorPositionSeq, row, column)
}

// HideCursor hides the cursor.
func (o Output) HideCursor() {
	fmt.Fprintf(o.tty, CSI+HideCursorSeq)
}

// ShowCursor shows the cursor.
func (o Output) ShowCursor() {
	fmt.Fprintf(o.tty, CSI+ShowCursorSeq)
}

// SaveCursorPosition saves the cursor position.
func (o Output) SaveCursorPosition() {
	fmt.Fprint(o.tty, CSI+SaveCursorPositionSeq)
}

// RestoreCursorPosition restores a saved cursor position.
func (o Output) RestoreCursorPosition() {
	fmt.Fprint(o.tty, CSI+RestoreCursorPositionSeq)
}

// CursorUp moves the cursor up a given number of lines.
func (o Output) CursorUp(n int) {
	fmt.Fprintf(o.tty, CSI+CursorUpSeq, n)
}

// CursorDown moves the cursor down a given number of lines.
func (o Output) CursorDown(n int) {
	fmt.Fprintf(o.tty, CSI+CursorDownSeq, n)
}

// CursorForward moves the cursor up a given number of lines.
func (o Output) CursorForward(n int) {
	fmt.Fprintf(o.tty, CSI+CursorForwardSeq, n)
}

// CursorBack moves the cursor backwards a given number of cells.
func (o Output) CursorBack(n int) {
	fmt.Fprintf(o.tty, CSI+CursorBackSeq, n)
}

// CursorNextLine moves the cursor down a given number of lines and places it at
// the beginning of the line.
func (o Output) CursorNextLine(n int) {
	fmt.Fprintf(o.tty, CSI+CursorNextLineSeq, n)
}

// CursorPrevLine moves the cursor up a given number of lines and places it at
// the beginning of the line.
func (o Output) CursorPrevLine(n int) {
	fmt.Fprintf(o.tty, CSI+CursorPreviousLineSeq, n)
}

// ClearLine clears the current line.
func (o Output) ClearLine() {
	fmt.Fprint(o.tty, CSI+EraseEntireLineSeq)
}

// ClearLineLeft clears the line to the left of the cursor.
func (o Output) ClearLineLeft() {
	fmt.Fprint(o.tty, CSI+EraseLineLeftSeq)
}

// ClearLineRight clears the line to the right of the cursor.
func (o Output) ClearLineRight() {
	fmt.Fprint(o.tty, CSI+EraseLineRightSeq)
}

// ClearLines clears a given number of lines.
func (o Output) ClearLines(n int) {
	clearLine := fmt.Sprintf(CSI+EraseLineSeq, 2)
	cursorUp := fmt.Sprintf(CSI+CursorUpSeq, 1)
	fmt.Fprint(o.tty, clearLine+strings.Repeat(cursorUp+clearLine, n))
}

// ChangeScrollingRegion sets the scrolling region of the terminal.
func (o Output) ChangeScrollingRegion(top, bottom int) {
	fmt.Fprintf(o.tty, CSI+ChangeScrollingRegionSeq, top, bottom)
}

// InsertLines inserts the given number of lines at the top of the scrollable
// region, pushing lines below down.
func (o Output) InsertLines(n int) {
	fmt.Fprintf(o.tty, CSI+InsertLineSeq, n)
}

// DeleteLines deletes the given number of lines, pulling any lines in
// the scrollable region below up.
func (o Output) DeleteLines(n int) {
	fmt.Fprintf(o.tty, CSI+DeleteLineSeq, n)
}

// EnableMousePress enables X10 mouse mode. Button press events are sent only.
func (o Output) EnableMousePress() {
	fmt.Fprint(o.tty, CSI+EnableMousePressSeq)
}

// DisableMousePress disables X10 mouse mode.
func (o Output) DisableMousePress() {
	fmt.Fprint(o.tty, CSI+DisableMousePressSeq)
}

// EnableMouse enables Mouse Tracking mode.
func (o Output) EnableMouse() {
	fmt.Fprint(o.tty, CSI+EnableMouseSeq)
}

// DisableMouse disables Mouse Tracking mode.
func (o Output) DisableMouse() {
	fmt.Fprint(o.tty, CSI+DisableMouseSeq)
}

// EnableMouseHilite enables Hilite Mouse Tracking mode.
func (o Output) EnableMouseHilite() {
	fmt.Fprint(o.tty, CSI+EnableMouseHiliteSeq)
}

// DisableMouseHilite disables Hilite Mouse Tracking mode.
func (o Output) DisableMouseHilite() {
	fmt.Fprint(o.tty, CSI+DisableMouseHiliteSeq)
}

// EnableMouseCellMotion enables Cell Motion Mouse Tracking mode.
func (o Output) EnableMouseCellMotion() {
	fmt.Fprint(o.tty, CSI+EnableMouseCellMotionSeq)
}

// DisableMouseCellMotion disables Cell Motion Mouse Tracking mode.
func (o Output) DisableMouseCellMotion() {
	fmt.Fprint(o.tty, CSI+DisableMouseCellMotionSeq)
}

// EnableMouseAllMotion enables All Motion Mouse mode.
func (o Output) EnableMouseAllMotion() {
	fmt.Fprint(o.tty, CSI+EnableMouseAllMotionSeq)
}

// DisableMouseAllMotion disables All Motion Mouse mode.
func (o Output) DisableMouseAllMotion() {
	fmt.Fprint(o.tty, CSI+DisableMouseAllMotionSeq)
}

// SetWindowTitle sets the terminal window title.
func (o Output) SetWindowTitle(title string) {
	fmt.Fprintf(o.tty, OSC+SetWindowTitleSeq, title)
}

// Legacy functions.

// Reset the terminal to its default style, removing any active styles.
func Reset() {
	NewOutputWithProfile(defaultOutputFile, ANSI).Reset()
}

// SetForegroundColor sets the default foreground color.
func SetForegroundColor(color Color) {
	NewOutputWithProfile(defaultOutputFile, ANSI).SetForegroundColor(color)
}

// SetBackgroundColor sets the default background color.
func SetBackgroundColor(color Color) {
	NewOutputWithProfile(defaultOutputFile, ANSI).SetBackgroundColor(color)
}

// SetCursorColor sets the cursor color.
func SetCursorColor(color Color) {
	NewOutputWithProfile(defaultOutputFile, ANSI).SetCursorColor(color)
}

// RestoreScreen restores a previously saved screen state.
func RestoreScreen() {
	NewOutputWithProfile(defaultOutputFile, ANSI).RestoreScreen()
}

// SaveScreen saves the screen state.
func SaveScreen() {
	NewOutputWithProfile(defaultOutputFile, ANSI).SaveScreen()
}

// AltScreen switches to the alternate screen buffer. The former view can be
// restored with ExitAltScreen().
func AltScreen() {
	NewOutputWithProfile(defaultOutputFile, ANSI).AltScreen()
}

// ExitAltScreen exits the alternate screen buffer and returns to the former
// terminal view.
func ExitAltScreen() {
	NewOutputWithProfile(defaultOutputFile, ANSI).ExitAltScreen()
}

// ClearScreen clears the visible portion of the terminal.
func ClearScreen() {
	NewOutputWithProfile(defaultOutputFile, ANSI).ClearScreen()
}

// MoveCursor moves the cursor to a given position.
func MoveCursor(row int, column int) {
	NewOutputWithProfile(defaultOutputFile, ANSI).MoveCursor(row, column)
}

// HideCursor hides the cursor.
func HideCursor() {
	NewOutputWithProfile(defaultOutputFile, ANSI).HideCursor()
}

// ShowCursor shows the cursor.
func ShowCursor() {
	NewOutputWithProfile(defaultOutputFile, ANSI).ShowCursor()
}

// SaveCursorPosition saves the cursor position.
func SaveCursorPosition() {
	NewOutputWithProfile(defaultOutputFile, ANSI).SaveCursorPosition()
}

// RestoreCursorPosition restores a saved cursor position.
func RestoreCursorPosition() {
	NewOutputWithProfile(defaultOutputFile, ANSI).RestoreCursorPosition()
}

// CursorUp moves the cursor up a given number of lines.
func CursorUp(n int) {
	NewOutputWithProfile(defaultOutputFile, ANSI).CursorUp(n)
}

// CursorDown moves the cursor down a given number of lines.
func CursorDown(n int) {
	NewOutputWithProfile(defaultOutputFile, ANSI).CursorDown(n)
}

// CursorForward moves the cursor up a given number of lines.
func CursorForward(n int) {
	NewOutputWithProfile(defaultOutputFile, ANSI).CursorForward(n)
}

// CursorBack moves the cursor backwards a given number of cells.
func CursorBack(n int) {
	NewOutputWithProfile(defaultOutputFile, ANSI).CursorBack(n)
}

// CursorNextLine moves the cursor down a given number of lines and places it at
// the beginning of the line.
func CursorNextLine(n int) {
	NewOutputWithProfile(defaultOutputFile, ANSI).CursorNextLine(n)
}

// CursorPrevLine moves the cursor up a given number of lines and places it at
// the beginning of the line.
func CursorPrevLine(n int) {
	NewOutputWithProfile(defaultOutputFile, ANSI).CursorPrevLine(n)
}

// ClearLine clears the current line.
func ClearLine() {
	NewOutputWithProfile(defaultOutputFile, ANSI).ClearLine()
}

// ClearLineLeft clears the line to the left of the cursor.
func ClearLineLeft() {
	NewOutputWithProfile(defaultOutputFile, ANSI).ClearLineLeft()
}

// ClearLineRight clears the line to the right of the cursor.
func ClearLineRight() {
	NewOutputWithProfile(defaultOutputFile, ANSI).ClearLineRight()
}

// ClearLines clears a given number of lines.
func ClearLines(n int) {
	NewOutputWithProfile(defaultOutputFile, ANSI).ClearLines(n)
}

// ChangeScrollingRegion sets the scrolling region of the terminal.
func ChangeScrollingRegion(top, bottom int) {
	NewOutputWithProfile(defaultOutputFile, ANSI).ChangeScrollingRegion(top, bottom)
}

// InsertLines inserts the given number of lines at the top of the scrollable
// region, pushing lines below down.
func InsertLines(n int) {
	NewOutputWithProfile(defaultOutputFile, ANSI).InsertLines(n)
}

// DeleteLines deletes the given number of lines, pulling any lines in
// the scrollable region below up.
func DeleteLines(n int) {
	NewOutputWithProfile(defaultOutputFile, ANSI).DeleteLines(n)
}

// EnableMousePress enables X10 mouse mode. Button press events are sent only.
func EnableMousePress() {
	NewOutputWithProfile(defaultOutputFile, ANSI).EnableMousePress()
}

// DisableMousePress disables X10 mouse mode.
func DisableMousePress() {
	NewOutputWithProfile(defaultOutputFile, ANSI).DisableMousePress()
}

// EnableMouse enables Mouse Tracking mode.
func EnableMouse() {
	NewOutputWithProfile(defaultOutputFile, ANSI).EnableMouse()
}

// DisableMouse disables Mouse Tracking mode.
func DisableMouse() {
	NewOutputWithProfile(defaultOutputFile, ANSI).DisableMouse()
}

// EnableMouseHilite enables Hilite Mouse Tracking mode.
func EnableMouseHilite() {
	NewOutputWithProfile(defaultOutputFile, ANSI).EnableMouseHilite()
}

// DisableMouseHilite disables Hilite Mouse Tracking mode.
func DisableMouseHilite() {
	NewOutputWithProfile(defaultOutputFile, ANSI).DisableMouseHilite()
}

// EnableMouseCellMotion enables Cell Motion Mouse Tracking mode.
func EnableMouseCellMotion() {
	NewOutputWithProfile(defaultOutputFile, ANSI).EnableMouseCellMotion()
}

// DisableMouseCellMotion disables Cell Motion Mouse Tracking mode.
func DisableMouseCellMotion() {
	NewOutputWithProfile(defaultOutputFile, ANSI).DisableMouseCellMotion()
}

// EnableMouseAllMotion enables All Motion Mouse mode.
func EnableMouseAllMotion() {
	NewOutputWithProfile(defaultOutputFile, ANSI).EnableMouseAllMotion()
}

// DisableMouseAllMotion disables All Motion Mouse mode.
func DisableMouseAllMotion() {
	NewOutputWithProfile(defaultOutputFile, ANSI).DisableMouseAllMotion()
}

// SetWindowTitle sets the terminal window title.
func SetWindowTitle(title string) {
	NewOutputWithProfile(defaultOutputFile, ANSI).SetWindowTitle(title)
}
