package placement

import (
	"github.com/imxyy1soope1/go-hyprland-ipc/hyprctl"
)

func Init(client *hyprctl.Client) (c Canvas, err error) {
	windows, err := client.Windows()
	if err != nil {
		return
	}

}
