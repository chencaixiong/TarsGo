all:
	tars2go -add-servant=false -tarsPath github.com/chencaixiong/TarsGo/tars *.tars
	#override file
	sed -i 's|"endpointf"|"github.com/chencaixiong/TarsGo/tars/protocol/res/endpointf"|g' queryf/QueryF_IF.go
