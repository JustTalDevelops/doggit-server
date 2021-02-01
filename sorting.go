package main

import (
	"github.com/JustTalDevelops/doggit"
	"sort"
)

func GetOrderedHandlers(libraries []*doggit.Library) []doggit.Handler {
	var sortedHandlers PriorityHandlerList

	for _, l := range libraries {
		for _, h := range l.AllHandlers() {
			sortedHandlers = append(sortedHandlers, h)
		}
	}

	sort.Sort(sort.Reverse(sortedHandlers))

	return sortedHandlers
}

type PriorityHandlerList []doggit.Handler

func (p PriorityHandlerList) Len() int           { return len(p) }
func (p PriorityHandlerList) Less(i, j int) bool { return p[i].Priority() < p[j].Priority() }
func (p PriorityHandlerList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
