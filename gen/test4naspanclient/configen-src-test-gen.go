package test4naspanclient
import (
    pfa2f69e31 "github.com/bitwormhole/naspan-client/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type pfa2f69e31.DemoUnit in package:github.com/bitwormhole/naspan-client/src/test/golang/unit
//
// id:com-fa2f69e31f34595e-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pfa2f69e31f_unit_DemoUnit struct {
}

func (inst* pfa2f69e31f_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fa2f69e31f34595e-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfa2f69e31f_unit_DemoUnit) new() any {
    return &pfa2f69e31.DemoUnit{}
}

func (inst* pfa2f69e31f_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfa2f69e31.DemoUnit)
	nop(ie, com)

	


    return nil
}


