# advent-of-code

This repository contains my Go solutions for [Advent of Code] programming
challenges, inspired and forked from [[https://codeberg.org/derat/advent-of-code]], along with 
[advent.sh](./advent.sh), a shell script that makes various common tasks easier.

[Advent of Code]: https://adventofcode.com/

## Usage

In order for the library code to be able to download input,
`$REPO_DIR/.cookie` should contain the value of the `session` cookie
that gets set for the `.adventofcode.com` domain after authenticating with the
website. The session needs to be updated every year or two. You can set it using `advent.sh cookie`command.

I add this line to .bashrc
`alias advent="<$REPO_DIR>/advent.sh"`

This lets me run various commands like the following:

*   `advent today` - Move to the directory for today's puzzle.
*   `advent web` - Open today's puzzle in a web browser.
*   `advent run` - Run `main.go` in the current directory with real input.
*   `advent stdin <example.txt` - Run `main.go` with other input.
*   `advent input` - Gets today's input.
*   `advent cookie <cookie>` - Sets session cookie.
*   `advent save` - Run `main.go` and save its output under `answers/`.
*   `advent check` - Run `main.go` and compare its output against saved output.

## Copyright

The original puzzles and corresponding text (portions of which are sometimes
quoted within comments in my solutions) are copyrighted by Advent of Code. From
<https://adventofcode.com/about>:

> Advent of Code is a registered trademark in the United States. The design
> elements, language, styles, and concept of Advent of Code are all the sole
> property of Advent of Code and may not be replicated or used by any other
> person or entity without express written consent of Advent of Code. Copyright
> 2015-2022 Advent of Code. All rights reserved.
>
> You may link to or reference puzzles from Advent of Code in discussions,
> classes, source code, printed material, etc., even in commercial contexts.
> Advent of Code does not claim ownership or copyright over your solution
> implementation.
