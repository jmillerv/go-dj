# go-dj
daemon for automating audio programming 

## Usage
The use case I built this for was automating content for an AM radio station I run around my house.

## Roadmap
- [ ] Announcements
- [ ] Podcast Support
- [ ] Web Radio Station Support
- [ ] Local File Support
- [ ] Randomized Programming
- [ ] Scheduled Programming
- [ ] Local Folder support
- [ ] Remote file support

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
Early: 5 AM - 7 AM
Morning: 7 MA - 8 AM
Breakfast: 8 AM - 11 AM
Midmorning: 11 AM - 2 PM
Afternoon: 2PM - 5 PM
Commute: 5PM - 7 PM
Evening: 7 PM - 11 PM
Late: 11 PM - 2 AM
Overnight: 2 AM - 5 AM

## Local Files
go-dj searches for files from the root of where the binary is stored. The LocalFile struct doesn't have a URL attached
but does have a path. When adding a LocalFile type to the programs be sure to use `Path` instead of `Source`. This should
be cleaned up so that source can stand in for URL/Path/etc but I haven't abstracted that yet.

## Types

go-dj recognizes four types of content: `announcement`, `file`, `podcast`, `web_radio`

The `announcement` type could be pulled from a file or a URL, but I wanted a distinct type for content interruptions.

The `file` type is intended to be a local file; however, I see the use case for being able to pull files from URLs and will
eventually add that functionality

The `podcast` file is for podcasts with published RSS feeds. go-dj will parse the RSS and select from there.

The `web_radio` file is able to take in a web radio address and play it through your go-dj.