#!/usr/bin/env bash
USAGE="Usage: $0 < multi_layer_tar > tar 

Extracts a multi-layer 'docker save' TAR archive from stdin,
finds the largest archives and then consolidates them within a single
.tar.xz rootfs archive on stdout.

Options:
    -p    Plugin file
    -h    Display this help and exit
    -v    Turn on verbose messages for debugging
"
while getopts "o:p:hv" OPTION; do
    case $OPTION in
    p) PLUGIN_FILE=$OPTARG ;;
    h) echo "$USAGE" && exit 0 ;;
    o) OUTPUT_FILE=$OPTARG ;;
    v) set -x && export VERBOSE=1 && VERBOSE_OPT='-v' ;;
    *) echo "$USAGE" && exit 1 ;;
    esac
done

say() {
	echo "$@" >&2
}

msg() { :; }

set -e

# Check for deps
for DEP in tar jq mktemp; do
	if ! type $DEP &>/dev/null; then
		say "Error: $DEP not found. Please install it first."
		exit 1
	fi
done

TEMP=$(mktemp -d -t cps-image-XXXXXXX)
# shellcheck disable=SC2064
trap "rm -rf $TEMP" EXIT
msg "Temp directory is: $TEMP"

PACKDIR=$TEMP/consolidated
mkdir "$PACKDIR"

cd "$TEMP"
msg "Extracting image"
tar --warning=no-timestamp -xf-
for LAYER in $(jq -r '.[0].Layers[]' manifest.json); do
    msg "Extracting layer $LAYER"
    tar --warning=no-timestamp -xf "$LAYER" -C "$PACKDIR"
    find "$PACKDIR" -name .wh.\* | while read -r PATH; do
        msg $PATH
        /bin/rm -rf "$PATH" "${PATH/.wh./}"
    done
done

cd "$PACKDIR"
msg "Creating consolidated image archive..."

if [[ -f "$PLUGIN_FILE" ]];then
   $PLUGIN_FILE -p $PACKDIR $VERBOSE_OPT
fi

if time tar -cf - . | xz -v --threads=24 > $OUTPUT_FILE; then
    msg "Flattened image archive created successfully"
else
    say "ERROR: Image archive creation failed"
fi
