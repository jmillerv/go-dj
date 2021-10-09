# go-dj
daemon for automating audio programming 

# Usage
The use case I built this for was automating content for an AM radio station I run around my house.

# Roadmap
- [ ] Announcements
- [ ] Podcast Support
- [ ] Web Radio Station Support
- [ ] Local File Support
- [ ] Randomized Programming
- [ ] Scheduled Programming

# Config Setup



## Config Example

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

Types