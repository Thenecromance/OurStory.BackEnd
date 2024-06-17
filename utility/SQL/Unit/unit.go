package Unit

import (
	"fmt"
	"github.com/Thenecromance/OurStories/utility/log"
	"regexp"
	"strings"
)

const (
	Infinity = -1
)

type placeHolder struct {
	mask  string
	count int
}

// the smallest Unit of the sql command
type Unit struct {
	uid       int64
	index     int
	command   string
	argsCount int
	args      []any // args container
	next      *Unit // other Unit
	placeHolder
}

func (u *Unit) appendArgs(args []any) {
	if len(args) == 0 {
		return
	}
	if u.argsCount == Infinity || u.argsCount == len(u.args) {
		u.args = append(u.args, args...)
		return
	}
	if len(args) > u.argsCount {
		log.Error("args count not match")
		return
	}
	// < u.argsCount
	u.args = append(u.args, args...)
}

func (u *Unit) SetArgs(idx int, args []any) {
	// check the index is valid
	if idx < u.index {
		log.Errorf("index out of range current index is %d, but the index is %d", u.index, idx)
		return
	}
	// if the index is equal to the current index, set the args
	if idx == u.index {
		u.appendArgs(args)
		return
	}
	if u.next != nil {
		u.next.SetArgs(idx, args)
	} else {
		log.Errorf("index out of range, the max index is %d, but the index is %d", u.index, idx)
	}
}

func (u *Unit) SetArg(idx int, arg any) {
	u.SetArgs(idx, []any{arg})
}

func (u *Unit) SetUID(uid int64) {
	u.uid = uid
}

func (u *Unit) UID() int64 {
	return u.uid
}
func (u *Unit) GetUIDs() (res []int64) {
	res = append(res, u.uid)
	tmp := u.next
	for tmp != nil {
		res = append(res, tmp.uid)
		tmp = tmp.next
	}
	return
}

func (u *Unit) Count() (cnt int) {
	cnt = 1
	tmp := u.next
	for tmp != nil {
		cnt++
		tmp = tmp.next
	}
	return
}
func (u *Unit) createPlacholder() {
	//u.command = strings.ReplaceAll(u.command, "%s", "?")
	start := strings.Index(u.command, "(")
	end := strings.Index(u.command, ")")
	cnt := strings.Count(u.command[start:end], ",")
	if cnt == 0 {
		u.mask = "?"
		return
	}
	u.count = cnt + 1
	u.mask = strings.Repeat("?,", u.count)
	u.mask = u.mask[:len(u.mask)-1]
	u.mask = "(" + u.mask + ")"
}
func (u *Unit) Command() string {
	if u.argsCount == Infinity {
		u.createPlacholder()
		argCount := len(u.args)
		if argCount%u.count != 0 {
			log.Errorf("args count not match, the args count is %d, but the mask count is %d", argCount, u.count)
			panic("args count not match")
			return u.command
		}
		final := strings.Repeat(u.mask+",", argCount/u.count)
		final = final[:len(final)-1]
		return fmt.Sprintf(u.command, final)
	}
	return u.command
}

func (u *Unit) Args() []any {
	if u.argsCount == Infinity && len(u.args) == 0 {
		log.Error("args count is dynamic, but args is empty")
		return nil
	}
	if u.argsCount > 0 && len(u.args) != u.argsCount {
		log.Errorf("args count not match %d %d", u.argsCount, len(u.args))
		return nil
	}
	return u.args
}

func (u *Unit) Next() *Unit {
	return u.next
}

func splitCommands(command string) []string {

	re := regexp.MustCompile(`(?m)--.*$`)
	command = re.ReplaceAllString(command, "")
	command = strings.ReplaceAll(command, "\n", " ")

	result := strings.SplitAfter(command, ";")
	if len(result) == 0 {
		return nil
	}
	if len(result) == 1 {
		return result
	} else {
		return result[:len(result)-1] // remove the last empty string
	}

}
func parseArgsCount(cmd string) int {
	if strings.Contains(cmd, "%s") {
		return Infinity
	}
	if strings.Count(cmd, "?") > 0 {
		return strings.Count(cmd, "?")
	}
	if strings.Count(cmd, "$") > 0 {
		return strings.Count(cmd, "$")
	}
	return 0
}

func newNode(idx int, commandGroup []string) *Unit {
	if len(commandGroup) == 0 {
		return nil
	}
	u := &Unit{
		index:   idx,
		command: commandGroup[0],
		next:    newNode(idx+1, commandGroup[1:]),
	}

	u.argsCount = parseArgsCount(u.command)
	if u.argsCount > 0 {
		u.args = make([]any, u.argsCount)
	}
	return u
}

func New(cmd string) *Unit {
	commandGroup := splitCommands(cmd)
	u := newNode(0, commandGroup)
	u.argsCount = parseArgsCount(u.command)
	if u.argsCount > 0 {
		u.args = make([]any, 0, u.argsCount)
	}
	return u
}
