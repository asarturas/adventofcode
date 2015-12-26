package solution
import (
	"strings"
	"strconv"
)

type BrightnessInstructionParser struct {}

func (this BrightnessInstructionParser) Parse(writing string) Instruction {
	ins := Instruction{}
	ins.Act, writing = this.parseAction(writing)
	ins.Start, ins.End = this.parseCoords(writing)
	return ins;
}

func (this BrightnessInstructionParser) parseAction(writing string) (Action, string) {
	if strings.HasPrefix(writing, "turn on") {
		return TurnOnBrightness{}, strings.Replace(writing, "turn on ", "", 1);
	} else if strings.HasPrefix(writing, "turn off") {
		return TurnOffBrightness{}, strings.Replace(writing, "turn off ", "", 1);
	} else {
		return ToggleBrightness{}, strings.Replace(writing, "toggle ", "", 1);
	}
}

func (this BrightnessInstructionParser) parseCoords(writing string) (Coord, Coord) {
	parts := strings.Split(writing, " ");
	start := strings.Split(parts[0], ",");
	startRow, _ := strconv.Atoi(start[0]);
	startCol, _ := strconv.Atoi(start[1]);
	end := strings.Split(parts[2], ",");
	endRow, _ := strconv.Atoi(end[0]);
	endCol, _ := strconv.Atoi(end[1]);
	return Coord{startRow, startCol}, Coord{endRow, endCol}
}
