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
- [ ] Podcast Support
- [x] Web Radio Station Support
- [x] Local File Support
- [x] Randomized Programming
- [ ] Scheduled Programming
- [x] Local Folder support
- [ ] Remote file support
- [x] Stop signal
- [x] WAV, OGG, FLAC Support
- [ ] Timeslots in config
- [ ] Unit Tests

### Extras
- [ ] UI for scheduler & running the daemon
- [ ] Pipeline 
- [ ] User Manual - Printable for workshops
- [ ] Docs

### Longshots
- [ ] Plugins for consuming additional types of audio 
- [ ] [Funkwhale API integration](https://docs.funkwhale.audio/api.html)
- [ ] Remove MPV dependency
- [ ] Remove beep & oto dependency

## Config Setup

go-dj consumes a file called `config.yml` the daemon searches for this file in the root of the directory
that the go-dj binary is stored in.

### Config Example

config.yml
```
version: 0.0.1
content:
  Programs:
    - Name:
      Type:
      Slot:
      Source:
```


## Timeslots
early: 5 AM - 7 AM  
morning: 7 MA - 8 AM  
breakfast: 8 AM - 11 AM  
midmorning: 11 AM - 2 PM  
afternoon: 2PM - 5 PM  
commute: 5PM - 7 PM  
evening: 7 PM - 11 PM  
late: 11 PM - 2 AM  
overnight: 2 AM - 5 AM  

## Files

### Local Files
go-dj searches for files from the root of where the binary is stored. The LocalFile struct doesn't have a URL attached
but does have a path. When adding a LocalFile type to the programs be sure to use `Path` instead of `Source`. This should
be cleaned up so that source can stand in for URL/Path/etc but I haven't abstracted that yet.

### Supported File Types
At the moment go-dj only supports mp3 files for local and remote files.

## Types

go-dj recognizes four types of content: `announcement`, `file`, `podcast`, `web_radio`

The `announcement` type could be pulled from a file or a URL, but I wanted a distinct type for content interruptions.
Since this would be hyper-local AM radio, I was thinking community service announcements, how to get involved, or donation requests.
The intent is that the `announcement` struct will be treated with more leniency within the automation.

The `file` type is intended to be a local file; however, I see the use case for being able to pull files from URLs and will
eventually add that functionality

The `podcast` file is for podcasts with published RSS feeds. go-dj will parse the RSS and select from there.

The `web_radio` file is able to take in a web radio address and play it through your go-dj.
