package puppy

import(
	"github.com/EllyABB/Example_sub_mod"
)

func Ladrar() string{
	return "woff"
}

func Ladridos() string{
	return "woff, woff, woff"
}

func La_que_llama_a_otra() string{
	return sub_pack.Algo(Ladrar)
}


func La_que_llama_a_otra2() string{
	return sub_pack.Algo(Ladridos)
}