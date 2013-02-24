TODO List
=========
* Priority is to wrap most of the Allegro API and then refactor it to comply better with GO idioms.
* Upgrade to latest 5.x branch (Right now, any func after 5.0.3 is commented out).
* Get rid of parts that can be done better with GO core libraries (File I/O, Path?).
* API naming cleanup: I'm trying to simplify things by omitting prefixes like "al_" in function names but sometimes it causes inconsistent names. Come up with a good scheme.
* Use GO's test facilities.
* Port some allegro examples.
* Documentation? (Is it worth copying over the Allegro docs? How will it be maintained?).
* Check class encapsulations as it doesn't make sense for some of the function (SetMouseXY belongs to Display but SetMouseZ is independant).
* Idiomatic error handling integration?
* Orginize code
