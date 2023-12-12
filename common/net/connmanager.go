package net

type ConnManager struct {
	connections ShardLockMaps
}

func newConnManager() *ConnManager {
	return &ConnManager{
		connections: NewShardLockMaps(),
	}
}
func (connMgr *ConnManager) GenConnID() string {

}

func (connMgr *ConnManager) Add(conn Connection) {

	connMgr.connections.Set(conn.GetConnIdStr(), conn)
}

func (connMgr *ConnManager) Remove(conn Connection) {

	connMgr.connections.Remove(conn.GetConnIdStr())
}
