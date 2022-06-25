# The Speech Only Toaster!
### by Giovanni D'Amico

*"toasts your Speech Only Tsv's to a crisp!"*

Speech Only Toaster is a small .exe I used for collecting Speech Only TSV's for audit. its designed to be chucked into your PATH and ran in CLI inside of the working folder.

## How to use

Inside of the working folder, call the application and provide `2` arguments. a `firstname` and a `lastname`
```bash
$ speechOnlyToaster giovanni damico
```

soToaster will use your arguments to generate a folder with your name and the current date with every `curated_speech_only.tsv` of every applicable `session-` folder in your directory.

**if there is already a folder generated for the day, SOtoaster will panic and close, you will need to delete the folder in order to create a new folder with SOtoaster.**

soToaster *will ignore* any `session-`  folder that does not contain a `curated_speech_only.tsv` and it will ignore any `session` folders placed elsewhere inside of the folder. ergo, under 2 or more layers.

