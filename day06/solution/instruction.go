package solution
import (
	"strings"
	"strconv"
)

type Coord struct {
	Row, Col int
}

type Instruction struct {
	Act Action
	Start, End Coord
}

type InstructionParser struct {
	actionOptions []Action
}

func NewInstructionParser(actions ...Action) InstructionParser {
	return InstructionParser{actions}
}

func (this InstructionParser) Parse(writing string) Instruction {
	ins := Instruction{}
	ins.Act, writing = this.parseAction(writing)
	ins.Start, ins.End = this.parseCoords(writing)
	return ins;
}

func (this InstructionParser) parseAction(writing string) (Action, string) {
	for _, action := range this.actionOptions {
		if strings.HasPrefix(writing, action.String()) {
			return action, strings.Replace(writing, action.String(), "", 1);
		}
	}
	return nil, writing
}

func (this InstructionParser) parseCoords(writing string) (Coord, Coord) {
	parts := strings.Split(writing, " ");
	start := strings.Split(parts[0], ",");
	startRow, _ := strconv.Atoi(start[0]);
	startCol, _ := strconv.Atoi(start[1]);
	end := strings.Split(parts[2], ",");
	endRow, _ := strconv.Atoi(end[0]);
	endCol, _ := strconv.Atoi(end[1]);
	return Coord{startRow, startCol}, Coord{endRow, endCol}
}
