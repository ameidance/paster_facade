#! /usr/bin/env bash
export PSM=${PSM:-ameidance.paster.facade}
CURDIR=$(cd $(dirname $0); pwd)

exec "$CURDIR/bin/ameidance.paster.facade"