# Running go-dj on a fresh machine

These instructions come from getting this working on a fresh install of [DietPi]() running on a [Raspberry Pi 3 B+]() There
is an [existing issue](https://github.com/jmillerv/go-dj/issues/26) to automate this process for ease of use.


## Install Go
1. mkdir ~/tmp
2. cd ~/tmp
3. wget https://go.dev/dl/go1.19.4.src.tar.gz
4. sudo tar -C /usr/local -xzf go1.19.4.src.tar.gz
5. add export PATH=$PATH:/usr/local/go/bin to .profile
6. source ~/.profile
7. go version
   returns  `go version go1.19.4 linux/arm64`

## Install dependencies
These are the dependencies that I installed during the process of troubleshooting. I need to go through and confirm what is/is not needed.

1. `sudo apt install libasound2 -y`
2. `sudo apt install alsa-utils -y`
3. `sudo apt-get install -y -qq libasound2-dev libssl-dev libpulse-dev libdbus-1-dev portaudio19-dev`
4. `sudo apt install build-essential`
5. `sudo apt install mpv` // this installs many of the previous packages, but I did this as step 5 when testing. 

## Clone Repo
1. `cd ~/`
2. `mkdir dev`
3. `git clone https://github.com/jmillerv/go-dj`

## Build
1. `cd ~/dev/go-dj`
2. `go build`
   3 chmod +x ./go-dj

## Add a config
1. cd ~/dev/go-dj
2. `mv config.example.yml config.yml`
3. populate the yaml with your desired content

## Configure your timezone
This will be different based on your OS and you may have done this during setup. I missed setting it to mine during the boot, so files weren't playing when I expected. 

## Configure audio outputs 
DietPi doesn't have the soundcards installed by default, I had to do this to get sound out of the 3.5mm jack. go-dj outputs to
wherever the default audio is sent to. 

### Test your speakers
The following command will test the speakers without playing loud noise
`speaker-test -t wav -c 6` 

To end the test press `ctrl+c`

## Run

./go-dj s

## Run on reboot
Should your SBC go down, you'll likely not want to have to reboot the daemon yourself. 

1. open crontab
`crontab -e`
2. add command to file 
`@rebott cd ~/dev/go && ./go-dj s`

## Killing the program
### Kill Signal
`Ctrl+C` will usually be enough to stop the program; however, there are known issues where it does not.

### Kill the process

Should you run into issues, knowing how to kill the program is useful

1. Get the process ID

`ps aux | grep go-dj`  
should yeild something that looks like  
`root       22328  0.0  1.2 1240808 12436 pts/0   Sl+  02:03   0:00 ./go-dj s`

2. stop the process with the procdess ID
   `kill -9 22328` 
