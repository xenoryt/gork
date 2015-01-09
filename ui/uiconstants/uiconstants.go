/*Package uiconstants is for storing constants for UI in a central location
for easy modification*/
package uiconstants

//Minimum dimensions for the window
const (
	MinWidth  = 80
	MinHeight = 24
)

//For sending commands
type CommandType int

const (
	CMD_UPDATE CommandType = iota
	CMD_TRACK
	CMD_REMOVE
	CMD_EXIT
)
