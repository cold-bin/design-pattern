// @author cold bin
// @date 2023/8/28

package foobarbaz

import "github.com/google/wire"

var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)

var MegaSet = wire.NewSet(SuperSet)
