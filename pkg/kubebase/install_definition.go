package kubebase

type PIPE_DEF map[interface{}][]string

type INSTALL_FUNCS map[string]func(input map[string]string)

type INSTALL_VALUES map[string]string

var IFN_CTX = make(INSTALL_FUNCS)

var IVL_CTX = make(INSTALL_VALUES)

func _InitPipeDef() PIPE_DEF {

	var sess_pdef = make(PIPE_DEF)

	return sess_pdef
}

var SESS_PIPEDEF = _InitPipeDef()
