package Role

const (
	Guest  = iota // when user is not logged in or is not registered yet
	User          // when user is logged in
	Admin         // in this role, user can do some crazy stuff
	Master        // this is the GOD!!!!!!! he can do anything!!!!!!!
)
