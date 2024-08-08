package naspanclient

import (
	naspanclient "github.com/bitwormhole/naspan-client"
	"github.com/bitwormhole/naspan-client/gen/main4naspanclient"
	"github.com/bitwormhole/naspan-client/gen/test4naspanclient"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/security/modules/security"
)

// Module  ...
func Module() application.Module {
	mb := naspanclient.NewMainModule()
	mb.Components(main4naspanclient.ExportComponents)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(security.Module())

	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := naspanclient.NewTestModule()
	mb.Components(test4naspanclient.ExportComponents)
	return mb.Create()
}
