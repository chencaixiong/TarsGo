APP       := StressTest
TARGET    := ContextTestServer
MFLAGS    :=
DFLAGS    :=
CONFIG    := client
STRIP_FLAG:= N
J2GO_FLAG:= 

libpath=${subst :, ,$(GOPATH)}
$(foreach path,$(libpath),$(eval -include $(path)/src/github.com/chencaixiong/TarsGo/tars/makefile.tars))
