package utils

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	PolicyDefault = "default"
	PolicyRandom  = "random"
	ProgramName   = "AntShell"
	Version       = "1.0"
	Author        = "Casstiel"
	GitHub        = "https://github.com/ooppwwqq0/AntShell-Go"
	Banner        = `
 _____     _   _____ _       _ _
|  _  |___| |_|   __| |_ ___| | |
|     |   |  _|__   |   | -_| | |
|__|__|_|_|_| |_____|_|_|___|_|_|
`
)

var (
	BannerList = []string{
		`
 _____     _   _____ _       _ _
|  _  |___| |_|   __| |_ ___| | |
|     |   |  _|__   |   | -_| | |
|__|__|_|_|_| |_____|_|_|___|_|_|
`,
		`
 _______         __   _______ __           __ __
|   _   |.-----.|  |_|     __|  |--.-----.|  |  |
|       ||     ||   _|__     |     |  -__||  |  |
|___|___||__|__||____|_______|__|__|_____||__|__|
`,
		`
 _______  __    _  _______  _______  __   __  _______  ___      ___
|   _   ||  |  | ||       ||       ||  | |  ||       ||   |    |   |
|  |_|  ||   |_| ||_     _||  _____||  |_|  ||    ___||   |    |   |
|       ||       |  |   |  | |_____ |       ||   |___ |   |    |   |
|       ||  _    |  |   |  |_____  ||       ||    ___||   |___ |   |___
|   _   || | |   |  |   |   _____| ||   _   ||   |___ |       ||       |
|__| |__||_|  |__|  |___|  |_______||__| |__||_______||_______||_______|
`,
		`
██      ▄     ▄▄▄▄▀ ▄▄▄▄▄    ▄  █ ▄███▄   █    █
█ █      █ ▀▀▀ █   █     ▀▄ █   █ █▀   ▀  █    █
█▄▄█ ██   █    █ ▄  ▀▀▀▀▄   ██▀▀█ ██▄▄    █    █
█  █ █ █  █   █   ▀▄▄▄▄▀    █   █ █▄   ▄▀ ███▄ ███▄
   █ █  █ █  ▀                 █  ▀███▀       ▀    ▀
  █  █   ██                   ▀
 ▀
`,
		`
   ▄████████ ███▄▄▄▄       ███        ▄████████    ▄█    █▄       ▄████████  ▄█        ▄█
  ███    ███ ███▀▀▀██▄ ▀█████████▄   ███    ███   ███    ███     ███    ███ ███       ███
  ███    ███ ███   ███    ▀███▀▀██   ███    █▀    ███    ███     ███    █▀  ███       ███
  ███    ███ ███   ███     ███   ▀   ███         ▄███▄▄▄▄███▄▄  ▄███▄▄▄     ███       ███
▀███████████ ███   ███     ███     ▀███████████ ▀▀███▀▀▀▀███▀  ▀▀███▀▀▀     ███       ███
  ███    ███ ███   ███     ███              ███   ███    ███     ███    █▄  ███       ███
  ███    ███ ███   ███     ███        ▄█    ███   ███    ███     ███    ███ ███▌    ▄ ███▌    ▄
  ███    █▀   ▀█   █▀     ▄████▀    ▄████████▀    ███    █▀      ██████████ █████▄▄██ █████▄▄██
                                                                            ▀         ▀
`,
		`
 █████╗ ███╗   ██╗████████╗███████╗██╗  ██╗███████╗██╗     ██╗
██╔══██╗████╗  ██║╚══██╔══╝██╔════╝██║  ██║██╔════╝██║     ██║
███████║██╔██╗ ██║   ██║   ███████╗███████║█████╗  ██║     ██║
██╔══██║██║╚██╗██║   ██║   ╚════██║██╔══██║██╔══╝  ██║     ██║
██║  ██║██║ ╚████║   ██║   ███████║██║  ██║███████╗███████╗███████╗
╚═╝  ╚═╝╚═╝  ╚═══╝   ╚═╝   ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝
`,
		`
    ___       ___       ___       ___       ___       ___       ___       ___
   /\  \     /\__\     /\  \     /\  \     /\__\     /\  \     /\__\     /\__\
  /::\  \   /:| _|_    \:\  \   /::\  \   /:/__/_   /::\  \   /:/  /    /:/  /
 /::\:\__\ /::|/\__\   /::\__\ /\:\:\__\ /::\/\__\ /::\:\__\ /:/__/    /:/__/
 \/\::/  / \/|::/  /  /:/\/__/ \:\:\/__/ \/\::/  / \:\:\/  / \:\  \    \:\  \
   /:/  /    |:/  /   \/__/     \::/  /    /:/  /   \:\/  /   \:\__\    \:\__\
   \/__/     \/__/               \/__/     \/__/     \/__/     \/__/     \/__/
`,
		`
      ___           ___           ___           ___           ___           ___           ___       ___
     /\  \         /\__\         /\  \         /\  \         /\__\         /\  \         /\__\     /\__\
    /::\  \       /::|  |        \:\  \       /::\  \       /:/  /        /::\  \       /:/  /    /:/  /
   /:/\:\  \     /:|:|  |         \:\  \     /:/\ \  \     /:/__/        /:/\:\  \     /:/  /    /:/  /
  /::\~\:\  \   /:/|:|  |__       /::\  \   _\:\~\ \  \   /::\  \ ___   /::\~\:\  \   /:/  /    /:/  /
 /:/\:\ \:\__\ /:/ |:| /\__\     /:/\:\__\ /\ \:\ \ \__\ /:/\:\  /\__\ /:/\:\ \:\__\ /:/__/    /:/__/
 \/__\:\/:/  / \/__|:|/:/  /    /:/  \/__/ \:\ \:\ \/__/ \/__\:\/:/  / \:\~\:\ \/__/ \:\  \    \:\  \
      \::/  /      |:/:/  /    /:/  /       \:\ \:\__\        \::/  /   \:\ \:\__\    \:\  \    \:\  \
      /:/  /       |::/  /     \/__/         \:\/:/  /        /:/  /     \:\ \/__/     \:\  \    \:\  \
     /:/  /        /:/  /                     \::/  /        /:/  /       \:\__\        \:\__\    \:\__\
     \/__/         \/__/                       \/__/         \/__/         \/__/         \/__/     \/__/
`,
		`
  ____  ____   ______  _____ __ __    ___  _      _     
 /    ||    \ |      |/ ___/|  |  |  /  _]| |    | |    
|  o  ||  _  ||      (   \_ |  |  | /  [_ | |    | |    
|     ||  |  ||_|  |_|\__  ||  _  ||    _]| |___ | |___ 
|  _  ||  |  |  |  |  /  \ ||  |  ||   [_ |     ||     |
|  |  ||  |  |  |  |  \    ||  |  ||     ||     ||     |
|__|__||__|__|  |__|   \___||__|__||_____||_____||_____|
`,
		`
   _____          __   _________.__           .__  .__   
  /  _  \   _____/  |_/   _____/|  |__   ____ |  | |  |  
 /  /_\  \ /    \   __\_____  \ |  |  \_/ __ \|  | |  |  
/    |    \   |  \  | /        \|   Y  \  ___/|  |_|  |__
\____|__  /___|  /__|/_______  /|___|  /\___  >____/____/
        \/     \/            \/      \/     \/
`,
		`
   ___        __  ______       ____
  / _ | ___  / /_/ __/ /  ___ / / /
 / __ |/ _ \/ __/\ \/ _ \/ -_) / / 
/_/ |_/_//_/\__/___/_//_/\__/_/_/
`,
		`
 ______     __   __     ______   ______     __  __     ______     __         __        
/\  __ \   /\ "-.\ \   /\__  _\ /\  ___\   /\ \_\ \   /\  ___\   /\ \       /\ \       
\ \  __ \  \ \ \-.  \  \/_/\ \/ \ \___  \  \ \  __ \  \ \  __\   \ \ \____  \ \ \____  
 \ \_\ \_\  \ \_\\"\_\    \ \_\  \/\_____\  \ \_\ \_\  \ \_____\  \ \_____\  \ \_____\ 
  \/_/\/_/   \/_/ \/_/     \/_/   \/_____/   \/_/\/_/   \/_____/   \/_____/   \/_____/
`,
		`
 ________  ________   _________  ________  ___  ___  _______   ___       ___          
|\   __  \|\   ___  \|\___   ___\\   ____\|\  \|\  \|\  ___ \ |\  \     |\  \         
\ \  \|\  \ \  \\ \  \|___ \  \_\ \  \___|\ \  \\\  \ \   __/|\ \  \    \ \  \        
 \ \   __  \ \  \\ \  \   \ \  \ \ \_____  \ \   __  \ \  \_|/_\ \  \    \ \  \       
  \ \  \ \  \ \  \\ \  \   \ \  \ \|____|\  \ \  \ \  \ \  \_|\ \ \  \____\ \  \____  
   \ \__\ \__\ \__\\ \__\   \ \__\  ____\_\  \ \__\ \__\ \_______\ \_______\ \_______\
    \|__|\|__|\|__| \|__|    \|__| |\_________\|__|\|__|\|_______|\|_______|\|_______|
                                   \|_________|
`,
		`
 ▄▄▄       ███▄    █ ▄▄▄█████▓  ██████  ██░ ██ ▓█████  ██▓     ██▓    
▒████▄     ██ ▀█   █ ▓  ██▒ ▓▒▒██    ▒ ▓██░ ██▒▓█   ▀ ▓██▒    ▓██▒    
▒██  ▀█▄  ▓██  ▀█ ██▒▒ ▓██░ ▒░░ ▓██▄   ▒██▀▀██░▒███   ▒██░    ▒██░    
░██▄▄▄▄██ ▓██▒  ▐▌██▒░ ▓██▓ ░   ▒   ██▒░▓█ ░██ ▒▓█  ▄ ▒██░    ▒██░    
 ▓█   ▓██▒▒██░   ▓██░  ▒██▒ ░ ▒██████▒▒░▓█▒░██▓░▒████▒░██████▒░██████▒
 ▒▒   ▓▒█░░ ▒░   ▒ ▒   ▒ ░░   ▒ ▒▓▒ ▒ ░ ▒ ░░▒░▒░░ ▒░ ░░ ▒░▓  ░░ ▒░▓  ░
  ▒   ▒▒ ░░ ░░   ░ ▒░    ░    ░ ░▒  ░ ░ ▒ ░▒░ ░ ░ ░  ░░ ░ ▒  ░░ ░ ▒  ░
  ░   ▒      ░   ░ ░   ░      ░  ░  ░   ░  ░░ ░   ░     ░ ░     ░ ░   
      ░  ░         ░                ░   ░  ░  ░   ░  ░    ░  ░    ░  ░
`,
		`
 ▄▄▄·  ▐ ▄ ▄▄▄▄▄.▄▄ ·  ▄ .▄▄▄▄ .▄▄▌  ▄▄▌  
▐█ ▀█ •█▌▐█•██  ▐█ ▀. ██▪▐█▀▄.▀·██•  ██•  
▄█▀▀█ ▐█▐▐▌ ▐█.▪▄▀▀▀█▄██▀▐█▐▀▀▪▄██▪  ██▪  
▐█ ▪▐▌██▐█▌ ▐█▌·▐█▄▪▐███▌▐▀▐█▄▄▌▐█▌▐▌▐█▌▐▌
 ▀  ▀ ▀▀ █▪ ▀▀▀  ▀▀▀▀ ▀▀▀ · ▀▀▀ .▀▀▀ .▀▀▀
`,
		`
.______  .______  _____._.________.___.__  ._______.___    .___    
:      \ :      \ \__ _:||    ___/:   |  \ : .____/|   |   |   |   
|   .   ||       |  |  :||___    \|   :   || : _/\ |   |   |   |   
|   :   ||   |   |  |   ||       /|   .   ||   /  \|   |/\ |   |/\ 
|___|   ||___|   |  |   ||__:___/ |___|   ||_.: __/|   /  \|   /  \
    |___|    |___|  |___|   :         |___|   :/   |______/|______/
`,
	}
)

func GetBanner(policy string) (banner string) {
	switch policy {
	case PolicyDefault:
		banner = Banner
		break
	case PolicyRandom:
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(BannerList))
		banner = BannerList[index]
		break
	default:
		if index, err := strconv.Atoi(policy); err == nil && index < len(BannerList) {
			banner = BannerList[index-1]
		} else {
			banner = Banner
		}
	}

	return
}
