package main

import (
	"fmt"
	"os"

	"github.com/bendahl/uinput"
	evdev "github.com/gvalkov/golang-evdev"
)

func main() {

	//创建虚拟输入设备
	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("testkeyboard"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//关闭设备
	defer keyboard.Close()

	//打开输入设备/dev/input/event3
	device, err := evdev.Open("/dev/input/event3")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//抓取/dev/input/event3
	device.Grab()

	//释放/dev/input/event3
	defer device.Release()

	for {
		//读取键盘输入
		events, err := device.Read()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//Esc终止程序 test
		for i := range events {
			e := events[i]
			if e.Code == evdev.KEY_ESC {
				fmt.Println("exit")
				os.Exit(0)
			}
		}
	}
}
