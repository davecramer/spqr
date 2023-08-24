package app_test

import (
	"io"
	"testing"

	"github.com/pg-sharding/spqr/pkg/config"
	app "github.com/pg-sharding/spqr/router"
	mockcl "github.com/pg-sharding/spqr/router/mock/client"
	mockqr "github.com/pg-sharding/spqr/router/mock/qrouter"

	mockcmgr "github.com/pg-sharding/spqr/router/mock/rulerouter"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestFrontendSimpleEOF(t *testing.T) {

	assert := assert.New(t)
	ctrl := gomock.NewController(t)

	cl := mockcl.NewMockRouterClient(ctrl)
	qr := mockqr.NewMockQueryRouter(ctrl)
	cmngr := mockcmgr.NewMockPoolMgr(ctrl)

	cl.EXPECT().Usr().AnyTimes().Return("user1")
	cl.EXPECT().DB().AnyTimes().Return("db1")
	cl.EXPECT().Close().Times(1)

	cl.EXPECT().Receive().Times(1).Return(nil, io.EOF)

	cmngr.EXPECT().UnRouteCB(gomock.Any(), gomock.Any()).Times(1)

	err := app.Frontend(qr, cl, cmngr, &config.Router{})

	assert.NoError(err, "")
}

/*

TODO: make this work

separate client interface and client-related operations

func TestFrontendSimple(t *testing.T) {

	assert := assert.New(t)
	ctrl := gomock.NewController(t)

	cl := mockcl.NewMockRouterClient(ctrl)
	srv := mocksrv.NewMockServer(ctrl)
	qr := mockqr.NewMockQueryRouter(ctrl)
	cmngr := mockcmgr.NewMockPoolMgr(ctrl)

	frrule := &config.FrontendRule{
		DB:  "db1",
		Usr: "user1",
	}

	beRule := &config.BackendRule{}

	cl.EXPECT().Usr().AnyTimes().Return("user1")
	cl.EXPECT().DB().AnyTimes().Return("db1")

	cl.EXPECT().ID().AnyTimes().Return("lolkekcheburek")

	cl.EXPECT().Close().Times(1)
	cl.EXPECT().Rule().AnyTimes().Return(
		frrule,
	)

	cl.EXPECT().ReplyDebugNotice(gomock.Any()).AnyTimes().Return(nil)
	cl.EXPECT().AssignServerConn(gomock.Any()).AnyTimes().Return(nil)

	query := &pgproto3.Query{
		String: "SELECT 1",
	}

	cl.EXPECT().Receive().Times(1).Return(query, nil)
	cl.EXPECT().Receive().Times(1).Return(nil, io.EOF)

	// reroute on first query in this case
	cmngr.EXPECT().ValidateReRoute(gomock.Any()).Return(true)

	cmngr.EXPECT().RouteCB(cl, gomock.Any()).AnyTimes()

	cmngr.EXPECT().UnRouteCB(gomock.Any(), gomock.Any()).AnyTimes()

	cmngr.EXPECT().TXBeginCB(gomock.Any()).AnyTimes()

	cmngr.EXPECT().TXEndCB(gomock.Any()).AnyTimes()

	qr.EXPECT().Route(gomock.Any(), &lyx.Select{
		TargetList: []lyx.Node{
			&lyx.AExprConst{Value: "1"},
		},
		Where: &lyx.AExprEmpty{},
	}, nil).Return(qrouter.ShardMatchState{
		Routes: []*qrouter.DataShardRoute{
			{
				Shkey: kr.ShardKey{
					Name: "sh1",
				},
			},
		},
	}, nil).Times(1)

	route := route.NewRoute(beRule, frrule, map[string]*config.Shard{
		"sh1": &config.Shard{},
	})

	cl.EXPECT().Route().AnyTimes().Return(route)

	cl.EXPECT().ProcQuery(query, true, true).Do(

	).Return(nil)

	err := app.Frontend(qr, cl, cmngr, &config.Router{})

	assert.NoError(err, "")
}

*/
