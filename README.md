# go-dj
daemon for automating audio programming 

## Usage
This code was developed for automating content for an AM radio station; however, it can be installed on any Linux device with an audio output.

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

## Roadmap
### [Feature Complete](https://github.com/jmillerv/go-dj/issues?q=is%3Aopen+is%3Aissue+milestone%3A%22Feature+Complete%22)
- [x] Podcast Support
- [x] Web Radio Station Support
- [x] Local File Support
- [x] Randomized Programming
- [x] Scheduled Programming
- [x] Local Folder support
- [ ] Remote file support
- [x] Stop signal
- [x] WAV, OGG, FLAC Support
- [x] Timeslots in config
- [ ] Unit Tests

### [Nice To Have](https://github.com/jmillerv/go-dj/milestone/2)
- [ ] UI for scheduler & running the daemon
- [x] Pipeline 
- [ ] User Manual - Printable for workshops
- [ ] Docs
- [ ] Automated environment setup
- [ ] Read file metadata 

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
  Programs:
    - Name: "gettysburg10"
      Type: "file"
      Source: "./static/gettysburg10.wav"
      Timeslot:
        Begin: "11:00PM"
        End: "11:30PM"
```

## Files

### Local Files
go-dj searches for files from the root of where the binary is stored. The LocalFile struct doesn't have a URL attached
but does have a path. When adding a LocalFile type to the programs be sure to use `Path` instead of `Source`. This should
be cleaned up so that source can stand in for URL/Path/etc but I haven't abstracted that yet.

### Supported File Types
At the moment go-dj only supports local files.

## Types

go-dj recognizes four types of content `file`, `folder`, `podcast`, `web_radio`


### File
The `file` type is intended to be a local file; however, I see the use case for being able to pull files from URLs and will
eventually add that functionality

### Podcast
The `podcast` file is for podcasts with published RSS feeds. go-dj will parse the RSS and select from there.

The podcast type defaults to playing newest. If the oldest or random flags are passed, podcasts will play in those orders. 
An improvement on this would be to allow this to be set on a per podcast basis in the config.yml. 

### Web Radio
The `web_radio` file is able to take in a web radio address and play it through your go-dj.

## Feature Requests 
I've built this out for my purposes and released it to the public when I considered it feature complete. Suggestions are welcome for adding additional features and I will accept PRs that extend go-dj. In addition, I will do paid development work on this project. Should you want a feature added, reach out at <insert_email>. Quotes start at $300 and the request must fall within the scope and vision of the project.
