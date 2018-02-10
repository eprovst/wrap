package ast

// ElementType is another representation of an elements type
type ElementType byte

// Defines the different ElementTypes
const (
	TScene ElementType = iota
	TBeginAct
	TEndAct
	TAction
	TDialogue
	TDualDialogue
	TLyrics
	TTransition
	TCenteredText
	TPageBreak
	TSection
	TSynopse
	TNote
	TNone
)

/*GetElementType returns the type of element as a
lowercase string and "TNone" if unknown. */
func GetElementType(elem Element) ElementType {
	switch elem.(type) {
	case Scene:
		return TScene
	case BeginAct:
		return TBeginAct
	case EndAct:
		return TEndAct
	case Action:
		return TAction
	case Dialogue:
		return TDialogue
	case DualDialogue:
		return TDualDialogue
	case Lyrics:
		return TLyrics
	case Transition:
		return TTransition
	case CenteredText:
		return TCenteredText
	case PageBreak:
		return TPageBreak
	case Section:
		return TSection
	case Synopse:
		return TSynopse
	case Note:
		return TNote
	default:
		return TNone
	}
}
