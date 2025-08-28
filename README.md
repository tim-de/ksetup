# Setup tool for Kohan Ahriman's Gift

A single, simple program for installing common fixes to get Kohan running on
modern systems.

## What does it do?

ksetup checks for the following fixes, downloading and installing the latest
released versions of any that are missing

- [KohanGold](https://github.com/Kohan-Citadel/kohangold-KG-) - An actively maintained community mod
- KohanLauncher.exe - The original launcher, used for selecting modfiles
- [khaldun.net client](https://github.com/Kohan-Citadel/khaldun.net-client) - A wrapper to allow online matchmaking
- [cnc-ddraw](https://github.com/FunkyFr3sh/cnc-ddraw) - A replacement for ddraw that fixes graphical bugs

## How do I use it?

It can be run by either dragging the exe into the install directory for Kohan
and running it from the file manager (may need to be run as admin depending on
how things are set up), or from a terminal by navigating to the Kohan install
directory and then running ksetup from wherever it is saved, as follows:

    > cd "C:\Path\to\Kohan Ahrimans Gift\"
    > C:\Path\to\ksetup.exe

Note that once again the terminal may have to be opened with administrator
privileges.
