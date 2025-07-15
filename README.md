Shell Time
----------

[<img src="https://xn--gckvb8fzb.com/images/chatroom.png" width="275">](https://xn--gckvb8fzb.com/contact/)

*Shell Time* is for your shell, what *Screen Time* is for your phone!

> *Shell Time* lets you know how much time you and your kids spend on CLIs, 
> TUIs, and more. This way, you can make more informed decisions about how you 
> use your terminals.

*Shell Time* shows you your top most used commands, the (rough) amount of time 
you spend in your terminal per day and the hours you seem to be most drawn to 
the command line! It can give you interesting insights into your shell usage and 
remind you of long forgotten tools.

TODO: GIFs


## Installation

Either download a build from the releases page or clone this repository and run:

```sh
go build
```

or

```sh
go install
```


## Configuration


### Zsh

Nothing to configure!


## Bash

TODO


## Usage

```sh
shell-time
```

Example output:

```sh

=== YOUR TOP 10 COMMANDS ===
 1. vim (2198 times)
 2. cd (1757 times)
 3. rm (1132 times)
 4. mv (1117 times)
 5. find (1115 times)
 6. ls (888 times)
 7. ga (729 times)
 8. rg (672 times)
 9. cat (655 times)
10. git (605 times)

=== LONG FORGOTTEN COMMANDS ===
 1. dmesg\
 2. head
 3. ks
 4. freecad
 5. lokinet-bootstrap
 6. uuidgen
 7. mbe
 8. docnf
 9. gpoat
10. 27*100

=== MOST PRODUCTIVE HOURS ===
 1. 1:00 (421 commands fired)
 2. 2:00 (188 commands fired)
 3. 21:00 (1643 commands fired)
 4. 22:00 (1216 commands fired)
 5. 0:00 (1048 commands fired)

On average you ran commands on the shell for about 19 minutes per day.

```

