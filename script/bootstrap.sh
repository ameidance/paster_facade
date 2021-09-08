#! /usr/bin/env bash
export PSM=${PSM:-paster_facade}
CURDIR=$(cd $(dirname $0); pwd)

exec "$CURDIR/bin/paster_facade"