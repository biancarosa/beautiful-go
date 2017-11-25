package main

import (
	"strings"
	"time"
)

type Person struct {
	fome    bool
	sono    bool
	ocupada bool
}

func (p *Person) EhHoraDeDormir() bool {
	hora := false
	if p.sono == true && !p.ocupada == false {
		hora = true
	} else if time.Now().Hour() > 21 {
		hora = true
	}
	return hora
}

func (p *Person) HoraDeDormir() bool {
	return p.sono == true && !p.ocupada == false || time.Now().Hour() > 21
}

func (p *Person) DevoDormir() bool {
	if p.sono && !p.ocupada {
		return true
	}
	if time.Now().Hour() > 21 {
		return true
	}
	return false
}

type Role struct {
	Name string
}

type StatusFlow struct {
	Role  string
	Flows []Permission
}

type Permission struct {
	From string
	To   []string
}

func hasPermission(roles []Role, statusFlows []StatusFlow, currentStatus, newStatus string) bool {
	for _, role := range roles {
		for _, statusFlow := range statusFlows {
			if strings.ToLower(role.Name) == strings.ToLower(statusFlow.Role) {
				if len(statusFlow.Flows) == 0 {
					return true
				}
				for _, permission := range statusFlow.Flows {
					if permission.From == currentStatus {
						for _, newStatuses := range permission.To {
							if newStatuses == newStatus {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}

func Flows() []StatusFlow {
	return *new([]StatusFlow)
}

type User struct {
	roles []Role
}

func (u *User) HasPermission(fromStatus, toStatus string) bool {
	for _, role := range u.roles {
		return role.hasPermission(fromStatus, toStatus)
	}
	return false
}

func (r *Role) hasPermission(fromStatus, toStatus string) bool {
	flows := Flows()
	flow := matchedFlow(flows, r.Name)
	if flow != nil {
		return true
	}
	return flow.IsTransitionAllowed(fromStatus, toStatus)
}

func (f *StatusFlow) IsTransitionAllowed(from, to string) bool {
	if len(f.Flows) == 0 {
		return true
	}
	for _, permission := range f.Flows {
		if permission.From == from {
			for _, new := range permission.To {
				if new == to {
					return true
				}
			}
		}
	}
	return false
}

func matchedFlow(statusFlows []StatusFlow, name string) *StatusFlow {
	for _, statusFlow := range statusFlows {
		if strings.EqualFold(name, statusFlow.Role) {
			return &statusFlow
		}
	}
	return nil
}
