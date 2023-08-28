# Introduction

Thanks for taking a look and considering contributing to go-dj. This isn't an ambitious tool, but an endeavor to get people interested 
in local first music solutions and possibly [Part 15 of Title 47 CFR](https://en.wikipedia.org/wiki/Title_47_CFR_Part_15).

Following these guidelines will help you to report and fix bugs, as well as, submit and suggests features.

## Contributions 

I've built this out for my specific use case and released it to the public when I considered it feature complete. Currently, my focus is on fixing bugs I've found during usage and making the development experience better.
My hope is to get this into a state where most people, especially non-developers, could get a hold of go-dj and start experimenting with it. To that end,
improving documentation, bug triaging, or writing tutorials are all welcome contributions. Suggestions are also welcome for adding additional features, they just might not be priority.

### What Isn't Sought 

At this time, the goal of go-dj isn't to become a competitor to any other dj software. The goal is to be an easily configurable, lightweight tool for scheduling audio content to be played locally. Anything that expands beyond 
those goals likely won't be prioritized. 


## Repository Rules

> Responsibilities
> * Ensure cross-platform compatibility for every change that's accepted. Currently, this is Ubuntu/Debian Linux. 
> * Create issues for any major changes and enhancements that you wish to make.
> * Keep feature versions as small as possible, preferably one new feature per version.
> * Be welcoming to newcomers and encourage diverse new contributors from all backgrounds. 

## Your First Contribution

Unsure where to begin contributing to go-dj? You can start by looking through the [good first issue](https://github.com/jmillerv/go-dj/labels/good%20first%20issue) and the [help wanted](https://github.com/jmillerv/go-dj/labels/help%20wanted) tags:

**Good first issues**: issues which should only require a few lines of code, and a test or two.  
**Help wanted issues**: issues which should be a bit more involved than beginner issues.

If you have never submitted a pull request before, check out: http://makeapullrequest.com/ and http://www.firsttimersonly.com/. 
You can also check out this series: [How to Contribute to an Open Source Project on GitHub](https://egghead.io/series/how-to-contribute-to-an-open-source-project-on-github).

# Getting started

* Code coverage should not fall below 80 percent when code is submitted to the repository. 
* The code should be linted before making the pull request, else it will fail the pipeline. 

For something that is bigger than a one or two line fix:
    1. Create your own fork of the code
    2. Do the changes in your fork
    3. If you like the change and think the project could use it run through the following checklist:  
- [ ] Code linted   
- [ ] Unit tests written 
- [ ] Unit tests passing 
- [ ] Code works locally without issue

    
### Reporting Bugs & Submitting Issues

> When filing an issue, make sure to answer these five questions:
>
> 1. What version of Go are you using (go version)?
> 2. What operating system and processor architecture are you using?
> 3. What did you do?
> 4. What did you expect to see?
> 5. What did you see instead?


# Code review process
When a merge request is submitted, after it passes the pipelines, I will test and review the changes locally and make any necessary comments on the suggested changes. 

I endeavor to check on this project at least once a week. 

## Acknowledgements 
https://github.com/nayafia/contributing-template