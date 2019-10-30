package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Propper interface {
	Go()
}

func New(seed int64) Propper {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	return &propper{seed, rand.New(rand.NewSource(seed))}
}

type propper struct {
	seed int64
	r *rand.Rand
}

func (p *propper) Go() {
	//noinspection GoUnhandledErrorResult
	fmt.Fprintf(os.Stderr, "seed: %d\n", p.seed)

	var previousLine string
	for _, line := range p.story() {
		if line == previousLine {
			// when a task follows a struggle we assume
			// "A false hero presents unfounded claims"
			// doesn't happen twice in a row
			continue
		}
		fmt.Println(line)
		previousLine = line
	}
}

func (p *propper) story() []string {
	return p.applyAllRules(p.act1, p.act2, p.act3)
}

func (p *propper) act1() []string {
	return p.applyAllFunctions("A","B","C","↑","D","E","F","G",)
}

func (p *propper) act2() []string {
	return p.applyAllRules(p.struggleOrNil, p.taskOrNil)
}

func (p *propper) act3() []string {
	return p.applyAllFunctions("Q", "Ex", "T", "U", "W")
}

func (p *propper) struggle() []string {
	return p.applyAllFunctions("H", "J", "I", "K", "↓", "Pr", /*"−",*/ "Rs", "o", "L")
}

func (p *propper) struggleOrNil() []string {
	return p.applyAnyRule(p.struggle, p.null)
}

func (p *propper) task() []string {
	return p.applyAllFunctions("L", "M", "J", "N", "K", "↓", "Pr", /*"−",*/ "Rs")
}

func (p *propper) taskOrNil() []string {
	return p.applyAnyRule(p.task, p.null)
}

func (p *propper) null() []string {
	return []string{}
}

func (p *propper) applyAllRules(rules ...func() []string) (result []string) {
	for _, rule := range rules {
		result = append(result, rule()...)
	}
	return
}

func (p *propper) applyAnyRule(rules ...func() []string) []string {
	index := p.intn(len(rules))
	rule := rules[index]
	return rule()
}

func (p *propper) applyAllFunctions(fns ...string) (result []string) {
	for _, fn := range fns {
		result = append(result, p.applyAnyFunction(fn))
	}
	return
}

func (p *propper) applyAnyFunction(fn string) string {
	return p.choose(functions[fn])
}

func (p *propper) intn(n int) int {
	return p.r.Intn(n)
}

func (p *propper) choose(choices []string) string {
	index := p.intn(len(choices))
	return choices[index]
}
