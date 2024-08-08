package main4naspanclient
import (
    p26a4d739d "github.com/bitwormhole/naspan-client/app/data/dao"
    p9eefbbba1 "github.com/bitwormhole/naspan-client/app/data/database"
    p2d015274b "github.com/bitwormhole/naspan-client/app/data/dxo"
    p0e8b51b48 "github.com/bitwormhole/naspan-client/app/implements/example"
    p9bbb03b92 "github.com/bitwormhole/naspan-client/app/web/controllers/apiv1"
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
     "github.com/starter-go/application"
)

// type p9eefbbba1.ThisGroup in package:github.com/bitwormhole/naspan-client/app/data/database
//
// id:com-9eefbbba1c40e338-database-ThisGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:alias-2d015274bc135dbf2a2761d31deb16b0-DatabaseAgent
// scope:singleton
//
type p9eefbbba1c_database_ThisGroup struct {
}

func (inst* p9eefbbba1c_database_ThisGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9eefbbba1c40e338-database-ThisGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = "alias-2d015274bc135dbf2a2761d31deb16b0-DatabaseAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9eefbbba1c_database_ThisGroup) new() any {
    return &p9eefbbba1.ThisGroup{}
}

func (inst* p9eefbbba1c_database_ThisGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9eefbbba1.ThisGroup)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Alias = inst.getAlias(ie)
    com.URI = inst.getURI(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)
    com.SourceManager = inst.getSourceManager(ie)


    return nil
}


func (inst*p9eefbbba1c_database_ThisGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.default.enabled}")
}


func (inst*p9eefbbba1c_database_ThisGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.alias}")
}


func (inst*p9eefbbba1c_database_ThisGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.uri}")
}


func (inst*p9eefbbba1c_database_ThisGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.table-name-prefix}")
}


func (inst*p9eefbbba1c_database_ThisGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.datasource}")
}


func (inst*p9eefbbba1c_database_ThisGroup) getSourceManager(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}



// type p0e8b51b48.DaoImpl in package:github.com/bitwormhole/naspan-client/app/implements/example
//
// id:com-0e8b51b48d7e697c-example-DaoImpl
// class:
// alias:alias-26a4d739d8aa0a726b196e8f40ff0c99-ExampleDAO
// scope:singleton
//
type p0e8b51b48d_example_DaoImpl struct {
}

func (inst* p0e8b51b48d_example_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0e8b51b48d7e697c-example-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-26a4d739d8aa0a726b196e8f40ff0c99-ExampleDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0e8b51b48d_example_DaoImpl) new() any {
    return &p0e8b51b48.DaoImpl{}
}

func (inst* p0e8b51b48d_example_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0e8b51b48.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)


    return nil
}


func (inst*p0e8b51b48d_example_DaoImpl) getAgent(ie application.InjectionExt)p2d015274b.DatabaseAgent{
    return ie.GetComponent("#alias-2d015274bc135dbf2a2761d31deb16b0-DatabaseAgent").(p2d015274b.DatabaseAgent)
}



// type p9bbb03b92.ExampleController in package:github.com/bitwormhole/naspan-client/app/web/controllers/apiv1
//
// id:com-9bbb03b9213f00d7-apiv1-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p9bbb03b921_apiv1_ExampleController struct {
}

func (inst* p9bbb03b921_apiv1_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9bbb03b9213f00d7-apiv1-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9bbb03b921_apiv1_ExampleController) new() any {
    return &p9bbb03b92.ExampleController{}
}

func (inst* p9bbb03b921_apiv1_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9bbb03b92.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p9bbb03b921_apiv1_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p9bbb03b921_apiv1_ExampleController) getDao(ie application.InjectionExt)p26a4d739d.ExampleDAO{
    return ie.GetComponent("#alias-26a4d739d8aa0a726b196e8f40ff0c99-ExampleDAO").(p26a4d739d.ExampleDAO)
}


