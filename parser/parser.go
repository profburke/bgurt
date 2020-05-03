package parser

// This file will include function(s) to parse tags in overtext strings,
// replacing the tag with the appropriate value.

// In https://boardgamegeek.com/thread/501653/bgg-randomizer-periodically-randomize-your-desktop/page/1
// supported tags are

/*


@Bail Organa
Posted,Mar 3, 2010 (edited)
BGG Randomizer HOWTO

Installation
Download the application from BGG Randomizer and simply unzip it to a folder. Run the executable file and you're good to go.
If you're updating the application to a newer version, simply unzip the update to the same folder were you had placed the previous version. Don't forget to close the application if it is running before doing so.

Generic configuration
The easiest way to use the application is simply to:
1Insert your BGG username and password.
2Configure the wallpaper and/or microbadge and/or (uber)geekbadge (image and/or overtext) and/or avatar overtext settings as described below.
3Tick the Auto update checkbox. That option will make the program run when you login to Windows (in most cases, when you turn your computer on) and proceed to periodically randomize your BGG profile according to your configuration. If you use this Auto update option, don't forget to tick the appropriate Randomize wallpaper/geekbadge/microbadge/avatar overtext according to whatever you want. After doing so, you can basically forget about the program.

Desktop wallpaper randomization
If you really like BGG's images, like me!, you can set the application to periodically change your desktop's wallpaper. Simply chose the appropriate options (Hot/Recent, Period, Gallery, etc) and, well, that's it!

Microbadge randomization
Just tick which microbadge slots you wish to be randomized; if you don't want a certain microbadge to be placed in any of your microbadge slots, simply retrieve the current configuration with the Retrieve Config button and untick the ones you don't want to be displayed. You may force a specific microbadge to be always set to a specific slot; simply select the badge in the list and click the properties button (small button with a pencil).

(Uber)Geekbadge randomization, including geekbadge overtext
First of all, if you don't have a (uber)geekbadge and/or overtext, don't bother using the appropriate option. It won't work! external image
To use, simply:
1 For the purposes of the application, the geekbadge image and overtext is considered a pair, so first of all you should select what is to be randomized, according to what you want and have: if the geekbadge image, its overtext or both the image and overtext.
2 Select the Ubergeekbadge or Geekbadge option.
3 Configure the colors and various options more or less as you would on BGG.
3 Select how the geekbadge will be randomized, using "dynamic" text, static text and or images, as described next.

Geekbadge random text and/or image for geekbadge image and overtext
Here you can insert a list of text lines and/or images, including some tags that are replaced with special information, for use in your badge and/or overtext. If you are using only the badge image or overtext, you can simply ignore the options for the other. Overtext should be reasonably simple to configure and use, and please note that all special tags defined below are also valid for it. As for the badge image configuration: Neither text or image is mandatory, which allows you several options, such as having a list of ubergeekbadge badge images, a list of phrases that are superimposed on the default or specific image or a mix of the two.

A few configuration examples (in which the default image has a camel on it):
no caption, "THERE IS NO TRY" text, using default image
no caption, "DO OR DO NOT" text, using default image
no caption, "GIVEN:[THG][LF]RECEIVED:[THR]", thumbs.png image
no caption, "Essen in [C] days!", essen.png image
no caption, "WUs:[FUW][LF]Score:[FUS]", folding.png image

Empty text is perfectly possible, such as:
no caption, no text, ubermicrobadge.png image

Regular geekbadges can also be used, but they're naturally more limited. A simple example:
no caption, "TR:[THR]TG:[THG]"

Avatar overtext randomization
Should be reasonably easy to use, just add new entries to the list and your avatar overtext will be updated according to those entries. All dynamic text tags - twitter, folding@home, BGG gaming data, etc - can also be used here.

Text tags
There are several tags that you can use while specifying text, which are:
[C] - Countdown to a date; the tag is replaced by the number of days of whatever you set the countdown text to, according to the number of missing days. You you can also specify a specific text that will be displayed when the date is reached.
[THG] - Your total number of given thumbs
[THR] - Your total number of received thumbs
[GO] - Your total number of games owned
[GP] - Your total number of games played
[30] - Your total number of games played in the last 30 days
[30e] - Your total number of games played in the last 30 days, excluding expansions
[30E] - Your total number of expansions played in the last 30 days
[LF] - Force line feed

Twitter text tags
[TW] - Your latest tweet
[TWD] - Date of latest tweet
In order to use these tags, you must configure your twitter username in the "Extra tag data" section.

Folding@Home text tags
[FTS] - Your team's score
[FTW] - Your team's WUs
[FTR] - Your team's rank
[FUR] - Your overall rank
[FUT] - Your rank in your team
[FUS] - Your score
[FUW] - Your WUs


Here are some more possibilites:
[MB###] -- microbadge name ###
[MBOM###] -- mouseover text for ###
[MBI###] -- the microbadge image (generate image and include mb)


*/

// TODO: should also be able to parse text for geek badges
// -- maybe move this out of overtext library and into utilities?

func Parse(text string) (result string) {
	return "dude"
}
