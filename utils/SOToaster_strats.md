in the main function:
1. run a loop on the directory the executable is placed.
2. target every **folder** in the directory that starts with the name `session-`
3. run a check inside the folder to see if it contains `curated_processed_speech_only.tsv` otherwise it will **ignore** the folder.
4. once all folders are targeted, I will generate a new directory in the same directory with the name `giovanni_damico_tsv_mmddyy`
5. loop through the names of every folder name captured in step 2 and copy the `curated_processed_speech_only.tsv` files renamed to the source folder inside.

this application will be PATH-able and usable on target locations.


look into current working directory


June 26th 2022 - folder is generated, tsv's are generated but files are **not** fully copied over.
