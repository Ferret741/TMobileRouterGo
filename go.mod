module localhost/tmobile

go 1.24.1

replace (
	localhost/tmobile => ./
	localhost/tmobile/router => ./router
)

require localhost/tmobile/router v0.0.0-00010101000000-000000000000
