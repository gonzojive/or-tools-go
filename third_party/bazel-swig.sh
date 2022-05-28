#!/bin/bash
set -e

# --- begin runfiles.bash initialization v2 ---
# Copy-pasted from the Bazel Bash runfiles library v2.
set -uo pipefail; f=bazel_tools/tools/bash/runfiles/runfiles.bash
source "${RUNFILES_DIR:-/dev/null}/$f" 2>/dev/null || \
source "$(grep -sm1 "^$f " "${RUNFILES_MANIFEST_FILE:-/dev/null}" | cut -f2- -d' ')" 2>/dev/null || \
source "$0.runfiles/$f" 2>/dev/null || \
source "$(grep -sm1 "^$f " "$0.runfiles_manifest" | cut -f2- -d' ')" 2>/dev/null || \
source "$(grep -sm1 "^$f " "$0.exe.runfiles_manifest" | cut -f2- -d' ')" 2>/dev/null || \
  { echo>&2 "ERROR: cannot find $f"; exit 1; }; f=; set -e
# --- end runfiles.bash initialization v2 ---

LIBDIR=`dirname $(rlocation swig/Lib/allkw.swg)`

# Include all of the standard swig templates
# See https://gitlab-extern.ivi.fraunhofer.de/philipps/or-tools/-/blob/9c5f3cffba703fd72d1ae6c75d0d8f10ee341d24/bazel/swig.bzl
$(rlocation swig/swig) \
  "-I$LIBDIR" "-I$LIBDIR/" \
  $@
