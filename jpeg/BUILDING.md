# Instructions to build the static libraries

## Current libraries version

The current static libraries are version 2.1.5.1 (see [brew page](https://formulae.brew.sh/formula/jpeg-turbo) (version might change), [official binaries](https://sourceforge.net/projects/libjpeg-turbo/files/2.1.5.1/) and [Github Project Page](https://github.com/libjpeg-turbo/libjpeg-turbo)).
## Darwin and Linux - extract from brew bottles (recommanded)
The static libraries are extracted from the brew bottles.  
To do so, you need to have [homebrew](https://brew.sh) installed.

1. Install jq `brew install jq` (not crucial but easier to process the json)
2. Run `brew info --json jpeg-turbo | jq -r ".[].bottle"` to display the URLs for the bottles of the stable releases are.
3. Choose the architecture and the OS you want to extract the static library for and copy its URL.
4. Get the bottle `curl -L -H "Authorization: Bearer QQ==" -o MyFilename.tar.gz $URL` (replace the $URL field)
5. Run `tar -xf MyFilename.tar.gz`which will create a `jpeg-turbo` repository.
6. Find libturbojpeg.a (if your version number is 2.1.5.1, it will be under `jpeg-turbo/2.1.5.1/lib/libturbojpeg.a`)
7. Copy/rename it for the target architectured: ex "libturbojpeg_linux_arm64.a" This is the ONLY file you need for each architecture.
8. On one architecture only (it doesn't matter which) you must also extract all the .h files from the include folder.

## Darwin and Linux - compile from source
You can compile the static libraries from [source](https://github.com/libjpeg-turbo/libjpeg-turbo) following the building requirements and instructions [here](https://github.com/libjpeg-turbo/libjpeg-turbo/blob/main/BUILDING.md#all-systems).
However, you need to add an additional flag to make it API/ABI-compatible with **libjpeg v8**. The building procedure becomes:

    cd {build_directory}
    cmake -G"Unix Makefiles" -DWITH_JPEG8=1 {source_directory}
    make


## Windows

For windows, the static libraries **and the header files** are extracted from the [official binary release gcc-64](https://sourceforge.net/projects/libjpeg-turbo/files/2.1.5.1/libjpeg-turbo-2.1.5.1-gcc64.exe/download).

1. Follow the instructions of the auto-installer. Then navigate to the repository libjpeg-turbo-gcc64 (probably `C:\libjpeg-turbo-gcc64\`)
2. grab the headers from `C:\libjpeg-turbo-gcc64\include\`
3. and the static library : `C:\libjpeg-turbo-gcc64\lib\libturbojpeg.a` (renamed in the repository libturbojpeg_windows.a)