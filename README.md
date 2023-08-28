# go-dj
daemon for automating audio programming 

## Usage
This code was developed for automating content for an AM radio station; however, it can be installed on any Linux device with an audio output.

For the first time setup, I'd recommend checking out this [doc](https://github.com/jmillerv/go-dj/blob/main/docs/first_time_setup.md)

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

### Future Releases
Check the [milestones](https://github.com/jmillerv/go-dj/milestones) for what is planned and when it's expected to release. 

### [Nice To Have](https://github.com/jmillerv/go-dj/milestone/2)
This is the milestone where I throw tickets that I may pull into current release milestones or work on as quick, one offs.  

### [Longshots](https://github.com/jmillerv/go-dj/milestone/3)
This is board is where neat ideas get thrown, but realistically won't be implemented anytime soon. 

## Config Setup

go-dj consumes a file called `config.yml` the daemon searches for this file in the root of the directory
that the go-dj binary is stored in.

You can set an alternative file by setting the `GODJ_CONFIG_OVERRIDE` environment variable. If this is set, it will use 
the file named in the variable. More useful for development than production, IMO. 

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

**Note** thus far only .pls streams are provably working. An outstanding [issue](https://github.com/jmillerv/go-dj/issues/37) exists to dig into why some other 
streams don't work. 

## Feature Requests 


## Development Setup

Follow the steps in [development.md](https://github.com/jmillerv/go-dj/blob/main/docs/development.md) to get started.

### Linting 
Linting is performed by [golangci-lint](https://golangci-lint.run/)
