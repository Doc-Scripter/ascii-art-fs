# Ascii-art Project

## Overview

Ascii-art is a command-line Go application that renders input text strings into ASCII art representations. Transform your simple text into an exciting visual that stands out in any plain text environment.

## Objectives

Ascii-art aims to:

- Receive input strings and convert them into graphic representations using ASCII characters.
- Support varied input including alphanumerics, spaces, special characters, and line breaks.
- Offer multiple ASCII art fonts such as shadow, standard, and thinkertoy.

## Technical Specifications

- **Language**: The entire project is implemented in Go (Golang).
- **Coding Standards**: Follow Go's best practices for a clean, maintainable codebase.
- **Unit Testing**: Includes a test to ensure the reliability and stability of the application.

## Banner Format

Ascii-art characters adhere to the following specifications:

- Fixed height of 8 lines.
- New line separation for each character representation.

## Usage
For you to use our project your have to clone the repository from gitea 

```shell
cd ~
git clone https://learn.zone01kisumu.ke/git/aadero/ascii-art.git
```

once you have cloned it you go to the terminal and enter to the directory where the program is contained.

```shell
cd ascii-art
```
To run Ascii-art, use the Go command-line interface with the following syntax:

```shell
go run . "Your desired string"
```

## Example

if we run 
```go
go run . "Hello\n" | cat -e
```
on the terminal we expect an ascii art to be printed for hello and new line i.e
```

 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
```

if you use characters which are out of range in ascii between 32 and 126 it print's out an error message displaying it's not an ascii character
```shell
go run . "さようなら" 
```
the output will be: "Not an ascii character"


## Banner file switch
If you look at the Resources file it contains a series of banner files that can be used to provide the various types of formats in which you want to print the word. We have handled a set of banner files including 
* standard
``` 
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
```

 * thinkertoy
```
                 $
o        o o     $
|        | |     $
O--o o-o | | o-o $
|  | |-' | | | | $
o  o o-o o o o-o $
                 $
                 $
```
* shadow
```
                                 $
_|                _| _|          $
_|_|_|     _|_|   _| _|   _|_|   $
_|    _| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $
```

## Contribution

Pull request are accepted
