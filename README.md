# Shark Sandwich

## Git Game Server Repo

An external git repo has been created on Google cloud compute instance for this app to use as a 'game server'. This server is wide open with no auth required.

### To Test

 * git clone http://104.154.43.42/git/shark-sandwich-game-server
 * add a line to the README, checkin and push

### To launch game

* When the game first launches, it will ask you for a folder path to a game repo. You can enter `testgame` here and it will use that folder in your local dir. I have added this folder to the gitignore.
* After you enter that, it will see that it is not a valid game location and it will ask you for a repo url. Enter the repo url above. It will clone it into that folder.
* Then it will ask you for a name. Enter whatever you want here.

The game should not be setup to play. When you run it later, you can enter the same folder path `testgame`, or if you run the game from within the `testgame` folder, it will load up straight away. You previously created character will be remembered.

Run `resetgame.sh` script to clear out repo.
