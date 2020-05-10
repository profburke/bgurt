
## BGURT - BoardGameGeek Graphical User Representation Toolkit

<!-- [![NPM Version][npm-image]][npm-url]
[![Build Status][travis-image]][travis-url]
[![Downloads Stats][npm-downloads]][npm-url] -->

Each user on [BoardGameGeek](https://boardgamegeek.com) (BGG) has a graphical user representation ([GUR](https://boardgamegeek.com/wiki/page/Graphical_User_Representation)) consisting of an avatar, a badge, and several microbadges. The libraries and programs in this project allow you to easily
change the displayed components of your GUR.

##### Motivation

A user's GUR can display up to five microbadges. However, many
(_most?_) users have more than five microbadges. BGG User [Bail Organa](https://boardgamegeek.com/user/Bail%20Organa) wrote a tool, [BGG Randomizer](https://boardgamegeek.com/thread/501653/bgg-randomizer-periodically-randomize-your-desktop), that randomly switches out the displayed microbadges (_and eventually included a lot more functionality_).

However, BGG Randomizer is Windows-only, and I don't use a Windows machine. 

Also, it's closed source and its creator has not been responding for several years.  It probably (_again, see closed source_) works by screen scraping. Thus, as BGG has evolved, the tool has slowly been losing functionality.


<!-- #### Build Status

badges from Travis or what not (does this need to be a separate section?) -->

<!-- #### Code Style

badges ??? see: https://medium.com/@meakaakka/a-beginners-guide-to-writing-a-kickass-readme-7ac01da88ab3 -->

## Features

The toolkit is designed using a layered approach.

- **Low-level libraries**: these provide functionality to get and set a user's avatar, geekbadge, microbadges, and overtext.

- **Command line utilities**: programs that function as blocks in a piped chain of commands. The intention is that these utilities are  combined to implement custom workflows. 

 The command line utiliities include _av-fetch_, _av-set_, _gb-fetch_, _gb-set_, _mb-fetch_, _mb-fetchslot_, _mb-set_, _mb-setslot_, _ot-fetch_, and _ot-set_.
 
- **Command line programs**:  these are more comprehensive than the previously mentioned utilities. 

 For example, _av-randomizer_ randomly picks an image from a specified directory of images and sets your avatar to that image, whereas the utility, _av-set_, sets your avatar to the specified file.

 The command line programs include _av-randomize_, _gb-randomize_, _mb-randomize_, and _ot-randomize_.

- **Graphical User Interface programs**: These programs implement all the functionality in a desktop context.  _In development._

- **Text User Interface programs**: Similar to the above, but with a text-based UI. TUI programs are particularly useful if you are logged in remotely to the system on which they are installed. _In development._

- **AWS Lambdas and Step Functions**: these functions allow you to run your GUR manipulation workflow _in the cloud_, using AWS's serverless functionality.

Binaries are available for macOS, Windows, and Linux (Specifically, Raspberry Pi; other Linux distributions are possible). AWS Lambda functions are currently only available as source code.

## How to Use

(_As mentioned above TUI and GUI programs are not available yet. This section will be updated when they are ready._)

There are several levels of tools available. The following instructions on using the programs are arranged by how much you want to bother with technical details.

In all cases, however, there are a few steps necessary to install the programs. Details on installation are found in the next section.

##### I want the tools requiring the least amount of bother: mb-randomize, av-randomize, gb-randomize, and ot-randomize

To update your displayed microbadges, first use the `mb-fetch` program to download a list of all your microbadges as follows:

```
mb-fetch > badges.json
```

After the first run of `mb-fetch`, you only need to re-run it after purchasing new microbadge(s). Then run 

```
mb-randomize badges.json
```

To update your avatar, run

```
av-randomize <avfolder>
```

where `<avdfolder>` is the name of a folder that contains one or more image files you would like to use as your avatar. The image files need to be either in GIF, JPG, or PNG format and must be 64x64 pixels or smaller.

To update your overtext, run

```
ot-randomize <overtextfile>
```

where `<overtextfile>` is a file containing the different possible overtexts you would like to display. Currently you need to specify both the avatar and geekbadge overtexts. The file should be formatted as a JSON array where each entry is a pair of overtexts. For example:

```
[
    {
        "Avatar": "Can sour cream go bad?",
        "Badge": "Go Speed Racer!"
    },
    {
        "Avatar": "Are Santa's helpers called subordinate clauses?",
        "Badge": "There's a fine line between a superpower and a chronic medical condition."
    },
    {
        "Avatar": "What happens if you get scared half to death twice?",
        "Badge": "You don't build a 100 foot robot out of nothing"
    }
]
```

To update your geekbadge, run 

```
gb-randomze <geekbadgefolder>
```

where `<geekbadgefolder>` is a folder containing files that describe different geekbadges. These geekbadge files are in JSON format. An example is as follows:

```
{
   "OuterBorder":{"R":106,"G":90,"B":205,"A":255},
   "InnerBorder"{"R":138,"G":43,"B":226,"A":255},
   "BarPosition":40,
   "LeftBox":{
          "Text":"Play",
          "Background":{"R":85,"G":107,"B":47,"A":255},
          "TextColor":{"R":152,"G":251,"B":152,"A":255},
          "TextStart":4
   },
   "RightBox":{
          "Text":"Always",
          "Background":{"R":152,"G":251,"B":152,"A":255},
          "TextColor":{"R":85,"G":107,"B":47,"A":255},
          "TextStart":44
   }
}
```

Tools to easily specify geekbadge descriptions are not yet developed, so currently you will have to create them by hand.


**NOTE:** The functionality of these programs will likely change.

--

##### I'd like some customization please: mb-set, mb-fetch, av-set, ot-set, gb-set, etc.

This next set of programs sits at a slightly lower level than the previous set. They are designed mainly to be combined using command line pipes. Alternatively they can be invoked by scripts written in other langauges (_such as Python, Lua, Ruby, etc., all of which have some sort of_ exec _functionality_).

```
av-fetch <filename>
```

Fetches your current avatar and saves it to the file named `<filename>`.

```
av-set <filename>
```

Sets your avatar using the image in the file named `<filename>`. This image must be either a GIF, JPG, or PNG and must be a maximum of 64x64 pixels.

```
gb-fetch
```

Fetches your current geekbadge and writes a JSON representation to standard out.

```
gb-set <filename>
```

Sets your geekbadge. The file, `<filename>` should contain a JSON representation of your geekbadge. See above for an example of the format.

```
mb-fetch
```

Retrieves all your microbadges and writes them to standard out as a JSON data structure (and array of objects).

```
mb-fetchlot <slotnumber>
```

Not actually implemented. Also not sure this is needed&mdash;probably should be wrapped into `mb-fetch` using an additional command line parameter.

```
mb-set <badgedID1> <badgeID2> ... <badgeID5>
```

Run this program passing in 5 badge IDs to updated your microbadges. 

```
mb-setslot --slot <number> --microbadge <badgeID>
```

Use the `--slot` flag to specify slot number and the `--microbadge` flag to specify a microbadge ID. Sets the specified slot to the given microbadge. Probably will merge this in with the `mb-set` program using appopriate flags.

```
ot-fetch [--json]
```

Retrieves your overtext and writes to standard out in a simple text format. If the `--json` switch is given, writes it as a JSON object.

```
ot-set --avatar "avatar overtext" --badge "badge overtext"
```

Sets your overtext to the values given on the command line.



##### I'm into gory technical details: library files....

At the bottom of the stack are the various library files: avatar.go, bggclient.go, geekbadge.go, microbadge.go, overtext.go, parser.go (_currently just a shell_) and a few other utility files.

Use these libraries to build your own command-line, TUI, or GUI tools.


### Running on a schedule

If you want to change your microbadges periodically, e.g. every day at 6pm, you will need a way to run the command line tools on a schedule. For macOS and Linux, your best bet is to use `cron`. 

A quick internet seach will turn up many articles on using cron, here's a reasonably good one: [How to use cron in Linux](https://opensource.com/article/17/11/how-use-cron-linux).

On Windows, you can make use of Scheduled Tasks. This [article](https://www.digitalcitizen.life/how-create-task-basic-task-wizard) was written relatively recently. I have not verified it since I don't easily have access to a Windows machine. I'll update as soon as possible. In the meantime, if you are a Windows user, I encourage you to propose an update to this README if there are any inaccuracies in the linked document.


_**For more examples and usage, please refer to the [Wiki](https://github.com/profburke/bgurt/wiki).**_


## How to Install

Installation consists of moving the executables to a directory in your executable path and creating a configuration file in the appropriate directory. 

The executable path is a set of directories on your computer where programs are located. For more details, see this [answer](https://superuser.com/questions/284342/what-are-path-and-other-environment-variables-and-how-can-i-set-or-use-them) on [superuser.com](https://superuser.com).

The configuration file is not strictly necessary because you can specify the needed information to each program using environment variables. However, a configuration file is more convenient.

##### About the Passhash

In order to manipulate your GUR, the BGURT tools need to authenticate with the BoardGameGeek website. In order to do so, they send your username and _a hash_ of your passowrd with each request.

Note that this is *not* your password. Instead, it is an encrypted form of your password. Moreover, it is not possible (_or, at least, extremely difficult_) to determine your password based on the password hash. There are two reasons bgurt does not use your password: 1) I don't want to know your password, and 2) I don't know what the algorithm is that BGG uses to hash your password and they require the hash be sent with each HTTP request. Since bgurt manipulates your GUR using HTTP requests, it needs the hash.

##### How to get Your Passhash

Take a look in the cookies your web browser sends to boardgamegeek.com. Details are beyond the scope of this README file, so please do a web search if you need help retrieving your cookies.

##### Configuration File Format

The configuration file should be named `config.toml` and its contents are as follows:

```bash
username = 'YOUR_USER_NAME'
passhash = 'YOUR_PASSWORD_HASH'
```


##### For macOS

A good location for your executables is `/usr/local/bin`, but any directory specified by your `PATH` environment variable is fine. The configuration file should go in `~/Library/Application Support/bgurt`.

##### For Linux

Executables should go in `/usr/local/bin` or any other directory specified by your `PATH` environment variable.  The configuration files should go in `~/.config/bgurt`.

##### For macOS and Linux

If you wish to use environment variables instead of a configuration file, set them with the `export` command.

```sh
export BGGUSERNAME=your_user_name
export BGGPASSHASH=your_password_hash
```

##### For Windows

Locations for executables - TBD; Location for configuration file - TBD; How to set environment variables - TBD.

## How to Contribute

Thank you for taking the time to contribute!

There are many ways to contribute in addition to submitting code. Bug reports, feature suggestions, a logo for the project, and improvements to documentation are all appreciated.

All contributors are expected to abide by BoardGameGeek's [Community Rules](https://videogamegeek.com/community_rules).

##### Bug Reports and Feature Suggestions

Please submit bug reports and feature suggestions by creating a [new issue](https://github.com/profburke/bgurt/issues/new). If possible, look for an existing [open issue](https://github.com/profburke/bgurt/issues) that is related and comment on it.

When creating an issue, the more detail, the better. For bug reports in partciular, try to include at least the following information:

* The application version
* The operating system (macOS, Windows, etc) and version
* The expected behavior
* The observed behavior
* Step-by-step instructions for reproducing the bug
* Screenshots for GUI issues


##### Pull Requests

Ensure the PR description clearly describes the problem and solution. It should include the relevant issue number, if applicable.


##### Documentation Improvements

Preferably, submit documentation changes by pull request. However, feel free to post your changes to an [issue](https://github.com/profburke/bgurt/issues/new) or send them to the project team.

It would be nice to have a site for this project that we could host on Github Pages. If you have an interest in helping build that, please let me know.


<!-- ### Credits -->



## License

This project is licensed under the BSD 3-Clause License. For details, please read the [LICENSE]() file.

	
<!-- Markdown link & img dfn's -- >
[npm-image]: https://img.shields.io/npm/v/datadog-metrics.svg?style=flat-square
[npm-url]: https://npmjs.org/package/datadog-metrics
[npm-downloads]: https://img.shields.io/npm/dm/datadog-metrics.svg?style=flat-square
[travis-image]: https://img.shields.io/travis/dbader/node-datadog-metrics/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/dbader/node-datadog-metrics
[wiki]: https://github.com/yourname/yourproject/wiki -->
