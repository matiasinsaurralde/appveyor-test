package mybinding

/*
#include <stdio.h>
typedef int (*MyFunc)(int);
MyFunc f;

void** function() {
	return (void**)&f;
}

int call(int i) {
	return f(i);
}

typedef char* (*GetStringFunc)();
GetStringFunc getString;

void** getStringFunc() {
	return (void**)&getString;
}

char* callGetString() {
	char* output = getString();
	printf("lol %s\n", output);
	return output;
}
*/
import "C"

import (
	"github.com/matiasinsaurralde/go-dotnet/dotnet"
)

func init() {
	dotnet.SetupDelegates(func() error {
		// f := C.thefunction()
		f := C.function()
		// fmt.Println("f = ", f)
		err := dotnet.CreateDelegate(
			"HelloWorld",
			"HelloWorld.HelloWorld",
			"Hello",
			0, f)

		f2 := C.getStringFunc()
		err = dotnet.CreateDelegate(
			"HelloWorld",
			"HelloWorld.HelloWorld",
			"GetString",
			0, f2)
		return err
	})
	err := dotnet.Init()
	if err != nil {
		panic(err)
	}
}

func Hello(a int) int {
	r := C.call(C.int(a))
	return int(r)
}

func GetString() string {
	r := C.callGetString()
	str := C.GoString(r)
	return str
}
