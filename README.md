# go-dj
daemon for automating audio programming 

## Usage
This code was developed for automating content for an AM radio station; however, it can be installed on any Linux device with an audio output.

```azure
NAME:
Go DJ - Daemon that schedules audio programming content

USAGE:
    [global options] command [command options] [arguments...]

VERSION:
    0.0.1

AUTHOR:
    Jeremiah Miller

COMMANDS:
   start, s  start
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
--help, -h     show help
--version, -v  print the version

```

### Start Command 

```azure
NAME:
    start - start

USAGE:
   starts the daemon based on the config

OPTIONS:
   --random      Start your radio station w/ randomized schedule
   --pod-oldest  podcasts will play starting with the oldest first
   --pod-random  podcasts will play in a random order
   

```

## Dependencies
Under the hood, this uses [beep](https://github.com/faiface/beep) to play the files. That package relies on [Oto](https://github.com/hajimehoshi/oto)
which has dependencies that may need to be installed depending on the system you're running go-dj off of.
[MPV](mpv.io) is used to support web radio streams. Ideally this gets removed by making a package for dealing with web streams, but is not considered necessary to reach the feature complete milestone. 

### Incompatibility Warning
MPEG 2.5 is not supported because of underlying go-mp3 dependency.


## Supported Content

### Local Files
Local files are supported and the config assumes that they are located in the same directory as the deamon.

### Local Folder
Folders are also supported and make the same directory assumptions as local files.

### Web Radio 
A web radio station can be supplied via URL in the config

## Roadmaps
### [Feature Complete](https://github.com/jmillerv/go-dj/issues?q=is%3Aopen+is%3Aissue+milestone%3A%22Feature+Complete%22)
- [x] Podcast Support
- [x] Web Radio Station Support
- [x] Local File Support
- [x] Randomized Programming
- [x] Scheduled Programming
- [x] Local Folder support
- [x] Stop signal
- [x] WAV, OGG, FLAC Support
- [x] Timeslots in config
- [x] Played podcast cache
- [ ] Unit Tests

### [Nice To Have](https://github.com/jmillerv/go-dj/milestone/2)
- [ ] UI for scheduler & running the daemon
- [x] Pipeline 
- [ ] User Manual - Printable for workshops
- [ ] Browsable Documentation 
- [ ] Automated environment setup
- [ ] Read file metadata 
- [ ] Remote file support

### [Longshots](https://github.com/jmillerv/go-dj/milestone/3)
- [ ] Plugins for consuming additional types of audio 
- [ ] [Funkwhale/Subsonic API integration](https://docs.funkwhale.audio/api.html)

## Config Setup

go-dj consumes a file called `config.yml` the daemon searches for this file in the root of the directory
that the go-dj binary is stored in.

### Config Example

config.yml
```

version: 1.0.0
content:
  CheckInterval: 10m 
  PlayedPodcastTTL: 730h 
  Programs:
    - Name: "gettysburg10"
      Type: "file"
      Source: "./static/gettysburg10.wav"
      Timeslot:
        Begin: "11:00PM"
        End: "11:30PM"
```

#### Config Explained

The yaml is loaded into a struct with viper. 

`PlayedPodcastTTL` this is the expire time for the played podcast cache. It defaults to about a month. If an episode
attempts to play a second time within a month, it will be skipped. The duration is done hours, represented by `h` ex `10h`. It could
also be done in minutes, seconds, milliseconds, and nanoseconds. I did not test and am not sure days, weeks, months would 
work. 

`CheckInterval` is the duration for how long the daemon pauses between checking the programs for their timeslots. 


## Content Types

go-dj recognizes four types of content `file`, `folder`, `podcast`, `web_radio`


### File
The `file` type is intended to be a local file.

go-dj searches for files from the root of where the binary is stored. The LocalFile struct doesn't have a URL attached
but does have a path. When adding a LocalFile type to the programs be sure to use `Path` instead of `Source`. This should
be cleaned up so that source can stand in for URL/Path/etc but I haven't abstracted that yet.

### Podcast
The `podcast` file is for podcasts with published RSS feeds. go-dj will parse the RSS and select from there.

The podcast type defaults to playing newest. If the oldest or random flags are passed, podcasts will play in those orders. 
An improvement on this would be to allow this to be set on a per podcast basis in the config.yml. 

#### Played Episodes

go-dj implements a cache to check if a podcast episode has already been played. The cache defaults to a month and can be reset in
the config.yml. 

### Web Radio
The `web_radio` file is able to take in a web radio address and play it through your go-dj.

## Feature Requests 
I've built this out for my specific use case and released it to the public when I considered it feature complete. Suggestions are welcome for adding additional features.

### Paid Development
In addition, I will do paid development work on this project. Should you want a feature added, reach out at <insert_email>. Quotes start at $300 and the request must fall within the scope and vision of the project.