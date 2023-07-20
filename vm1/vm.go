package vm1

import (
	"fmt"     // for tracing only
	"reflect" // for optional CallX only
	"strconv" // for tracing only
)

// Byte-code instruction set.
const (
	// instruction effect on stack: values consumed -- values produced
	Nop      = iota // --
	Add             // n1 n2 -- sum ; sum = n1+n2
	Assign          // addr val -- ; mem[addr] = val
	Call            // f [a1 .. ai] -- [r1 .. rj] ; r1, ... = prog[f](a1, ...)
	CallX           // f [a1 .. ai] -- [r1 .. rj] ; r1, ... = mem[f](a1, ...)
	Dup             // addr -- value ; value = mem[addr]
	Fdup            // addr -- value ; value = mem[addr]
	Enter           // -- ; enter frame: push(fp), fp = sp
	Exit            // -- ;
	Jump            // -- ; ip += $1
	JumpTrue        // cond -- ; if cond { ip += $1 }
	Lower           // n1 n2 -- cond ; cond = n1 < n2
	Pop             // v --
	Push            // -- v
	Return          // [r1 .. ri] -- ; exit frame: sp = fp, fp = pop
	Sub             // n1 n2 -- diff ; diff = n1 - n2
)

var strop = [...]string{ // for VM tracing.
	Nop:      "Nop",
	Add:      "Add",
	Assign:   "Assign",
	Call:     "Call",
	CallX:    "CallX",
	Dup:      "Dup",
	Fdup:     "Fdup",
	Enter:    "Enter",
	Exit:     "Exit",
	Jump:     "Jump",
	JumpTrue: "JumpTrue",
	Lower:    "Lower",
	Pop:      "Pop",
	Push:     "Push",
	Return:   "Return",
	Sub:      "Sub",
}

// Machine represents a virtual machine.
type Machine struct {
	code   [][]int64 // code to execute
	mem    []any     // memory, as a stack
	ip, fp int       // instruction and frame pointer
	// flags  uint      // to set debug mode, restrict CallX, etc...
}

// Run runs a program.
func (m *Machine) Run() {
	code, mem, ip, fp, sp := m.code, m.mem, m.ip, m.fp, 0

	defer func() { m.mem, m.ip, m.fp = mem, ip, fp }()

	trace := func() {
		var op1 string
		if len(code[ip]) > 1 {
			op1 = strconv.Itoa(int(code[ip][1]))
		}
		fmt.Printf("ip:%-4d sp:%-4d fp:%-4d op:[%-8s %-4s] mem:%v\n", ip, sp, fp, strop[code[ip][0]], op1, mem)
	}
	_ = trace

	for {
		sp = len(mem) // stack pointer
		trace()
		switch op := code[ip]; op[0] { // TODO: op[0] will contain file pos ?
		case Add:
			mem[sp-2] = mem[sp-2].(int) + mem[sp-1].(int)
			mem = mem[:sp-1]
		case Assign:
			mem[mem[sp-2].(int)] = mem[sp-1]
			mem = mem[:sp-1]
		case Call:
			mem = append(mem, ip+1)
			ip += int(op[1])
			continue
		case CallX: // Should be made optional.
			in := make([]reflect.Value, int(op[1]))
			for i := range in {
				in[i] = reflect.ValueOf(mem[sp-1-i])
			}
			f := reflect.ValueOf(mem[sp-len(in)-1])
			mem = mem[:sp-len(in)-1]
			for _, v := range f.Call(in) {
				mem = append(mem, v.Interface())
			}
		case Dup:
			mem = append(mem, mem[int(op[1])])
		case Enter:
			mem = append(mem, fp)
			fp = sp + 1
		case Exit:
			return
		case Fdup:
			mem = append(mem, mem[int(op[1])+fp-1])
		case Jump:
			ip += int(op[1])
			continue
		case JumpTrue:
			cond := mem[sp-1].(bool)
			mem = mem[:sp-1]
			if cond {
				ip += int(op[1])
				continue
			}
		case Lower:
			mem[sp-2] = mem[sp-2].(int) < mem[sp-1].(int)
			mem = mem[:sp-1]
		case Pop:
			mem = mem[:sp-1]
		case Push:
			mem = append(mem, int(op[1]))
		case Return:
			ip = mem[fp-2].(int)
			ofp := fp
			fp = mem[fp-1].(int)
			mem = append(mem[:ofp-int(op[1])-2], mem[sp-int(op[1]):]...)
			continue
		case Sub:
			mem[sp-2] = mem[sp-2].(int) - mem[sp-1].(int)
			mem = mem[:sp-1]
		}
		ip++
	}
}

func (m *Machine) PushCode(code [][]int64) (p int) {
	p = len(m.code)
	m.code = append(m.code, code...)
	return
}
func (m *Machine) SetIP(ip int)       { m.ip = ip }
func (m *Machine) Push(v any) (l int) { l = len(m.mem); m.mem = append(m.mem, v); return }
func (m *Machine) Pop() (v any)       { l := len(m.mem) - 1; v = m.mem[l]; m.mem = m.mem[:l]; return }
