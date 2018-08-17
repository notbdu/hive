// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package tcliservice

import (
	"bytes"
	"reflect"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var PRIMITIVE_TYPES []TTypeId
var COMPLEX_TYPES []TTypeId
var COLLECTION_TYPES []TTypeId
var TYPE_NAMES map[TTypeId]string
const CHARACTER_MAXIMUM_LENGTH = "characterMaximumLength"
const PRECISION = "precision"
const SCALE = "scale"

func init() {
PRIMITIVE_TYPES = []TTypeId{
    0,     1,     2,     3,     4,     5,     6,     7,     8,     9,     15,     16,     17,     18,     19,     20,     21,     22, }

COMPLEX_TYPES = []TTypeId{
    10,     11,     12,     13,     14, }

COLLECTION_TYPES = []TTypeId{
    10,     11, }

TYPE_NAMES = map[TTypeId]string{
    0: "BOOLEAN",
    1: "TINYINT",
    2: "SMALLINT",
    3: "INT",
    4: "BIGINT",
    5: "FLOAT",
    6: "DOUBLE",
    7: "STRING",
    8: "TIMESTAMP",
    9: "BINARY",
    10: "ARRAY",
    11: "MAP",
    12: "STRUCT",
    13: "UNIONTYPE",
    15: "DECIMAL",
    16: "NULL",
    17: "DATE",
    18: "VARCHAR",
    19: "CHAR",
    20: "INTERVAL_YEAR_MONTH",
    21: "INTERVAL_DAY_TIME",
    22: "TIMESTAMP WITH LOCAL TIME ZONE",
}

}

