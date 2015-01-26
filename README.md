# Shark Sandwich

## Git Game Server Repo

An external git repo has been created on Google cloud compute instance for this app to use as a 'game server'. This server is wide open with no auth required.

### To Test

 * git clone http://104.154.43.42/git/shark-sandwich-game-server

### To launch game

* When the game first launches, it will ask you for a folder path to a game repo. You can enter `testgame` here and it will use that folder in your local dir. I have added this folder to the gitignore.
* After you enter that, it will see that it is not a valid game location and it will ask you for a repo url. Enter the repo url above. It will clone it into that folder.
* Then it will ask you for a name. Enter whatever you want here.

The game should not be setup to play. When you run it later, you can enter the same folder path `testgame`, or if you run the game from within the `testgame` folder, it will load up straight away. You previously created character will be remembered.

Run `resetgame.sh` script to clear out repo.

## Introduction

*Shark Sandwich* is a text based game with a new twist. As your hero adventures, you encounter monsters and make discoveries. Expeirience the world through a command line interface, pitting your hero against the environment and eventually other players.

Every moment of your character is saved away; moments in time etched in the stone tablets of history... until you smash them to bits.

The game records your hero's timeline. It also records your freind's timelines as well. Eventually, you'll be able to travel back in time to attack your friend's hero. It's a sketchy endeavor. Will you win, or lose? How will that alter the course of history?

If your character dies, your present day character becomes another descendent in the long line of your hero's heritage. You can become stronger or even weaker based on the outcome!

### Feature List

- [x] Create a hero
- [x] Go on an adventure
- [x] Record events for each hero
- [x] Persist state
- [x] Resume state
- [ ] Interact with friends
- [ ] PvP
- [ ] Rewrite history after death
- [ ] Make current character a new decendent from altered time
- [ ] Loot
- [ ] Gear
- [ ] Character building

