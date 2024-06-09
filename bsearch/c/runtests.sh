#!/bin/bash
#------------------------------------------------------------------------------
#
# Bulk test runner
#
#------------------------------------------------------------------------------

TESTS=$1

(
for i in ${TESTS}/*.dat
do
  echo "$(basename ${i})"
  ./bsearch.x < ${i};
  echo ""
done
) > all.log

diff -w all.log ${TESTS}/corr.log
