package Unit

import (
	"reflect"
	"testing"
)

const (
	singleCommand         = "select * from table;"
	multiCommand          = "select * from table;select * from table2;"
	singleCommandWithArgs = "select * from table where id = ?;"
	multiCommandWithArgs  = "select * from table where id = ?;select * from table where id = ?;"

	commandWithDynamicArgs = "Insert into `test`(`a`,`b`,`c`)VALUES %s;"
)

func TestNew(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name string
		args args
		want *Unit
	}{
		// basic test
		{name: "newSingleCommand", args: args{cmd: singleCommand}, want: &Unit{
			index:     0,
			command:   singleCommand,
			argsCount: 0,
			args:      nil,
			next:      nil,
		}},
		{name: "newMultiCommand", args: args{cmd: multiCommand}, want: &Unit{
			index:     0,
			command:   "select * from table;",
			argsCount: 0,
			args:      nil,
			next: &Unit{
				index:   1,
				command: "select * from table2;",
				next:    nil,
			},
		}},
		{name: "newSingleCommandWithArgs", args: args{cmd: singleCommandWithArgs}, want: &Unit{
			index:     0,
			command:   "select * from table where id = ?;",
			argsCount: 1,
			args:      make([]any, 1),
			next:      nil,
		}},
		{name: "newMultiCommandWithArgs", args: args{cmd: multiCommandWithArgs}, want: &Unit{
			index:     0,
			command:   "select * from table where id = ?;",
			argsCount: 1,
			args:      make([]any, 1),
			next: &Unit{
				index:     1,
				command:   "select * from table where id = ?;",
				argsCount: 1,
				args:      make([]any, 1),
				next:      nil,
			},
		}},
		{name: "newDynamicArgs", args: args{cmd: commandWithDynamicArgs}, want: &Unit{
			index:     0,
			command:   "Insert into `test`(`a`,`b`,`c`)VALUES %s;",
			argsCount: Infinity,
			args:      nil,
			next:      nil,
		}},
		// empty test
		{name: "newEmptyCommand", args: args{cmd: ""}, want: nil},

		// idiot test only for avoid monkey using this shit
		{name: "newIdiotCommand", args: args{cmd: "select * fro"}, want: nil},
		{name: "onlySplash", args: args{cmd: ";"}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.cmd)
			for got != nil && tt.want != nil {
				if got.index != tt.want.index ||
					got.command != tt.want.command ||
					got.argsCount != tt.want.argsCount ||
					!reflect.DeepEqual(got.args, tt.want.args) {
					t.Errorf("New() = %v, want %v", got, tt.want)
				}
				got = got.next
				tt.want = tt.want.next
			}
		})
	}
}

/*
func Test_newNode(t *testing.T) {
	type args struct {
		idx          int
		commandGroup []string
	}
	tests := []struct {
		name string
		args args
		want *unit
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newNode(tt.args.idx, tt.args.commandGroup); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNode() = %v, want %v", got, tt.want)
			}
		})
	}
}*/

func Test_parseArgsCount(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "singleCommand", args: args{cmd: singleCommand}, want: 0},
		{name: "multiCommand", args: args{cmd: multiCommand}, want: 0},
		{name: "singleCommandWithArgs", args: args{cmd: singleCommandWithArgs}, want: 1},
		{name: "commandWithDynamicArgs", args: args{cmd: commandWithDynamicArgs}, want: Infinity},
		{name: "emptyCommand", args: args{cmd: ""}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseArgsCount(tt.args.cmd); got != tt.want {
				t.Errorf("parseArgsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitCommands(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "singleCommand", args: args{command: singleCommand}, want: []string{"select * from table;"}},
		{name: "multiCommand", args: args{command: multiCommand}, want: []string{"select * from table;", "select * from table2;"}},
		{name: "singleCommandWithArgs", args: args{command: singleCommandWithArgs}, want: []string{"select * from table where id = ?;"}},
		{name: "multiCommandWithArgs", args: args{command: multiCommandWithArgs}, want: []string{"select * from table where id = ?;", "select * from table where id = ?;"}},
		{name: "commandWithDynamicArgs", args: args{command: commandWithDynamicArgs}, want: []string{"Insert into `test`(`a`,`b`,`c`)VALUES %s;"}},
		{name: "emptyCommand", args: args{command: ""}, want: []string{""}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitCommands(tt.args.command); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unit_Args(t *testing.T) {
	type fields struct {
		index     int
		command   string
		argsCount int
		args      []any
		next      *Unit
	}
	tests := []struct {
		name   string
		fields fields
		want   []any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Unit{
				index:     tt.fields.index,
				command:   tt.fields.command,
				argsCount: tt.fields.argsCount,
				args:      tt.fields.args,
				next:      tt.fields.next,
			}
			if got := u.Args(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Args() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unit_Command(t *testing.T) {
	type fields struct {
		index     int
		command   string
		argsCount int
		args      []any
		next      *Unit
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Unit{
				index:     tt.fields.index,
				command:   tt.fields.command,
				argsCount: tt.fields.argsCount,
				args:      tt.fields.args,
				next:      tt.fields.next,
			}
			if got := u.Command(); got != tt.want {
				t.Errorf("Command() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unit_Next(t *testing.T) {
	type fields struct {
		index     int
		command   string
		argsCount int
		args      []any
		next      *Unit
	}
	tests := []struct {
		name   string
		fields fields
		want   *Unit
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Unit{
				index:     tt.fields.index,
				command:   tt.fields.command,
				argsCount: tt.fields.argsCount,
				args:      tt.fields.args,
				next:      tt.fields.next,
			}
			if got := u.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unit_SetArg(t *testing.T) {
	type fields struct {
		index     int
		command   string
		argsCount int
		args      []any
		next      *Unit
	}
	type args struct {
		idx int
		arg any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Unit{
				index:     tt.fields.index,
				command:   tt.fields.command,
				argsCount: tt.fields.argsCount,
				args:      tt.fields.args,
				next:      tt.fields.next,
			}
			u.SetArg(tt.args.idx, tt.args.arg)
		})
	}
}

func Test_unit_SetArgs(t *testing.T) {
	type fields struct {
		index     int
		command   string
		argsCount int
		args      []any
		next      *Unit
	}
	type args struct {
		idx  int
		args []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Unit{
				index:     tt.fields.index,
				command:   tt.fields.command,
				argsCount: tt.fields.argsCount,
				args:      tt.fields.args,
				next:      tt.fields.next,
			}
			u.SetArgs(tt.args.idx, tt.args.args)
		})
	}
}

func Test_unit_appendArgs(t *testing.T) {
	type fields struct {
		index     int
		command   string
		argsCount int
		args      []any
		next      *Unit
	}
	type args struct {
		args []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Unit{
				index:     tt.fields.index,
				command:   tt.fields.command,
				argsCount: tt.fields.argsCount,
				args:      tt.fields.args,
				next:      tt.fields.next,
			}
			u.appendArgs(tt.args.args)
		})
	}
}
